package promise

import (
	"crypto/rand"

	"github.com/minio/highwayhash"
)

type Request struct {
	Key   string
	Value chan interface{}
	Fetch func() interface{}
}

type Bucket struct {
	buff chan *Request
	data map[uint64]interface{}
	seed []byte
}

func New(n int) *Bucket {
	bucket := &Bucket{
		buff: make(chan *Request, n),
		data: make(map[uint64]interface{}),
		seed: make([]byte, 32, 32),
	}

	// highwayhash seed
	rand.Read(bucket.seed)

	// start listening for incoming requests
	go bucket.Start()

	return bucket
}

func (b *Bucket) Start() {
	for {
		var (
			request = <-b.buff
			iden    = highwayhash.Sum64([]byte(request.Key), b.seed)
		)

		if value, exists := b.data[iden]; exists {
			request.Value <- value
		} else {
			b.data[iden] = request.Fetch()
			request.Value <- b.data[iden]
		}
	}
}

func (b *Bucket) Fetch(k string, f func() interface{}) (interface{}, bool) {
	request := &Request{
		Key:   k,
		Value: make(chan interface{}, 1),
		Fetch: f,
	}

	b.buff <- request

	return request, true
}
