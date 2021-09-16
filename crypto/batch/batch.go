package batch

import (
	"github.com/number571/tendermint/crypto"
	"github.com/number571/tendermint/crypto/gost512"
)

// CreateBatchVerifier checks if a key type implements the batch verifier interface.
// Currently only gost512 supports batch verification.
func CreateBatchVerifier(pk crypto.PubKey) (crypto.BatchVerifier, bool) {

	switch pk.Type() {
	case gost512.KeyType:
		return gost512.NewBatchVerifier(), true
	}

	return nil, false
}

// SupportsBatchVerifier checks if a key type implements the batch verifier
// interface.
func SupportsBatchVerifier(pk crypto.PubKey) bool {
	switch pk.Type() {
	case gost512.KeyType:
		return true
	}

	return false
}
