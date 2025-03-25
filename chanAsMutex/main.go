package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "golang.org/x/sync/errgroup"
)

func makeRequest(x int) <-chan string {
	responseChan := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		responseChan <- fmt.Sprintf("Запрос №%d был отправлен\n", x)
	}()
	return responseChan
}

func chanAsPromise() {
	firstResponseChan := makeRequest(1)
	secondResponseChan := makeRequest(2)

	fmt.Println(<-firstResponseChan, <-secondResponseChan)
}

func chanAsMutex() {
	var counter int
	muxtexChan := make(chan struct{}, 1) // пустая структура ничего не весит, самые легковесные данные
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			muxtexChan <- struct{}{}
			counter++
			<-muxtexChan
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func withOutErrGroup() {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("First started")
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("second started")
			err = fmt.Errorf("any error")

			cancel()
		}
	}()

	go func() {
		defer wg.Done()

		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("third started")
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	fmt.Println(err)
}

func main() {
	//chanAsPromise()
	//chanAsMutex()
	withOutErrGroup()
}
