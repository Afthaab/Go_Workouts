package main

import "fmt"

// d. Create a group of variables using var () variable declaration syntax, assign it some values & print it.

func main() {
	var (
		name       string
		age        int
		percentage float64
	)

	name = "Afthab"
	age = 22
	percentage = 97.9

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("percentage : ", percentage)

}
