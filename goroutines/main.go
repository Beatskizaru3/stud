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

			mu.RLock() // блокируем только чтение

			_ = counter // Защищаем участок кода от гонки данных для чтения

			mu.RUnlock() // разблокируем чтение
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // для записи тоже самое
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
