package main

import "fmt"

func main() {
	var fruits = []string{
		"Apple",
		"Dates",
		"Watermelon",
		"Orange",
		"Grapes",
	}
	fmt.Println("The fruits are : ")
	for _, v := range fruits {
		fmt.Println(v)
	}
}
