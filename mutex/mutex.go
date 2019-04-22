package mutex

import (
	"crypto/rand"
	"sync"

	"github.com/minio/highwayhash"
)

type Router struct {
	Mutex sync.Mutex
	Seed  []byte
}

type Request struct {
	Key   string
	Value uint64
}

func (r *Request) Nop() {}

func NewRouter() *Router {
	router := &Router{
		Seed: make([]byte, 32, 32),
	}

	rand.Read(router.Seed)

	return router
}

func (r *Router) Get(key string) *Request {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	return &Request{
		Key:   key,
		Value: highwayhash.Sum64([]byte(key), r.Seed),
	}
}
