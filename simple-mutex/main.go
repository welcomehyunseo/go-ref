package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	name  string
	count int
}

func main() {
	counter1 := &Counter{
		name:  "Herry",
		count: 0,
	}
	counter2 := &Counter{
		name:  "Elex",
		count: 0,
	}
	for i := 0; i < 100; i++ {
		go handleCounter(counter1, 0)
		go handleCounter(counter2, 2)
	}
	for {
	}
}

func handleCounter(c *Counter, sec int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	old := c.count
	c.count++
	new := c.count
	name := c.name
	fmt.Printf("%s: %d -> %d\n", name, old, new)

	sleep := time.Duration(sec) * time.Second
	time.Sleep(sleep)
}
