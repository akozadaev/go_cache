package main

import (
	"custom_thread_safe"
	"fmt"
)

func main() {
	cache := custom_thread_safe.NewLRUCache[string, int](2)

	cache.Add("a", 1)
	cache.Add("b", 2)

	fmt.Println(cache.Get("a")) // 1 true

	cache.Add("c", 3) // "b" будет удален

	fmt.Println(cache.Get("b")) // 0 false
	fmt.Println(cache.Size())   // 2 Последовательно идущие единицы
}
