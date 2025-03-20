package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 10), make(chan int, 10)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctxCancel, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i <= 1000; i++ {
			if i == 500 {
				cancel()
			}
			numbersToProcess <- i
		}
		cancel()
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}
	fmt.Println(counter)
}

func worker(ctx context.Context, processingTasks <-chan int, processedTasks chan<- int) {
	for {
		select {
		case <-ctx.Done(): // блокируещияся операция
			return

		case value, ok := <-processingTasks:
			if !ok {
				return // если канал закрыт, выходим из воркера
			}
			time.Sleep(time.Millisecond)
			processedTasks <- value * value
		}
	}
}
