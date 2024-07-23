package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var counters sync.Map
	for i := 0; i < 10; i++ {
		go func(c *sync.Map) {
			for {
				c.Store(rand.Intn(1000), "WB top")
				time.Sleep(10 * time.Millisecond)
			}

		}(&counters)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(counters)
}
