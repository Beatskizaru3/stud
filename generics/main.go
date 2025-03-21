package main

import (
	"fmt"
)

func sum[V int64 | float64](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if searchEl == el {
			return true
		}
	}
	return false
}

func showContains() {
	type Person struct {
		name     string
		age      int64
		jobTitle string
	}

	ints := []int64{1, 2, 3, 4, 5}
	fmt.Println("int:", contains(ints, 4))

	stirngs := []string{"Vasya", "Katya", "Seva"}
	fmt.Println("name:", contains(stirngs, "Seva"))
	fmt.Println("name:", contains(stirngs, "Katya"))
}

func show[T any](enteties ...T) {
	fmt.Println(enteties)
}

func showAny() {
	show(1, 2, 3)
	show("test", "test1", "test2")
	show([]int64{1, 2, 3}, []int64{4, 5, 6})
	show(map[string]int64{
		"number1": 20,
		"number2": 30,
	})

	show(interface{}(1), interface{}("string"), any(struct{ name string }{name: "Art"}))
}

type Number interface {
	int64 | float64
}

type Numbers[T Number] []T

func unionInterfaceAndType() {
	var ints Numbers[int64]

	ints = append(ints, []int64{1, 5, 44, 33, 22}...) // ... передает элементы массива как отдельные аргументы

	floats := Numbers[float64]{2.22, 3.114, 5.665, 3.77}
	fmt.Println(ints)
	fmt.Println(floats)
}

func sumUnionInterface[V Number](numbers []V) V {
	var sum V
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	floats := []float64{1.2, 3.2, 5.4}
	ints := []int64{1, 3, 5}

	fmt.Println(sum(floats))
	fmt.Println(sum(ints))

	showContains()

	showAny()

	fmt.Println(sumUnionInterface(ints))

	unionInterfaceAndType()
}
