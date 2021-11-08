package ed25519consensus_test

import (
	"testing"

	ed25519sha3 "github.com/crpt/go-ed25519-sha3-512"

	"github.com/crpt/go-ed25519consensus-sha3-512"
)

func BenchmarkVerification(b *testing.B) {
	b.ReportAllocs()
	pub, priv, _ := ed25519sha3.GenerateKey(nil)
	hash := []byte("Single key verification")
	signature := ed25519sha3.Sign(priv, hash)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ed25519consensus.Verify(pub, hash, signature)
	}
}
