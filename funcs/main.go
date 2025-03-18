package main

import "fmt"

// func multiplies(x1 int) func(x2 int) int {
// 	multipliesFunc := func(x2 int) int {
// 		return x1 * x2
// 	}
// 	return multipliesFunc
// }

// func devides(x1 int) func(x2 int) int {

// 	deviderFunc := func(x2 int) int {
// 		return x1 / x2
// 	} // функция которая возвращает функцию, при возврате может возврщаать результат уже функции

// 	return deviderFunc
// }

// func counter() func() int { //замыкание
// 	localCounter := 0
// 	return func() int {
// 		localCounter += 1
// 		return localCounter
// 	}
// }

func square(num *int) {
	*num = *num * *num
	return
}

func main() {
	// result := multiplies(5)(10)
	// fmt.Println(result)
	// result = devides(50)(5) // передаем во вторых скобках значение в возврщаему функцию
	// fmt.Println(result)

	// value := 40
	// getUahValue := func() int {
	// 	return value
	// }
	// fmt.Println(getUahValue())
	// value = 42
	// fmt.Println(getUahValue())
	// increase := counter()

	// fmt.Println(increase())
	// fmt.Println(increase())
	// fmt.Println(increase())

	//pointers
	a := 9
	square(&a)
	fmt.Println(a)
}
