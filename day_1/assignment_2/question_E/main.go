package main

import "fmt"

// e. Create a group of constants using  const ()  group constant declaration syntax, assign it some values & print it.
// Try to change its values & check if there is some error.

func main() {
	const (
		name       string  = "Afthab"
		age        int     = 22
		percentage float64 = 97.99
	)

	// tried changing the value of name, but getting an error

	// name = "suhail"
	// cannot assign to name (constant "Afthab" of type string)

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("Percentage : ", percentage)

}
