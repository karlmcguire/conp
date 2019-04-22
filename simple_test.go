package conp

import (
	"testing"

	"github.com/karlmcguire/conp/promise"
	"github.com/karlmcguire/conp/simple"
)

const N = 10000

func BenchmarkSimpleSingle(b *testing.B) {
	bucket := simple.New()
	zipf := MockZipf(b.N)

	for n := 0; n < b.N; n++ {
		bucket.Fetch(MockFetch(zipf).Unpack())
	}
}

func BenchmarkSimpleMultiple(b *testing.B) {
	bucket := simple.New()

	b.SetParallelism(N)
	b.RunParallel(func(pb *testing.PB) {
		zipf := MockZipf(b.N)

		for pb.Next() {
			bucket.Fetch(MockFetch(zipf).Unpack())
		}
	})
}

func BenchmarkPromiseSingle(b *testing.B) {
	bucket := promise.New(1000)
	zipf := MockZipf(b.N)

	for n := 0; n < b.N; n++ {
		request, _ := bucket.Fetch(MockFetch(zipf).Unpack())
		<-request.(*promise.Request).Value
	}
}

func BenchmarkPromiseMultiple(b *testing.B) {
	bucket := promise.New(1000)

	b.SetParallelism(N)
	b.RunParallel(func(pb *testing.PB) {
		zipf := MockZipf(b.N)

		for pb.Next() {
			request, _ := bucket.Fetch(MockFetch(zipf).Unpack())
			<-request.(*promise.Request).Value
		}
	})
}
