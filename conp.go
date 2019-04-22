package conp

import (
	"crypto/rand"
	"fmt"

	"github.com/minio/highwayhash"
)

type (
	Request struct {
		Key   string
		Value chan string
	}

	Bucket struct {
		Id       int
		Data     map[uint64]string
		Requests chan *Request
		Stop     chan struct{}
	}

	Router struct {
		Buckets []*Bucket
		Count   int
		Seed    []byte
	}
)

func (b *Bucket) Start() {
	for {
		select {
		case request := <-b.Requests:
			request.Value <- fmt.Sprintf("%d", b.Id)
		case <-b.Stop:
			return
		}
	}
}

func NewRouter(count, buffer int) *Router {
	var (
		router = &Router{
			Count: count,
			Seed:  make([]byte, 32, 32),
		}
	)

	// create 32 byte seed for highwayhash
	rand.Read(router.Seed)

	for i := 0; i < count; i++ {
		router.Buckets = append(router.Buckets, &Bucket{
			Id:       i,
			Data:     make(map[uint64]string),
			Requests: make(chan *Request, buffer),
			Stop:     make(chan struct{}),
		})

		// start listening for requests
		go router.Buckets[i].Start()
	}

	return router
}

func (r *Router) Get(key string) *Request {
	var (
		// hash key to get uint64
		hashed = highwayhash.Sum64([]byte(key), r.Seed)
		// create request with promised value
		request = &Request{
			Key:   key,
			Value: make(chan string, 1),
		}
	)

	// add requests to queue
	r.Buckets[hashed&(uint64(r.Count)-1)].Requests <- request

	return request
}
