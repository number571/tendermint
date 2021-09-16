package types

import (
	"github.com/number571/tendermint/crypto/gost512"
	tmmath "github.com/number571/tendermint/libs/math"
)

var (
	// MaxSignatureSize is a maximum allowed signature size for the Proposal
	// and Vote.
	MaxSignatureSize = tmmath.MaxInt(gost512.SignatureSize, 64)
)

// Signable is an interface for all signable things.
// It typically removes signatures before serializing.
// SignBytes returns the bytes to be signed
// NOTE: chainIDs are part of the SignBytes but not
// necessarily the object themselves.
// NOTE: Expected to panic if there is an error marshaling.
type Signable interface {
	SignBytes(chainID string) []byte
}
