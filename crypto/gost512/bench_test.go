package gost512

import (
	"testing"

	"github.com/number571/tendermint/crypto/internal/benchmarking"
)

const (
	TEST_SUBJECT  = "subject"
	TEST_PASSWORD = "password"
)

func BenchmarkSigning(b *testing.B) {
	priv := GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
	benchmarking.BenchmarkSigning(b, priv)
}

func BenchmarkVerification(b *testing.B) {
	priv := GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
	benchmarking.BenchmarkVerification(b, priv)
}

//--- FROM NEW VERSIONS

// func BenchmarkVerifyBatch(b *testing.B) {
// 	msg := []byte("BatchVerifyTest")

// 	for _, sigsCount := range []int{1, 8, 64, 512} {
// 		sigsCount := sigsCount
// 		b.Run(fmt.Sprintf("sig-count-%d", sigsCount), func(b *testing.B) {
// 			// Pre-generate all of the keys, and signatures, but do not
// 			// benchmark key-generation and signing.
// 			pubs := make([]crypto.PubKey, 0, sigsCount)
// 			sigs := make([][]byte, 0, sigsCount)
// 			for i := 0; i < sigsCount; i++ {
// 				priv := GenPrivKeyWithInput(TEST_SUBJECT, TEST_PASSWORD)
// 				sig, _ := priv.Sign(msg)
// 				pubs = append(pubs, priv.PubKey().(PubKey))
// 				sigs = append(sigs, sig)
// 			}
// 			b.ResetTimer()

// 			b.ReportAllocs()
// 			// NOTE: dividing by n so that metrics are per-signature
// 			for i := 0; i < b.N/sigsCount; i++ {
// 				// The benchmark could just benchmark the Verify()
// 				// routine, but there is non-trivial overhead associated
// 				// with BatchVerifier.Add(), which should be included
// 				// in the benchmark.
// 				v := NewBatchVerifier()
// 				for i := 0; i < sigsCount; i++ {
// 					err := v.Add(pubs[i], msg, sigs[i])
// 					require.NoError(b, err)
// 				}

// 				if ok, _ := v.Verify(); !ok {
// 					b.Fatal("signature set failed batch verification")
// 				}
// 			}
// 		})
// 	}
// }
