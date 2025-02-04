package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"bitsnake-server/internal/ent"
	"bitsnake-server/internal/ent/users"
)

// GameSession structure
type GameSession struct {
	GameHashID      string    `json:"gameHashID"`
	ExpirationDate  time.Time `json:"expirationDate"`
	WalletAddress   string    `json:"walletAddress"`
	TransactionHash string    `json:"transactionHash"`
}

// CreateGameSession generates a session and stores it as an HTTP cookie
func CreateGameSession(ctx context.Context, db *ent.Client, walletAddress, transactionHash string) (*GameSession, error) {
	// Check if the user has access
	user, err := db.Users.Query().
		Where(users.WalletAddress(walletAddress)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if !user.HasAccess {
		return nil, fmt.Errorf("user does not have verified access")
	}

	// Generate a unique GameHashID and expiration time
	expirationTime := time.Now().Add(1 * time.Hour) // Session expires in 1 hour
	gameHashID := fmt.Sprintf("%x", rand.Int63())

	// Save the session in the database
	session, err := db.Matches.Create().
		SetGameHashID(gameHashID).
		SetExpirationDate(expirationTime).
		SetWalletAddress(walletAddress).
		SetTransactionHash(transactionHash).
		// Associate the session with the user. This is required by the schema.
		SetUser(user).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to create game session: %w", err)
	}

	// Return session details
	return &GameSession{
		GameHashID:      session.GameHashID,
		ExpirationDate:  session.ExpirationDate,
		WalletAddress:   session.WalletAddress,
		TransactionHash: session.TransactionHash,
	}, nil
}

// GameSessionHandler handles session creation and stores it as a cookie
func GameSessionHandler(db *ent.Client, w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	walletAddress := r.URL.Query().Get("walletAddress")
	transactionHash := r.URL.Query().Get("transactionHash")

	if walletAddress == "" || transactionHash == "" {
		http.Error(w, `{"error": "Missing walletAddress or transactionHash"}`, http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	session, err := CreateGameSession(ctx, db, walletAddress, transactionHash)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusUnauthorized)
		return
	}

	// Convert session details to JSON
	sessionData, err := json.Marshal(session)
	if err != nil {
		http.Error(w, `{"error": "Failed to serialize session data"}`, http.StatusInternalServerError)
		return
	}

	// Set cookie with session data
	cookie := http.Cookie{
		Name:     "game_session",
		Value:    string(sessionData),
		Path:     "/",
		HttpOnly: true, // Prevents client-side access
		Secure:   true, // Ensures it's only sent over HTTPS
		Expires:  session.ExpirationDate,
	}

	http.SetCookie(w, &cookie)

	// Send response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(sessionData)
}
