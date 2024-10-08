package structure

import "hash/fnv"

type entry struct {
	key   string
	value interface{}
}

type bucket struct {
	slots [8]*entry
	count int
}

type HashTable struct {
	bucketCount uint64
	buckets     []*bucket
	overFlow    *HashTable
	entryCount  int // Number of elements in the table
	loadFactor  float64
}

func NewHashTable(bucketCount uint64, loadFactor float64) *HashTable {
	if bucketCount < 2 {
		bucketCount = 2
	}
	return &HashTable{
		bucketCount: bucketCount,
		buckets:     make([]*bucket, bucketCount),
		overFlow:    nil,
		loadFactor:  loadFactor,
	}
}

func (h *HashTable) getBucketIndex(key string) (uint64, error) {
	hf := fnv.New64a()
	_, err := hf.Write([]byte(key))
	if err != nil {
		return 0, err
	}
	hash := hf.Sum64()

	return hash % h.bucketCount, nil
}

func (h *HashTable) getLoadFactor() float64 {
	return float64(h.entryCount) / float64(h.bucketCount*8)
}

func (h *HashTable) Insert(key string, value interface{}) error {
	// Resize the table if load factor exceeds the threshold
	if h.getLoadFactor() > h.loadFactor {
		h.resize()
	}

	bucketIndex, err := h.getBucketIndex(key)
	if err != nil {
		return err
	}

	nominatedBucket := h.buckets[bucketIndex]

	if nominatedBucket == nil {
		slots := [8]*entry{}
		nominatedBucket = &bucket{
			slots: slots,
			count: 0,
		}
		h.buckets[bucketIndex] = nominatedBucket
	}
	for i := 0; i < len(nominatedBucket.slots); i++ {
		if nominatedBucket.slots[i] == nil {
			// Empty slot found, insert the new entry
			nominatedBucket.slots[i] = &entry{key: key, value: value}
			nominatedBucket.count++
			h.entryCount++
			return nil
		} else if nominatedBucket.slots[i].key == key {
			// Key already exists, update the value
			nominatedBucket.slots[i].value = value
			return nil
		}
	}
	if h.overFlow == nil {
		h.overFlow = &HashTable{
			bucketCount: h.bucketCount,
			buckets:     make([]*bucket, h.bucketCount),
			loadFactor:  h.loadFactor,
		}
	}

	return h.overFlow.Insert(key, value)
}

func (h *HashTable) Get(key string) (interface{}, bool) {
	bucketIndex, err := h.getBucketIndex(key)
	if err != nil {
		return nil, false
	}

	nominatedBucket := h.buckets[bucketIndex]

	if nominatedBucket != nil {
		for _, slot := range nominatedBucket.slots {
			if slot != nil && slot.key == key {
				return slot.value, true
			}
		}
	}

	if h.overFlow != nil {
		return h.overFlow.Get(key)
	}
	return nil, false
}

func (h *HashTable) EntryCount() int {
	return h.entryCount
}

func (h *HashTable) resize() {
	newBucketCount := h.bucketCount * 2
	newTable := NewHashTable(newBucketCount, h.loadFactor)

	// Rehash all existing entries and insert into the new table
	for _, bucket := range h.buckets {
		if bucket == nil {
			continue
		}
		for _, slot := range bucket.slots {
			if slot != nil {
				_ = newTable.Insert(slot.key, slot.value)
			}
		}
	}

	// Handle overflow elements
	if h.overFlow != nil {
		for _, bucket := range h.overFlow.buckets {
			if bucket == nil {
				continue
			}
			for _, slot := range bucket.slots {
				if slot != nil {
					_ = newTable.Insert(slot.key, slot.value)
				}
			}
		}
	}

	// Replace current table with the resized table
	h.bucketCount = newBucketCount
	h.buckets = newTable.buckets
	h.entryCount = newTable.entryCount
	h.overFlow = newTable.overFlow
}
