package conp

import (
	"testing"

	"github.com/karlmcguire/conp/simple"
)

func BenchmarkSimpleSingle(b *testing.B) {
	var (
		bucket = simple.New()
		zipf   = MockZipf(b.N)
	)

	for n := 0; n < b.N; n++ {
		bucket.Fetch(MockFetch(zipf).Unpack())
	}
}

func BenchmarkSimpleMultiple(b *testing.B) {
	bucket := simple.New()

	b.RunParallel(func(pb *testing.PB) {
		zipf := MockZipf(b.N)

		for pb.Next() {
			bucket.Fetch(MockFetch(zipf).Unpack())
		}
	})
}
