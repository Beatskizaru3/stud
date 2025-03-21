package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func addToAtomic() {
	start := time.Now()
	var counter int64
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("With atomic:", time.Now().Sub(start).Seconds())
}

func main() {
	addToAtomic()
}
