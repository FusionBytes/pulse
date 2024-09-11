package main

import (
	"fmt"
	"pulse/hashmap"
)

func main() {
	hm := hashmap.New(10)

	hm.Put("name", "Alice")
	hm.Put("age", 30)
	hm.Put("city", "Wonderland")

	if value, found := hm.Get("name"); found {
		fmt.Println("Name:", value)
	} else {
		fmt.Println("Name not found")
	}

	hm.Delete("age")

	if value, found := hm.Get("age"); found {
		fmt.Println("Age:", value)
	} else {
		fmt.Println("Age not found")
	}
}
