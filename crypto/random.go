package crypto

import (
	"encoding/hex"
	"io"

	grand "bitbucket.org/number571/go-cryptopro/gost_r_iso_28640_2012"
)

func CRandBytes(numBytes int) []byte {
	return grand.Rand(numBytes)
}

func CRandHex(numDigits int) string {
	return hex.EncodeToString(CRandBytes(numDigits / 2))
}

func CReader() io.Reader {
	return grand.Reader
}
