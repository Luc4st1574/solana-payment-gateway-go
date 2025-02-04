package modules

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"bitsnake-server/internal/ent"
	"bitsnake-server/internal/ent/users"
	"bitsnake-server/internal/solana/client"
	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/types"
)

func VerifyPayment(ctx context.Context, db *ent.Client, solanaClient *client.Client, w http.ResponseWriter, receiverWallet string, expectedAmount float64) (bool, error) {
	// 1. Get or create user
	u, err := db.Users.Query().
		Where(users.WalletAddress(receiverWallet)).
		Only(ctx)

	if ent.IsNotFound(err) {
		u, err = db.Users.Create().
			SetWalletAddress(receiverWallet).
			Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to create user: %w", err)
		}
	} else if err != nil {
		return false, fmt.Errorf("failed to query user: %w", err)
	}

	// 2. Create payment verification record
	pv, err := db.PaymentVerifications.Create().
		SetUser(u).
		SetWalletAddress(receiverWallet).
		SetAmount(expectedAmount).
		SetStatus("pending").
		Save(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to create payment verification: %w", err)
	}

	// 3. Retry Mechanism for Verification
	const maxRetries = 5
	const retryDelay = 5 * time.Second
	var verified bool
	var txHash string

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting verification (%d/%d)...", i+1, maxRetries)
		verified, txHash, err = verifyTransaction(ctx, solanaClient, receiverWallet, expectedAmount)
		if err == nil && verified {
			break
		}
		time.Sleep(retryDelay)
	}

	if !verified {
		_, updateErr := pv.Update().
			SetStatus("failed").
			Save(ctx)
		if updateErr != nil {
			log.Printf("Failed to update payment verification: %v", updateErr)
		}
		return false, fmt.Errorf("verification failed: %w", err)
	}

	// 4. Update verification status
	update := pv.Update().
		SetStatus("verified").
		SetTransactionHash(txHash).
		SetAccessGranted(verified)

	if verified {
		_, err = u.Update().
			SetHasAccess(true).
			Save(ctx)
		if err != nil {
			log.Printf("Failed to update user access: %v", err)
		}
	}

	_, err = update.Save(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to update payment verification: %w", err)
	}

	// 5. Create Game Session and set cookie
	gameSession, err := CreateGameSession(ctx, db, receiverWallet, txHash)
	if err != nil {
		return false, fmt.Errorf("failed to create game session: %w", err)
	}

	// Convert session details to JSON
	sessionData, err := json.Marshal(gameSession)
	if err != nil {
		http.Error(w, `{"error": "Failed to serialize session data"}`, http.StatusInternalServerError)
		return false, err
	}

	// Set session cookie
	cookie := http.Cookie{
		Name:     "game_session",
		Value:    string(sessionData),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  gameSession.ExpirationDate,
	}

	http.SetCookie(w, &cookie)

	// Send response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sessionData)

	return verified, nil
}

func verifyTransaction(ctx context.Context, c *client.Client, receiver string, expectedSOL float64) (bool, string, error) {
	expectedLamports := uint64(expectedSOL * 1e9)

	// Get recent transactions
	sigs, err := c.GetSignaturesForAddress(ctx, receiver)
	if err != nil {
		return false, "", fmt.Errorf("failed to get signatures: %w", err)
	}

	// Check last 20 transactions
	maxTransactions := 20
	if len(sigs) < maxTransactions {
		maxTransactions = len(sigs)
	}

	for _, sig := range sigs[:maxTransactions] {
		tx, err := c.GetTransaction(ctx, sig.Signature)
		if err != nil {
			log.Printf("Error fetching transaction %s: %v", sig.Signature, err)
			continue
		}

		if tx.Meta.Err != nil {
			continue // Ignore failed transactions
		}

		log.Printf("Checking transaction: %s", sig.Signature)

		// Validate transfer details
		for _, instr := range tx.Transaction.Message.Instructions {
			if isSystemTransfer(instr, tx.AccountKeys) {
				if len(instr.Data) >= 12 {
					amount := binary.LittleEndian.Uint64(instr.Data[4:12])
					if amount == expectedLamports {
						log.Printf("Transaction %s matches expected amount!", sig.Signature)
						return true, sig.Signature, nil
					}
				}
			}
		}
	}

	return false, "", nil
}

func isSystemTransfer(instr types.CompiledInstruction, accounts []common.PublicKey) bool {
	if len(accounts) <= instr.ProgramIDIndex {
		return false
	}
	return accounts[instr.ProgramIDIndex] == common.SystemProgramID
}
