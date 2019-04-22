package promise

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	var (
		N      = 10
		router = NewRouter(10, 10)
		done   = make(chan struct{}, N)

		bench = func() {
			for n := 0; n < b.N; n++ {
				request := router.Get(fmt.Sprintf("%d", n))
				<-request.Value
			}

			done <- struct{}{}
		}
	)

	for i := 0; i < N; i++ {
		go bench()
	}

	for i := 0; i < N; i++ {
		<-done
	}
}
