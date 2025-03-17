package main

import "fmt"

func multiplies(x1 int) func(x2 int) int {
	multipliesFunc := func(x2 int) int {
		return x1 * x2
	}
	return multipliesFunc
}

func devides(x1 int) func(x2 int) int {

	deviderFunc := func(x2 int) int {
		return x1 / x2
	} // функция которая возвращает функцию, при возврате может возврщаать результат уже функции

	return deviderFunc
}

func main() {
	result := multiplies(5)(10)
	fmt.Println(result)
	result = devides(50)(5) // передаем во вторых скобках значение в возврщаему функцию
	fmt.Println(result)
}
