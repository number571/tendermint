package blockchain

import (
	bcproto "github.com/number571/tendermint/proto/tendermint/blockchain"
	"github.com/number571/tendermint/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		bcproto.BlockResponseMessagePrefixSize +
		bcproto.BlockResponseMessageFieldKeySize
)
