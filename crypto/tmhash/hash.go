package tmhash

import (
	"hash"

	ghash "github.com/number571/go-cryptopro/gost_r_34_11_2012"
)

const (
	Size      = ghash.Size256
	BlockSize = ghash.BlockSize
)

func New() hash.Hash {
	return ghash.New(ghash.H256)
}

func Sum(bz []byte) []byte {
	h := ghash.Sum(ghash.H256, bz)
	return h[:]
}

//-------------------------------------------------------------

const (
	TruncatedSize = 20
)

type gostTrunc struct {
	ghash hash.Hash
}

func (h gostTrunc) Write(p []byte) (n int, err error) {
	return h.ghash.Write(p)
}

func (h gostTrunc) Sum(b []byte) []byte {
	shasum := h.ghash.Sum(b)
	return shasum[:TruncatedSize]
}

func (h gostTrunc) Reset() {
	h.ghash.Reset()
}

func (h gostTrunc) Size() int {
	return TruncatedSize
}

func (h gostTrunc) BlockSize() int {
	return h.ghash.BlockSize()
}

func NewTruncated() hash.Hash {
	return gostTrunc{
		ghash: ghash.New(ghash.H256),
	}
}

func SumTruncated(bz []byte) []byte {
	hash := ghash.Sum(ghash.H256, bz)
	return hash[:TruncatedSize]
}
