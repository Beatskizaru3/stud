package main

import "fmt"

func main() {
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr1)

	arr2 := []int{1, 2, 3, 4, 5, 6, 7}

	for indx, values := range arr1 {
		fmt.Printf("Index: %d Values: %s\n", indx, values)
	}

	changeInArr(&arr1)

	for indx, values := range arr1 {
		fmt.Printf("Index: %d Values: %s\n", indx, values)
	}

	showAllElements(1, 2, 3, 4, 5, 6, 7)
	showAllElements(arr2...)
}

func showAllElements(values ...int) {
	fmt.Println(values)
}

func changeInArr(arr1 *[5]string) *[5]string {
	arr1[2] = "sss"
	return arr1
}
