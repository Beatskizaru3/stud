package main

import (
	"fmt"
	"sync"
)

func main() {
	bufChan := make(chan int, 10)
	wg := sync.WaitGroup{}
	go func() {
		for i := 1; i < 20; i++ {
			bufChan <- i
		}
		close(bufChan)
	}()

	for i := range bufChan {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Val1 ", i*i)
		}(i)

		go func(i int) {
			defer wg.Done()
			fmt.Println("Val2 ", i*i*i)
		}(i)
	}
	wg.Wait()
}
