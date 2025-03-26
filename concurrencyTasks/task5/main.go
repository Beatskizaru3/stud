package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	result := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 5; i <= 50; i += 5 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for v := range ch {
			wg.Add(2)

			go func(num int) {
				defer wg.Done()
				mult1 := num * num
				result <- mult1
			}(v)

			go func(num int) {
				defer wg.Done()
				mult2 := num * num * num
				result <- mult2
			}(v)
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for r := range result {
		fmt.Println(r)
	}

}
