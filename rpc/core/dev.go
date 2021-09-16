package core

import (
	ctypes "github.com/number571/tendermint/rpc/core/types"
	rpctypes "github.com/number571/tendermint/rpc/jsonrpc/types"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func (env *Environment) UnsafeFlushMempool(ctx *rpctypes.Context) (*ctypes.ResultUnsafeFlushMempool, error) {
	env.Mempool.Flush()
	return &ctypes.ResultUnsafeFlushMempool{}, nil
}
