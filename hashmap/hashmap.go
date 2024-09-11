package hashmap

import (
	"hash/fnv"
)

type HashMap struct {
	Buckets []Bucket
	Size    int
}

type Bucket struct {
	Pairs []KeyValuePair
}

type KeyValuePair struct {
	Key   string
	Value interface{}
}

func New(bucketCount int) *HashMap {
	buckets := make([]Bucket, bucketCount)
	return &HashMap{
		Buckets: buckets,
		Size:    bucketCount,
	}
}

// Hash function to convert a key into an index
func (hm *HashMap) hash(key string) int {
	h := fnv.New64a()
	h.Write([]byte(key))
	hash := h.Sum64()
	return int(hash % uint64(hm.Size))
}

// Put a key-value pair into the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	index := hm.hash(key)
	for i, kv := range hm.Buckets[index].Pairs {
		if kv.Key == key {
			hm.Buckets[index].Pairs[i].Value = value
			return
		}
	}
	hm.Buckets[index].Pairs = append(hm.Buckets[index].Pairs, KeyValuePair{Key: key, Value: value})
}

// Get a value by key from the hash map
func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.hash(key)
	for _, kv := range hm.Buckets[index].Pairs {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return nil, false
}

// Delete a key-value pair from the hash map
func (hm *HashMap) Delete(key string) {
	index := hm.hash(key)
	for i, kv := range hm.Buckets[index].Pairs {
		if kv.Key == key {
			hm.Buckets[index].Pairs = append(hm.Buckets[index].Pairs[:i], hm.Buckets[index].Pairs[i+1:]...)
			return
		}
	}
}
