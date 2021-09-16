package gost256

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmjson "github.com/tendermint/tendermint/libs/json"

	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
)

//-------------------------------------
const (
	PrivKeyName = "tendermint/PrivKey256"
	PubKeyName  = "tendermint/PubKey256"

	ProvType = "256"
	KeyType  = gkeys.KeyType + " " + ProvType

	PubKeySize  = gkeys.PubKeySize256
	PrivKeySize = gkeys.PrivKeySize256
)

func init() {
	tmjson.RegisterType(PubKey{}, PubKeyName)
	tmjson.RegisterType(PrivKey{}, PrivKeyName)
}

var _ crypto.PrivKey = PrivKey{}

type PrivKey gkeys.PrivKey256

func (privKey PrivKey) Bytes() []byte {
	return gkeys.PrivKey256(privKey).Bytes()
}

func (privKey PrivKey) PubKey() crypto.PubKey {
	return PubKey(gkeys.PrivKey256(privKey).PubKey().(gkeys.PubKey256))
}

func (privKey PrivKey) Equals(other crypto.PrivKey) bool {
	return bytes.Equal(gkeys.PrivKey256(privKey).Bytes(), other.Bytes())
}

func (privKey PrivKey) Type() string {
	return gkeys.PrivKey256(privKey).Type()
}

func GenPrivKey() PrivKey {
	fmt.Printf("Generating private key [%s]...\n", KeyType)
	return GenPrivKeyWithInput(
		inputString("Subject >>> "),
		inputString("Password >>> "),
	)
}

func GenPrivKeyWithInput(subject, password string) PrivKey {
	cfg := gkeys.NewConfig(gkeys.K256, subject, password)
	gkeys.GenPrivKey(cfg)

	priv, err := gkeys.NewPrivKey(cfg)
	if err != nil {
		panic(err)
	}

	return PrivKey(priv.(gkeys.PrivKey256))
}

func inputString(begin string) string {
	fmt.Print(begin)
	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(data)
}

//-------------------------------------

var _ crypto.PubKey = PubKey{}

type PubKey gkeys.PubKey256

func (pubKey PubKey) Address() crypto.Address {
	return crypto.Address(tmhash.SumTruncated(pubKey.Bytes()))
}

func (pubKey PubKey) Bytes() []byte {
	return gkeys.PubKey256(pubKey).Bytes()
}

func (pubKey PubKey) String() string {
	return gkeys.PubKey256(pubKey).String()
}

func (pubKey PubKey) Equals(other crypto.PubKey) bool {
	return bytes.Equal(gkeys.PubKey256(pubKey).Bytes(), other.Bytes())
}

func (pubKey PubKey) Type() string {
	return gkeys.PubKey256(pubKey).Type()
}
