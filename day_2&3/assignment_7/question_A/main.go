package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	City string
}

func main() {
	Employee1 := Employee{
		Name: "Mohammed Afthab",
		Age:  22,
		City: "Bangalore",
	}
	Employee2 := Employee{
		Name: "Poorvi",
		Age:  22,
		City: "Bangalore",
	}
	Employee3 := Employee{
		Name: "Jeevan",
		Age:  22,
		City: "Bangalore",
	}

	fmt.Println("The Structs are : ")
	fmt.Println(Employee1)
	fmt.Println(Employee2)
	fmt.Println(Employee3)

}
