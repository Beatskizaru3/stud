package main

import "fmt"

type Runner interface {
	Run() string
}

type Swimmer interface {
	Swim() string
}

type Driver interface {
	Drive() string
}

type Sportsman interface {
	Runner
	Swimmer
	Driver
}

type Human struct {
	Name     string
	Age      int
	Distance float32
}

func (h Human) Run() string {
	return fmt.Sprintf("Person %s, ran %dm", h.Name, h.Distance)
}

func (h Human) Swim() string {
	return fmt.Sprintf("Person %s, swam %dm", h.Name, h.Distance)
}

func (h Human) Drive() string {
	return fmt.Sprintf("Person %s, drived %dm", h.Name, h.Distance)
}

func main() {
	var r Runner
	fmt.Printf("type: %T, value: %v\n", r, r)

	var unnamedSportsman *Human
	fmt.Printf("type: %T, value: %v\n", unnamedSportsman, unnamedSportsman)

	r = unnamedSportsman
	fmt.Printf("type: %T, value: %v\n", r, r)

	r = &Human{Name: "Seva", Age: 22, Distance: 100}
	fmt.Printf("Type: %T, Value: %v\n", r, r)
}
