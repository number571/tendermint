package crypto

import (
	ghash "github.com/number571/go-cryptopro/gost_r_34_11_2012"
)

func HashSum(bytes []byte) []byte {
	return ghash.Sum(ghash.H256, bytes)
}
