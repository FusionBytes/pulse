package main

import (
	"fmt"
	"pulse/hashmap"
)

func main() {
	bucketCount := 3
	bucketSizeLimit := 8
	hm := hashmap.New(bucketCount, bucketSizeLimit)

	for i := 0; i < bucketCount*bucketSizeLimit; i++ {
		key := fmt.Sprintf("key-%d", i)
		value := fmt.Sprintf("value-%d", i)
		hm.Put(key, value)
	}

	hm.PrintBuckets()
	fmt.Println(hm.Get("key-20"))
	hm.Delete("key-20")
	fmt.Println(hm.Get("key-20"))
}
