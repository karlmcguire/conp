package simple

import (
	"crypto/rand"
	"sync"

	"github.com/minio/highwayhash"
)

type Bucket struct {
	sync.Mutex

	data map[uint64]interface{}
	seed []byte
}

func New() *Bucket {
	bucket := &Bucket{
		data: make(map[uint64]interface{}),
		seed: make([]byte, 32, 32),
	}

	// highwayhash seed
	rand.Read(bucket.seed)

	return bucket
}

func (b *Bucket) Fetch(k string, f func() interface{}) (interface{}, bool) {
	b.Lock()
	defer b.Unlock()

	// if exists, get and return
	iden := highwayhash.Sum64([]byte(k), b.seed)
	if elem, exists := b.data[iden]; exists {
		return elem, true
	}

	// add and return
	b.data[iden] = f()
	return b.data[iden], false
}
