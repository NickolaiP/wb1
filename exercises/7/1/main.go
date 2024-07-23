package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Counters struct {
	mx sync.Mutex
	m  map[int]string
}

func NewCounters() *Counters {
	return &Counters{
		m: make(map[int]string),
	}
}

func (c *Counters) Load(key int) (string, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Counters) Store(key int, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func main() {
	counters := NewCounters()
	for i := 0; i < 10; i++ {
		go func(c *Counters) {
			for {
				c.Store(rand.Intn(1000), "WB top")
				time.Sleep(10 * time.Millisecond)
			}

		}(counters)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(counters.m)
}
