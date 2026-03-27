package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}

	pool.Put("Alief")
	pool.Put("Sufi")
	pool.Put("Uzan")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
}
