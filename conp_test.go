package conp

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	router := NewRouter(8, 10)

	for n := 0; n < b.N; n++ {
		request := router.Get(fmt.Sprintf("%d", n))
		<-request.Value
	}
}
