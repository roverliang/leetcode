package main

import (
	"fmt"
	"github.com/roverliang/leetcode/algorithm/lru"
)

func main() {

	cache := lru.New(5)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	cache.Put(5, 5)

	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	cache.Put(6, 6)
	fmt.Println(cache.Get(5))

}
