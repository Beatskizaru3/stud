package main

import "fmt"

type kvadrat struct {
	Size int
}

func (k kvadrat) Perimetr() {
	k.Size = k.Size * 4
}

func (k *kvadrat) Scale() {
	k.Size = k.Size * k.Size
}
func main() {

	square1 := kvadrat{Size: 4}
	//square1link := &square1

	square2 := kvadrat{Size: 16}
	//square2link := &square2
	square1.Perimetr()
	square2.Perimetr()
	fmt.Println(square1)
	fmt.Println(square2)

	square3 := kvadrat{Size: 5}
	sqaure3link := &square3
	sqaure3link.Scale()
	fmt.Println(square3.Size)
}
