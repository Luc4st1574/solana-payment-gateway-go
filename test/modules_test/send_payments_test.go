package test

import (
	"testing"

	"bitsnake-server/internal/modules" // Ensure this path matches your module name & structure

	"github.com/stretchr/testify/assert"
)

func TestTransferSOLAndVerifyPayment(t *testing.T) {
	// Log the start of the test.
	t.Log("=== Starting TestTransferSOLAndVerifyPayment ===")
	t.Log("This test simulates transferring SOL and verifies that a valid transaction hash is returned.")

	// Define test variables.
	fromAccountPrivateKey := "2cur21jM4RBN4rzu32KF5gwHL6hTkZwReXgAzRJbRnCzAzNj8WvYckoMALLZiAaKwo9iSc9Cv3pVM8U49n2AhVXw"
	toAccountPublicKey := "D7LwfYCjLLCaeLTTijwBagFAmB3aPSm2Fx8K2DzvqLrz"
	transferAmount := uint64(10000000)

	t.Logf("From Account Private Key: %s", fromAccountPrivateKey)
	t.Logf("To Account Public Key: %s", toAccountPublicKey)
	t.Logf("Transfer Amount (in lamports): %d", transferAmount)

	// Invoke the SendPayment function to initiate the SOL transfer.
	t.Log("Calling modules.SendPayment to initiate the SOL transfer...")
	txHash, err := modules.SendPayment(fromAccountPrivateKey, toAccountPublicKey, transferAmount)
	if err != nil {
		t.Logf("Error encountered during SOL transfer: %v", err)
	} else {
		t.Log("SOL transfer initiated successfully.")
	}
	t.Logf("Transaction hash returned: %s", txHash)

	// Validate that no error occurred and that a transaction hash was returned.
	assert.NoError(t, err, "TransferSOL should not return an error")
	assert.NotEmpty(t, txHash, "Transaction hash should not be empty")

	t.Log("=== TestTransferSOLAndVerifyPayment completed successfully ===")
}
