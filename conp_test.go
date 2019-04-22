package conp

import (
	"fmt"

	"github.com/xba/stress"
)

// zipf seed (use the same across all goroutines, need to copy because the
// generator isn't concurrent-safe)
const SEED = 23459283745

func MockZipf(n int) *stress.Trace {
	return stress.NewSimulator(stress.GenerateZipf(1.1, 1, uint64(n), SEED))
}

func MockFetch(t *stress.Trace) *Fetcher {
	n, err := t.Next()
	if err != nil {
		panic(err)
	}

	return &Fetcher{
		Key:   fmt.Sprintf("%d", n),
		Value: func() interface{} { return nil },
	}
}
