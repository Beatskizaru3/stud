package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // счетчик задач
	var mu sync.RWMutex   // синхронизатор горутин деля их на чтение и запись
	var counter int

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.RLock() // блокируем RaceCondition

			_ = counter // Защищаем участок кода от гонки данных для чтения

			mu.RUnlock() // разблокируем
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
