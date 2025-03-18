package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type BrickBuilder struct {
	Person
	Name string
}

func (p Person) ShowInfo() {
	fmt.Println(p)
}

func main() {
	// seva := Person{
	// 	Age:  22,
	// 	Name: "Seva",
	// }
	// user1 := &seva
	// fmt.Println(*user1)
	// fmt.Println((*user1).Age)
	// fmt.Println(user1.Age)

	// ivan := &Person{
	// 	Age:  22,
	// 	Name: "Ivan",
	// }
	// fmt.Println(ivan)
	worker1 := BrickBuilder{Person{Name: "Seva", Age: 22}, "Alex"}
	//equal
	worker2 := BrickBuilder{
		Person: Person{
			Name: "Marat",
			Age:  25,
		},
		Name: "Bob",
	}

	fmt.Println(worker1)            //{{22 Seva} Alex}
	fmt.Println(worker1.Age)        //22
	fmt.Println(worker1.Person.Age) //22

	worker1.ShowInfo()        //{22 Seva}
	worker1.Person.ShowInfo() //{22 Seva}

	worker2.ShowInfo()        //{25 Marat}
	fmt.Println(worker2.Name) //Bob

}
