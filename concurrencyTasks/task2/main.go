package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 20; i += 2 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := range ch {

			fmt.Println("Value", i)

		}
	}()

	time.Sleep(time.Second)
	fmt.Println(ch)
}
