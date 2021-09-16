package gost512

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	gkeys "bitbucket.org/number571/go-cryptopro/gost_r_34_10_2012"

	"github.com/number571/tendermint/crypto"
	"github.com/number571/tendermint/crypto/tmhash"
	tmjson "github.com/number571/tendermint/libs/json"
)

//-------------------------------------

var (
	_ crypto.PrivKey = PrivKey{}
)

const (
	PrivKeyName = "tendermint/PrivKey512"
	PubKeyName  = "tendermint/PubKey512"

	PubKeySize     = gkeys.PubKeySize512
	PrivateKeySize = gkeys.PrivKeySize512

	SignatureSize = gkeys.SignatureSize512

	ProvType = "512"
	KeyType  = gkeys.KeyType + " " + ProvType
)

func init() {
	tmjson.RegisterType(PubKey{}, PubKeyName)
	tmjson.RegisterType(PrivKey{}, PrivKeyName)
}

type PrivKey gkeys.PrivKey512

func (privKey PrivKey) Bytes() []byte {
	return gkeys.PrivKey512(privKey).Bytes()
}

func (privKey PrivKey) Sign(msg []byte) ([]byte, error) {
	return gkeys.PrivKey512(privKey).Sign(msg)
}

func (privKey PrivKey) PubKey() crypto.PubKey {
	return PubKey(gkeys.PrivKey512(privKey).PubKey().(gkeys.PubKey512))
}

func (privKey PrivKey) Equals(other crypto.PrivKey) bool {
	return bytes.Equal(gkeys.PrivKey512(privKey).Bytes(), other.Bytes())
}

func (privKey PrivKey) Type() string {
	return gkeys.PrivKey512(privKey).Type()
}

func GenPrivKey() PrivKey {
	fmt.Printf("Generating private key [%s]...\n", KeyType)
	return GenPrivKeyWithInput(
		inputString("Subject >>> "),
		inputString("Password >>> "),
	)
}

func GenPrivKeyWithInput(subject, password string) PrivKey {
	cfg := gkeys.NewConfig(gkeys.K512, subject, password)
	gkeys.GenPrivKey(cfg)

	priv, err := gkeys.NewPrivKey(cfg)
	if err != nil {
		panic(err)
	}

	return PrivKey(priv.(gkeys.PrivKey512))
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

type PubKey gkeys.PubKey512

func (pubKey PubKey) Address() crypto.Address {
	return crypto.Address(tmhash.SumTruncated(pubKey.Bytes()))
}

func (pubKey PubKey) Bytes() []byte {
	return gkeys.PubKey512(pubKey).Bytes()
}

func (pubKey PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return gkeys.PubKey512(pubKey).VerifySignature(msg, sig)
}

func (pubKey PubKey) String() string {
	return gkeys.PubKey512(pubKey).String()
}

func (pubKey PubKey) Type() string {
	return gkeys.PubKey512(pubKey).Type()
}

func (pubKey PubKey) Equals(other crypto.PubKey) bool {
	return bytes.Equal(gkeys.PubKey512(pubKey).Bytes(), other.Bytes())
}

//-------------------------------------

var _ crypto.BatchVerifier = &BatchVerifier{}

type BatchVerifier gkeys.BatchVerifierX

func NewBatchVerifier() crypto.BatchVerifier {
	return &BatchVerifier{}
}

func (b *BatchVerifier) Add(key crypto.PubKey, msg, signature []byte) error {
	return (*gkeys.BatchVerifierX)(b).Add(gkeys.PubKey512(key.(PubKey)), msg, signature)
}

func (b *BatchVerifier) Verify() (bool, []bool) {
	return (*gkeys.BatchVerifierX)(b).Verify()
}
