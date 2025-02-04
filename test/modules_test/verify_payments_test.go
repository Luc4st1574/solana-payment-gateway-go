package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"bitsnake-server/internal/ent"
	"bitsnake-server/internal/ent/paymentverifications"
	"bitsnake-server/internal/ent/users"
	"bitsnake-server/internal/modules"
	"bitsnake-server/internal/solana/client"
	"bitsnake-server/internal/solana/rpc"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestVerifyPaymentIntegration simulates the full flow of sending a payment, waiting for its confirmation,
// verifying it on the blockchain, and ensuring the database and user state are updated accordingly.
func TestVerifyPaymentIntegration(t *testing.T) {
	// Define test constants. These represent the sender's private key, the receiver's wallet address,
	// the amount to be sent (in lamports), the expected amount (in SOL), and delays for confirmation and test timeout.
	const (
		senderPrivateKey  = "2cur21jM4RBN4rzu32KF5gwHL6hTkZwReXgAzRJbRnCzAzNj8WvYckoMALLZiAaKwo9iSc9Cv3pVM8U49n2AhVXw"
		receiverWallet    = "D7LwfYCjLLCaeLTTijwBagFAmB3aPSm2Fx8K2DzvqLrz"
		amountLamports    = uint64(10000000)
		expectedSOL       = 0.01
		testTimeout       = 60 * time.Second
		confirmationDelay = 30 * time.Second
	)

	// Create a context with timeout to ensure the test does not run indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	// Step 0: Setup - Initialize the test database and Solana client.

	t.Log("Initializing test database connection...")
	dbClient := setupTestDB(t)
	defer dbClient.Close()

	t.Log("Initializing Solana client (Devnet)...")
	solanaClient := client.NewClient(rpc.DevnetRPCEndpoint)

	// Step 1: Payment Submission
	// In this stage, we simulate a payment being sent from a sender to the receiver.

	t.Log("Step 1: Initiating payment transaction.")
	txHash, err := modules.SendPayment(senderPrivateKey, receiverWallet, amountLamports)
	require.NoError(t, err, "Failed to send payment")
	require.NotEmpty(t, txHash, "Empty transaction hash returned")
	t.Logf("Payment transaction sent successfully. Transaction hash: %s", txHash)

	// Step 2: Blockchain Confirmation Wait
	// We wait for a predetermined delay to allow the blockchain to confirm the transaction.

	t.Logf("Step 2: Waiting for %s to allow the transaction to be confirmed...", confirmationDelay)
	time.Sleep(confirmationDelay)

	// Step 3: Payment Verification Process
	// Now that the payment is assumed to be confirmed, we start the verification process.
	// This includes checking recent transactions on the blockchain and updating our records.
	t.Log("Step 3: Verifying the payment transaction and updating database records.")

	// Create a dummy HTTP response writer to capture the cookie and response output.
	recorder := httptest.NewRecorder()
	verified, err := modules.VerifyPayment(ctx, dbClient, solanaClient, recorder, receiverWallet, expectedSOL)
	require.NoError(t, err, "Payment verification failed")
	assert.True(t, verified, "Payment should be verified successfully")
	t.Log("Payment verification succeeded.")

	// Log details from the HTTP response (cookie details, session info, etc.)
	t.Logf("HTTP Response Code: %d", recorder.Code)
	t.Logf("HTTP Response Body: %s", recorder.Body.String())

	// Optionally decode the returned game session to log its details.
	var session modules.GameSession
	if err := json.Unmarshal(recorder.Body.Bytes(), &session); err == nil {
		t.Logf("Created game session details: %+v", session)
	} else {
		t.Log("Failed to decode game session details from HTTP response.")
	}

	// Step 4: Database State Validation
	// Verify that the payment verification record in the database reflects the successful verification.
	t.Log("Step 4: Validating database state for the payment verification record.")
	verifyDatabaseState(t, dbClient, receiverWallet, txHash, expectedSOL)

	// Step 5: User Access Check
	// Confirm that the user's record has been updated to grant access post verification.

	t.Log("Step 5: Checking that the user's access rights have been correctly granted.")
	verifyUserAccess(t, dbClient, receiverWallet)

	t.Log("Test completed: Payment verification integration test passed successfully.")
}

// setupTestDB initializes the test database connection and applies the necessary schema migrations.
func setupTestDB(t *testing.T) *ent.Client {
	t.Log("Setting up the test database...")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=testdb password=admin123 sslmode=disable")
	require.NoError(t, err, "Database connection failed")

	err = client.Schema.Create(context.Background())
	require.NoError(t, err, "Schema migration failed")
	t.Log("Test database setup complete.")

	return client
}

// verifyDatabaseState fetches the payment verification record from the database and checks that
// all fields reflect a successful verification process.
func verifyDatabaseState(t *testing.T, db *ent.Client, walletAddress string, txHash string, expectedSOL float64) {
	t.Log("Fetching payment verification record from the database...")
	pv, err := db.PaymentVerifications.Query().
		Where(
			paymentverifications.WalletAddress(walletAddress),
			paymentverifications.TransactionHash(txHash),
		).
		WithUser().
		Only(context.Background())
	require.NoError(t, err, "Payment verification record not found")

	t.Logf("Payment verification record found: %+v", pv)

	// Check that the status is marked as "verified".
	t.Log("Verifying that the status is 'verified'.")
	assert.Equal(t, paymentverifications.Status("verified"), pv.Status, "Incorrect verification status")

	// Confirm that access was granted.
	t.Log("Confirming that access has been granted.")
	assert.True(t, pv.AccessGranted, "Access should be granted")

	// Validate that the stored transaction hash matches the one from the payment.
	t.Log("Validating the transaction hash in the record.")
	assert.Equal(t, txHash, pv.TransactionHash, "Transaction hash mismatch")

	// Compare the expected SOL amount with the recorded amount.
	t.Log("Comparing the expected SOL amount with the record.")
	assert.InDelta(t, expectedSOL, pv.Amount, 0.00000001, "Amount verification failed")

	// Ensure the user relationship is properly linked.
	t.Log("Ensuring that the user associated with the record is correctly linked.")
	require.NotNil(t, pv.Edges.User, "User relationship not loaded")
	assert.Equal(t, walletAddress, pv.Edges.User.WalletAddress, "User wallet address mismatch")
}

// verifyUserAccess checks that the user record in the database has been updated to grant access.
func verifyUserAccess(t *testing.T, db *ent.Client, walletAddress string) {
	t.Log("Fetching user record from the database...")
	user, err := db.Users.Query().
		Where(users.WalletAddress(walletAddress)).
		Only(context.Background())
	require.NoError(t, err, "User record not found")

	t.Logf("User record retrieved: %+v", user)
	t.Log("Verifying that the user's access has been granted.")
	assert.True(t, user.HasAccess, "User access not granted")
}

// go test -v ./...
