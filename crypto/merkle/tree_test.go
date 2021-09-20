package merkle

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/number571/tendermint/crypto/tmhash"
	tmrand "github.com/number571/tendermint/libs/rand"
	ctest "github.com/tendermint/tendermint/libs/test"
)

type testItem []byte

func (tI testItem) Hash() []byte {
	return []byte(tI)
}

func TestHashFromByteSlices(t *testing.T) {
	testcases := map[string]struct {
		slices     [][]byte
		expectHash string // in hex format
	}{
		"nil":          {nil, "3f539a213e97c802cc229d474c6aa32a825a360b2a933a949fd925208d9ce1bb"},
		"empty":        {[][]byte{}, "3f539a213e97c802cc229d474c6aa32a825a360b2a933a949fd925208d9ce1bb"},
		"single":       {[][]byte{{1, 2, 3}}, "a19da7f6904a04f5a963c3a3af15e6dc12d8dd676dd5b8402341dc9d2876a2c2"},
		"single blank": {[][]byte{{}}, "6f7305265dc0937440881f9493ef1260f61a9d47742d369e952d41bdb2a9edd1"},
		"two":          {[][]byte{{1, 2, 3}, {4, 5, 6}}, "50dc2f71b990d8b4ee83435730a4364d9f2a26d18ffa591556f4578fae248174"},
		"many": {
			[][]byte{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			"3c84088a0ba8e600164620b4d6efe371a83f652ebd42f7240544b2fa98e43b24",
		},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			hash := HashFromByteSlices(tc.slices)
			assert.Equal(t, tc.expectHash, hex.EncodeToString(hash))
		})
	}
}

func TestProof(t *testing.T) {

	// Try an empty proof first
	rootHash, proofs := ProofsFromByteSlices([][]byte{})

	require.Equal(t, "3f539a213e97c802cc229d474c6aa32a825a360b2a933a949fd925208d9ce1bb", hex.EncodeToString(rootHash))
	require.Empty(t, proofs)

	total := 100

	items := make([][]byte, total)
	for i := 0; i < total; i++ {
		items[i] = testItem(tmrand.Bytes(tmhash.Size))
	}

	rootHash = HashFromByteSlices(items)

	rootHash2, proofs := ProofsFromByteSlices(items)

	require.Equal(t, rootHash, rootHash2, "Unmatched root hashes: %X vs %X", rootHash, rootHash2)

	// For each item, check the trail.
	for i, item := range items {
		proof := proofs[i]

		// Check total/index
		require.EqualValues(t, proof.Index, i, "Unmatched indicies: %d vs %d", proof.Index, i)

		require.EqualValues(t, proof.Total, total, "Unmatched totals: %d vs %d", proof.Total, total)

		// Verify success
		err := proof.Verify(rootHash, item)
		require.NoError(t, err, "Verification failed: %v.", err)

		// Trail too long should make it fail
		origAunts := proof.Aunts
		proof.Aunts = append(proof.Aunts, tmrand.Bytes(32))
		err = proof.Verify(rootHash, item)
		require.Error(t, err, "Expected verification to fail for wrong trail length")

		proof.Aunts = origAunts

		// Trail too short should make it fail
		proof.Aunts = proof.Aunts[0 : len(proof.Aunts)-1]
		err = proof.Verify(rootHash, item)
		require.Error(t, err, "Expected verification to fail for wrong trail length")

		proof.Aunts = origAunts

		// Mutating the itemHash should make it fail.
		err = proof.Verify(rootHash, ctest.MutateByteSlice(item))
		require.Error(t, err, "Expected verification to fail for mutated leaf hash")

		// Mutating the rootHash should make it fail.
		err = proof.Verify(ctest.MutateByteSlice(rootHash), item)
		require.Error(t, err, "Expected verification to fail for mutated root hash")
	}
}

func TestHashAlternatives(t *testing.T) {

	total := 100

	items := make([][]byte, total)
	for i := 0; i < total; i++ {
		items[i] = testItem(tmrand.Bytes(tmhash.Size))
	}

	rootHash1 := HashFromByteSlicesIterative(items)
	rootHash2 := HashFromByteSlices(items)
	require.Equal(t, rootHash1, rootHash2, "Unmatched root hashes: %X vs %X", rootHash1, rootHash2)
}

func BenchmarkHashAlternatives(b *testing.B) {
	total := 100

	items := make([][]byte, total)
	for i := 0; i < total; i++ {
		items[i] = testItem(tmrand.Bytes(tmhash.Size))
	}

	b.ResetTimer()
	b.Run("recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = HashFromByteSlices(items)
		}
	})

	b.Run("iterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = HashFromByteSlicesIterative(items)
		}
	})
}

func Test_getSplitPoint(t *testing.T) {
	tests := []struct {
		length int64
		want   int64
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{10, 8},
		{20, 16},
		{100, 64},
		{255, 128},
		{256, 128},
		{257, 256},
	}
	for _, tt := range tests {
		got := getSplitPoint(tt.length)
		require.EqualValues(t, tt.want, got, "getSplitPoint(%d) = %v, want %v", tt.length, got, tt.want)
	}
}
