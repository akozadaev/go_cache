package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

// go get github.com/patrickmn/go-cache
func main() {
	c := cache.New(5*time.Minute, 10*time.Minute) // TTL = 5 мин, cleanup каждые 10 мин

	c.Set("key", "value", cache.DefaultExpiration)
	if val, found := c.Get("key"); found {
		fmt.Println(val) // value
	}

	time.Sleep(6 * time.Minute)
	if _, found := c.Get("key"); !found {
		fmt.Println("expired") // будет напечатано
	}
}
