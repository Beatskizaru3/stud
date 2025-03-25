package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const numberOfLinks int = 1000000

	var codes sync.Map               // мьютекс для мап
	workersCount := runtime.NumCPU() // оптимальное кол-во воркеров

	ch := make(chan string, workersCount*2) // буфферизированный канал

	wg := sync.WaitGroup{}
	for i := 0; i < workersCount; i++ { // запускаем 5 воркеров
		wg.Add(1)
		go func() { // запускаем код паралельно
			defer wg.Done()
			for url := range ch { //бесконечно перебираем канал, до его
				code := sendRequest(url)
				val, _ := codes.LoadOrStore(code, 0) // добавляет ключ и передает значение 0 в мапу
				codes.Store(code, val.(int)+1)       // пишем на ключ значение
			}
		}()
	}

	go func() { // фрагмент кода, который выполняется паралелльно с основным потоком
		for i := 0; i < numberOfLinks; i++ {
			ch <- fmt.Sprintf("https://links%d", i)
		}
		time.Sleep(time.Second)
		close(ch) // закрываем канал, тем самым завершаем горутины, которые читают
	}()

	wg.Wait()

	codes.Range(func(key, value any) bool {
		fmt.Printf("HTTP код: %v, Количество: %v\n", key, value)
		return true
	})
}

func sendRequest(url string) int {
	statusCodes := []int{200, 404, 503}
	randomIndex := rand.Intn(len(statusCodes))
	return statusCodes[randomIndex]
}
