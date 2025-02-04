package modules

import (
	"context"

	"bitsnake-server/internal/solana/client"
	"bitsnake-server/internal/solana/common"
	"bitsnake-server/internal/solana/program/system"
	"bitsnake-server/internal/solana/rpc"
	"bitsnake-server/internal/solana/types"
)

func SendPayment(fromAccountPrivateKey, toAccountPublicKey string, amount uint64) (string, error) {
	// Create a new Solana client
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// Create the sender account from the provided private key
	fromAccount, err := types.AccountFromBase58(fromAccountPrivateKey)
	if err != nil {
		return "", err
	}

	// Fetch recent blockhash
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		return "", err
	}

	// Create a transfer transaction
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{fromAccount},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        fromAccount.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   fromAccount.PublicKey,
					To:     common.PublicKeyFromString(toAccountPublicKey),
					Amount: amount, // Amount in lamports (1 SOL = 1e9 lamports)
				}),
			},
		}),
	})
	if err != nil {
		return "", err
	}

	// Send the transaction
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return txhash, nil
}
