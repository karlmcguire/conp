package mutex

import (
	"fmt"
	"testing"
)

func BenchmarkSimple(b *testing.B) {
	var (
		N      = 10
		done   = make(chan struct{}, N)
		router = NewRouter()

		bench = func() {
			for n := 0; n < b.N; n++ {
				request := router.Get(fmt.Sprintf("%d", n))
				request.Nop()
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
