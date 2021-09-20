package gost512_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/number571/tendermint/crypto"
	"github.com/number571/tendermint/crypto/gost512"
)

const (
	TEST_SUBJECT  = "subject"
	TEST_PASSWORD = "password"
)

func TestSignAndValidateGost512(t *testing.T) {

	privKey := gost512.GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
	pubKey := privKey.PubKey()

	msg := crypto.CRandBytes(128)
	sig, err := privKey.Sign(msg)
	require.Nil(t, err)

	assert.True(t, pubKey.VerifySignature(msg, sig))

	sig[7] ^= byte(0x01)

	assert.False(t, pubKey.VerifySignature(msg, sig))
}

//--- FROM NEW VERSIONS

// func TestBatchSafe(t *testing.T) {
// 	v := gost512.NewBatchVerifier()

// 	for i := 0; i <= 38; i++ {
// 		priv := gost512.GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
// 		pub := priv.PubKey()

// 		var msg []byte
// 		if i%2 == 0 {
// 			msg = []byte("easter")
// 		} else {
// 			msg = []byte("egg")
// 		}

// 		sig, err := priv.Sign(msg)
// 		require.NoError(t, err)

// 		err = v.Add(pub, msg, sig)
// 		require.NoError(t, err)
// 	}

// 	ok, _ := v.Verify()
// 	require.True(t, ok)
// }
