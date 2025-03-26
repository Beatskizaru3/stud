package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello1")
	}()
	go func() {
		fmt.Println("Hello2")
	}()

	time.Sleep(time.Second)
}
