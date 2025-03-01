package types

import (
	"bitsnake-server/internal/solana/common"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
)

var (
	ErrAccountFailedToBase58Decode     = errors.New("failed to base58 decode")
	ErrAccountFailedToHexDecode        = errors.New("failed to hex decode")
	ErrAccountPrivateKeyLengthMismatch = errors.New("key length mismatch")
)

type Account struct {
	PublicKey  common.PublicKey
	PrivateKey ed25519.PrivateKey
}

func NewAccount() Account {
	_, X, _ := ed25519.GenerateKey(nil)
	account, _ := AccountFromBytes(X)
	return account
}

// AccountFromBytes generate a account by bytes private key
func AccountFromBytes(key []byte) (Account, error) {
	if len(key) != ed25519.PrivateKeySize {
		return Account{}, fmt.Errorf("%w, expected: %v, got: %v", ErrAccountPrivateKeyLengthMismatch, ed25519.PrivateKeySize, len(key))
	}

	// Check if the key can derive a valid public key
	priKey := ed25519.PrivateKey(key)
	pubKey := priKey.Public()
	if _, ok := pubKey.(ed25519.PublicKey); !ok {
		return Account{}, fmt.Errorf("invalid Ed25519 private key")
	}

	return Account{
		PublicKey:  common.PublicKeyFromBytes(pubKey.(ed25519.PublicKey)),
		PrivateKey: priKey,
	}, nil
}

// AccountFromBase58 generate a account by base58 private key
func AccountFromBase58(key string) (Account, error) {
	b, err := base58.Decode(key)
	if err != nil {
		return Account{}, fmt.Errorf("%w, err: %v", ErrAccountFailedToBase58Decode, err)
	}
	return AccountFromBytes(b)
}

// AccountFromHex generate a account by hex private key
func AccountFromHex(key string) (Account, error) {
	b, err := hex.DecodeString(key)
	if err != nil {
		return Account{}, fmt.Errorf("%w, err: %v", ErrAccountFailedToHexDecode, err)
	}
	return AccountFromBytes(b)
}

// AccountFromSeed generate a account by seed
func AccountFromSeed(seed []byte) (Account, error) {
	pk := ed25519.NewKeyFromSeed(seed)
	return AccountFromBytes(pk)
}

func (a Account) Sign(message []byte) []byte {
	return ed25519.Sign(a.PrivateKey, message)
}
