package gost256_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/number571/tendermint/crypto"
	"github.com/number571/tendermint/crypto/gost256"
)

const (
	TEST_SUBJECT  = "subject"
	TEST_PASSWORD = "password"
)

func TestSignAndValidateGost256(t *testing.T) {
	privKey := gost256.GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
	pubKey := privKey.PubKey()

	msg := crypto.CRandBytes(128)
	sig, err := privKey.Sign(msg)
	require.Nil(t, err)

	assert.True(t, pubKey.VerifySignature(msg, sig))

	sig[3] ^= byte(0x01)

	assert.False(t, pubKey.VerifySignature(msg, sig))
}
