// +build !libgost256

package gost256

import (
	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
)

func (privKey PrivKey) Sign(msg []byte) ([]byte, error) {
	return gkeys.PrivKey256(privKey).Sign(msg)
}

func (pubKey PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return gkeys.PubKey256(pubKey).VerifySignature(msg, sig)
}
