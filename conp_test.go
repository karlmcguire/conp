package conp

import (
	"fmt"

	"github.com/xba/stress"
)

const (
	MULTIPLE = 10
	// use the same seed accross all zipf distributions
	SEED = 23459283745
)

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
		Value: func() interface{} { return "datadatadata" },
	}
}
