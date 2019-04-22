package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/xba/stress"
)

const (
	// number of goroutines to use
	N = 100
	// zipf seed
	SEED = 245250
)

func Once(f func()) func(*testing.B) {
	return func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			f()
		}
	}
}

func Many(f func()) func(*testing.B) {
	return func(b *testing.B) {
		wg := sync.WaitGroup{}
		for i := 0; i < N; i++ {
			wg.Add(1)
			go func() {
				for n := 0; n < b.N; n++ {
					f()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func Zipf() *stress.Trace {
	return stress.NewSimulator(stress.GenerateZipf(
		1.1,
		1,
		0xffffffffffffffff,
		SEED))
}

func Iden(k interface{}, err error) int {
	if err != nil {
		panic(err)
	}

	return int(k.(uint64))
}

func main() {
	// zipf
	{
		zipf := Zipf()

		fmt.Println("zipf_create", testing.Benchmark(Once(func() {
			Zipf()
		})))

		fmt.Println("zipf_next", testing.Benchmark(Once(func() {
			zipf.Next()
		})))
	}

	// default map
	{
		data := make(map[int]interface{})
		zipf := Zipf()

		fmt.Println("map_default_get", testing.Benchmark(Once(func() {
			_ = data[Iden(zipf.Next())]
		})))

		fmt.Println("map_default_put", testing.Benchmark(Once(func() {
			data[Iden(zipf.Next())] = nil
		})))
	}

	// mutex map
	{
		data := make(map[int]interface{})
		mute := sync.Mutex{}
		zipf := Zipf()

		fmt.Println("map_mutex_get", testing.Benchmark(Many(func() {
			mute.Lock()
			_ = data[Iden(zipf.Next())]
			mute.Unlock()
		})))

		fmt.Println("map_mutex_put", testing.Benchmark(Many(func() {
			mute.Lock()
			data[Iden(zipf.Next())] = nil
			mute.Unlock()
		})))
	}
}
