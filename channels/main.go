package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// bufferedChan := make(chan string, 3)
	// //

	// bufferedChan <- "first"

	// select {
	// case str := <-bufferedChan:
	// 	fmt.Println("read", str)
	// case bufferedChan <- "second":
	// 	fmt.Println("write", <-bufferedChan, <-bufferedChan)
	// }

	// unBufferedChan := make(chan int)

	// go func() {
	// 	time.Sleep(time.Second)
	// 	unBufferedChan <- 1
	// }()

	// select {
	// case bufferedChan <- "third": // выполниться первым потому, что не блокирующиеся операция
	// 	fmt.Println("unblocking writing")
	// case val := <-unBufferedChan: // блокируемся в селекте и ждет, пока не появятся данные
	// 	fmt.Println("blocking reading", val) // но это не отменяет того что могут выполниться за это время другие кейсы
	// case <-time.After(time.Millisecond * 500): // удобно для тайм-аута
	// 	fmt.Println("Time is up") // логика при тайм-ауте
	// default:
	// 	fmt.Println("default case") // выпонлиться если нет неблокирующихся операций
	// }

	// resultChan := make(chan int)
	// timer := time.After(time.Second)

	// go func() {
	// 	defer close(resultChan)

	// 	for i := 0; i < 1000; i++ {

	// 		select {
	// 		case <-timer:
	// 			fmt.Println("times's up")
	// 			return
	// 		default:
	// 			resultChan <- i
	// 		}
	// 	}

	// }()
	// for val := range resultChan {
	// 	fmt.Println(val)
	// }

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(5 * time.Second)

	select {
	case <-timer:
		fmt.Println("time's up")
		return

	case sig := <-sigChan:
		fmt.Println("stopped by signal ", sig)
		return
	}
}
