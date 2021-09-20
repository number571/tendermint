package mempool

import (
	"crypto/rand"
	"testing"

	ghash "github.com/number571/go-cryptopro/gost_r_34_11_2012"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/abci/example/kvstore"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tendermint/types"
)

func TestCacheRemove(t *testing.T) {
	cache := newMapTxCache(100)
	numTxs := 10
	txs := make([][]byte, numTxs)
	for i := 0; i < numTxs; i++ {
		// probability of collision is 2**-256
		txBytes := make([]byte, 32)
		_, err := rand.Read(txBytes)
		require.NoError(t, err)
		txs[i] = txBytes
		cache.Push(txBytes)
		// make sure its added to both the linked list and the map
		require.Equal(t, i+1, len(cache.cacheMap))
		require.Equal(t, i+1, cache.list.Len())
	}
	for i := 0; i < numTxs; i++ {
		cache.Remove(txs[i])
		// make sure its removed from both the map and the linked list
		require.Equal(t, numTxs-(i+1), len(cache.cacheMap))
		require.Equal(t, numTxs-(i+1), cache.list.Len())
	}
}

func TestCacheAfterUpdate(t *testing.T) {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	mempool, cleanup := newMempoolWithApp(cc)
	defer cleanup()

	// reAddIndices & txsInCache can have elements > numTxsToCreate
	// also assumes max index is 255 for convenience
	// txs in cache also checks order of elements
	tests := []struct {
		numTxsToCreate int
		updateIndices  []int
		reAddIndices   []int
		txsInCache     []int
	}{
		{1, []int{}, []int{1}, []int{1, 0}},    // adding new txs works
		{2, []int{1}, []int{}, []int{1, 0}},    // update doesn't remove tx from cache
		{2, []int{2}, []int{}, []int{2, 1, 0}}, // update adds new tx to cache
		{2, []int{1}, []int{1}, []int{1, 0}},   // re-adding after update doesn't make dupe
	}
	for tcIndex, tc := range tests {
		for i := 0; i < tc.numTxsToCreate; i++ {
			tx := types.Tx{byte(i)}
			err := mempool.CheckTx(tx, nil, TxInfo{})
			require.NoError(t, err)
		}

		updateTxs := []types.Tx{}
		for _, v := range tc.updateIndices {
			tx := types.Tx{byte(v)}
			updateTxs = append(updateTxs, tx)
		}
		err := mempool.Update(int64(tcIndex), updateTxs, abciResponses(len(updateTxs), abci.CodeTypeOK), nil, nil)
		require.NoError(t, err)

		for _, v := range tc.reAddIndices {
			tx := types.Tx{byte(v)}
			_ = mempool.CheckTx(tx, nil, TxInfo{})
		}

		cache := mempool.cache.(*mapTxCache)
		node := cache.list.Front()
		counter := 0
		for node != nil {
			require.NotEqual(t, len(tc.txsInCache), counter,
				"cache larger than expected on testcase %d", tcIndex)

			nodeVal := node.Value.([ghash.Size256]byte)
			expectedBz := ghash.Sum(ghash.H256, []byte{byte(tc.txsInCache[len(tc.txsInCache)-counter-1])})

			require.Equal(t, expectedBz, nodeVal, "Equality failed on index %d, tc %d", counter, tcIndex)
			counter++
			node = node.Next()
		}
		require.Equal(t, len(tc.txsInCache), counter,
			"cache smaller than expected on testcase %d", tcIndex)
		mempool.Flush()
	}
}
