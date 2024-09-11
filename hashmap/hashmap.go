package hashmap

import (
	"fmt"
	"hash/fnv"
)

const loadFactor = 0.7

type HashMap struct {
	Buckets         []Bucket
	size            int
	bucketSizeLimit int
}

type Bucket struct {
	Pairs       []KeyValuePair
	OverFlowMap *HashMap
}

type KeyValuePair struct {
	Key   string
	Value interface{}
}

func New(bucketCount int, becketSizeLimit int) *HashMap {
	buckets := make([]Bucket, bucketCount)
	return &HashMap{
		Buckets:         buckets,
		size:            bucketCount,
		bucketSizeLimit: becketSizeLimit,
	}
}

// Hash function to convert a key into an index
func (hm *HashMap) hash(key string) int {
	h := fnv.New64a()
	h.Write([]byte(key))
	hash := h.Sum64()
	return int(hash % uint64(hm.size))
}

// Put a key-value pair into the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	index := hm.hash(key)
	if float32(len(hm.Buckets[index].Pairs))/float32(hm.bucketSizeLimit) >= loadFactor {
		if hm.Buckets[index].OverFlowMap == nil {
			hm.Buckets[index].OverFlowMap = New(hm.size, hm.bucketSizeLimit)
		}
		hm.Buckets[index].OverFlowMap.Put(key, value)
		return
	}
	for i, kv := range hm.Buckets[index].Pairs {
		if kv.Key == key {
			hm.Buckets[index].Pairs[i].Value = value
			return
		}
	}
	if hm.Buckets[index].Pairs == nil {
		hm.Buckets[index].Pairs = make([]KeyValuePair, 0, hm.bucketSizeLimit)
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
	if hm.Buckets[index].OverFlowMap != nil {
		return hm.Buckets[index].OverFlowMap.Get(key)
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
	if hm.Buckets[index].OverFlowMap != nil {
		hm.Buckets[index].OverFlowMap.Delete(key)
	}
}

func (hm *HashMap) PrintBuckets() {
	for i, bucket := range hm.Buckets {
		fmt.Printf("Bucket %d:\n", i)
		printBucketPairs(bucket, 0)

		// Print overflow maps if they exist
		overflow := bucket.OverFlowMap
		level := 1
		for overflow != nil {
			fmt.Printf("  Overflow Bucket Level %d:\n", level)
			for j, overflowBucket := range overflow.Buckets {
				fmt.Printf("    Overflow Bucket %d-%d:\n", i, j)
				printBucketPairs(overflowBucket, level)
			}
			overflow = overflow.Buckets[i].OverFlowMap
			level++
		}
		fmt.Println()
	}
}

func printBucketPairs(bucket Bucket, level int) {
	if len(bucket.Pairs) == 0 {
		fmt.Printf("  [Bucket Empty]\n")
		return
	}
	for _, kv := range bucket.Pairs {
		fmt.Printf("  Key: %s, Value: %v (Level %d)\n", kv.Key, kv.Value, level)
	}
}
