package main

import "fmt"

func main() {
	var numbers = []int{}
	numbers = append(numbers, 5)
	fmt.Println(numbers)
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)

	fmt.Println("The length of the slice : ", len(numbers))
	fmt.Println("The Capacity of the slice : ", cap(numbers))

	numbers = deletetion(numbers, 2)

	fmt.Println(numbers)
}

func deletetion(numbers []int, index int) []int {
	numbers = append(numbers[:index], numbers[index+1:]...)
	return numbers
}
