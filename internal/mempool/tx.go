package mempool

import (
	ghash "bitbucket.org/number571/go-cryptopro/gost_r_34_11_2012"

	"github.com/number571/tendermint/types"
)

// TxKeySize defines the size of the transaction's key used for indexing.
const TxKeySize = ghash.Size256

// TxKey is the fixed length array key used as an index.
func TxKey(tx types.Tx) [TxKeySize]byte {
	var res [TxKeySize]byte
	copy(res[:], ghash.Sum(ghash.H256, tx))
	return res
}

// TxHashFromBytes returns the hash of a transaction from raw bytes.
func TxHashFromBytes(tx []byte) []byte {
	return types.Tx(tx).Hash()
}

// TxInfo are parameters that get passed when attempting to add a tx to the
// mempool.
type TxInfo struct {
	// SenderID is the internal peer ID used in the mempool to identify the
	// sender, storing two bytes with each transaction instead of 20 bytes for
	// the types.NodeID.
	SenderID uint16

	// SenderNodeID is the actual types.NodeID of the sender.
	SenderNodeID types.NodeID
}
