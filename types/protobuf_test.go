package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	abci "github.com/number571/tendermint/abci/types"
	"github.com/number571/tendermint/crypto"
	cryptoenc "github.com/number571/tendermint/crypto/encoding"
	"github.com/number571/tendermint/crypto/gost512"
)

const (
	testSubject  = "testSubject"
	testPassword = "testPassword"
)

func TestABCIPubKey(t *testing.T) {
	pkEd := gost512.GenPrivKeyWithInput(testSubject, testPassword).PubKey()
	err := testABCIPubKey(t, pkEd, ABCIPubKeyTypeGost512)
	assert.NoError(t, err)
}

func testABCIPubKey(t *testing.T, pk crypto.PubKey, typeStr string) error {
	abciPubKey, err := cryptoenc.PubKeyToProto(pk)
	require.NoError(t, err)
	pk2, err := cryptoenc.PubKeyFromProto(abciPubKey)
	require.NoError(t, err)
	require.Equal(t, pk, pk2)
	return nil
}

func TestABCIValidators(t *testing.T) {
	pkEd := gost512.GenPrivKeyWithInput(testSubject, testPassword).PubKey()

	// correct validator
	tmValExpected := NewValidator(pkEd, 10)

	tmVal := NewValidator(pkEd, 10)

	abciVal := TM2PB.ValidatorUpdate(tmVal)
	tmVals, err := PB2TM.ValidatorUpdates([]abci.ValidatorUpdate{abciVal})
	assert.Nil(t, err)
	assert.Equal(t, tmValExpected, tmVals[0])

	abciVals := TM2PB.ValidatorUpdates(NewValidatorSet(tmVals))
	assert.Equal(t, []abci.ValidatorUpdate{abciVal}, abciVals)

	// val with address
	tmVal.Address = pkEd.Address()

	abciVal = TM2PB.ValidatorUpdate(tmVal)
	tmVals, err = PB2TM.ValidatorUpdates([]abci.ValidatorUpdate{abciVal})
	assert.Nil(t, err)
	assert.Equal(t, tmValExpected, tmVals[0])
}

func TestABCIConsensusParams(t *testing.T) {
	cp := DefaultConsensusParams()
	abciCP := TM2PB.ConsensusParams(cp)
	cp2 := UpdateConsensusParams(*cp, abciCP)

	assert.Equal(t, *cp, cp2)
}

type pubKeyEddie struct{}

func (pubKeyEddie) Address() Address                            { return []byte{} }
func (pubKeyEddie) Bytes() []byte                               { return []byte{} }
func (pubKeyEddie) VerifySignature(msg []byte, sig []byte) bool { return false }
func (pubKeyEddie) Equals(crypto.PubKey) bool                   { return false }
func (pubKeyEddie) String() string                              { return "" }
func (pubKeyEddie) Type() string                                { return "pubKeyEddie" }

func TestABCIValidatorFromPubKeyAndPower(t *testing.T) {
	pubkey := gost512.GenPrivKeyWithInput(testSubject, testPassword).PubKey()

	abciVal := TM2PB.NewValidatorUpdate(pubkey, 10)
	assert.Equal(t, int64(10), abciVal.Power)

	assert.Panics(t, func() { TM2PB.NewValidatorUpdate(nil, 10) })
	assert.Panics(t, func() { TM2PB.NewValidatorUpdate(pubKeyEddie{}, 10) })
}

func TestABCIValidatorWithoutPubKey(t *testing.T) {
	pkEd := gost512.GenPrivKeyWithInput(testSubject, testPassword).PubKey()

	abciVal := TM2PB.Validator(NewValidator(pkEd, 10))

	// pubkey must be nil
	tmValExpected := abci.Validator{
		Address: pkEd.Address(),
		Power:   10,
	}

	assert.Equal(t, tmValExpected, abciVal)
}
