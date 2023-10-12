package main

import "fmt"

func main() {
	fruits := make(map[string]int)
	fruits["Mango"] = 100
	fruits["Apple"] = 20
	fruits["Chicko"] = 75

	for fruit, v := range fruits {
		fmt.Printf("The Fruit is '%d' and the quantity is '%d' \n", fruit, v)
	}
}
