package types

import (
	fmt "fmt"

	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
	cryptoenc "github.com/number571/tendermint/crypto/encoding"
	"github.com/number571/tendermint/crypto/gost256"
	"github.com/number571/tendermint/crypto/gost512"
)

func Gost512ValidatorUpdate(pk []byte, power int64) ValidatorUpdate {
	pkx, err := gkeys.LoadPubKey(pk)
	if err != nil {
		return ValidatorUpdate{}
	}
	pke := gost512.PubKey(pkx.(gkeys.PubKey512))

	pkp, err := cryptoenc.PubKeyToProto(pke)
	if err != nil {
		panic(err)
	}

	return ValidatorUpdate{
		PubKey: pkp,
		Power:  power,
	}
}

func UpdateValidator(pk []byte, power int64, keyType string) ValidatorUpdate {
	switch keyType {
	case "", gost512.KeyType:
		return Gost512ValidatorUpdate(pk, power)
	case gost256.KeyType:
		pke := gost256.PubKey(pk)
		pkp, err := cryptoenc.PubKeyToProto(pke)
		if err != nil {
			panic(err)
		}
		return ValidatorUpdate{
			PubKey: pkp,
			Power:  power,
		}
	default:
		panic(fmt.Sprintf("key type %s not supported", keyType))
	}
}
