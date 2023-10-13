package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Person struct {
	Name string
	Age  int
	Address
}

func main() {
	p := Person{
		Name: "Mohammed Afthab",
		Age:  22,
		Address: Address{
			Street:  "BTM Layout",
			City:    "Bengaluru",
			ZipCode: "560076",
		},
	}
	p1 := Person{
		Name: "Jeevan",
		Age:  22,
		Address: Address{
			Street:  "BTM Layout",
			City:    "Bengaluru",
			ZipCode: "560076",
		},
	}
	fmt.Println("The Details of Pesron : ", p)
	fmt.Println("The Details of Pesron : ", p1)

}
