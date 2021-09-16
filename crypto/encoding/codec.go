package encoding

import (
	"fmt"

	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/gost256"
	"github.com/tendermint/tendermint/crypto/gost512"
	"github.com/tendermint/tendermint/libs/json"
	pc "github.com/tendermint/tendermint/proto/tendermint/crypto"
)

func init() {
	json.RegisterType((*pc.PublicKey)(nil), "tendermint.crypto.PublicKey")
	json.RegisterType((*pc.PublicKey_Gost512)(nil), "tendermint.crypto.PublicKey_Gost512")
	json.RegisterType((*pc.PublicKey_Gost256)(nil), "tendermint.crypto.PublicKey_Gost256")
}

// PubKeyToProto takes crypto.PubKey and transforms it to a protobuf Pubkey
func PubKeyToProto(k crypto.PubKey) (pc.PublicKey, error) {
	var kp pc.PublicKey
	switch k := k.(type) {
	case gost512.PubKey:
		kp = pc.PublicKey{
			Sum: &pc.PublicKey_Gost512{
				Gost512: k.Bytes(),
			},
		}
	case gost256.PubKey:
		kp = pc.PublicKey{
			Sum: &pc.PublicKey_Gost256{
				Gost256: k.Bytes(),
			},
		}
	default:
		return kp, fmt.Errorf("toproto: key type %v is not supported", k)
	}
	return kp, nil
}

// PubKeyFromProto takes a protobuf Pubkey and transforms it to a crypto.Pubkey
func PubKeyFromProto(k pc.PublicKey) (crypto.PubKey, error) {
	switch k := k.Sum.(type) {
	case *pc.PublicKey_Gost512:
		if len(k.Gost512) != gost512.PubKeySize {
			return nil, fmt.Errorf("invalid size for PubKeyGost512. Got %d, expected %d",
				len(k.Gost512), gost512.PubKeySize)
		}
		pk, err := gkeys.LoadPubKey(k.Gost512)
		if err != nil {
			return nil, err
		}
		return gost512.PubKey(pk.(gkeys.PubKey512)), nil
	case *pc.PublicKey_Gost256:
		if len(k.Gost256) != gost256.PubKeySize {
			return nil, fmt.Errorf("invalid size for PubKeyGost256. Got %d, expected %d",
				len(k.Gost256), gost256.PubKeySize)
		}
		pk, err := gkeys.LoadPubKey(k.Gost256)
		if err != nil {
			return nil, err
		}
		return gost256.PubKey(pk.(gkeys.PubKey256)), nil
	default:
		return nil, fmt.Errorf("fromproto: key type %v is not supported", k)
	}
}
