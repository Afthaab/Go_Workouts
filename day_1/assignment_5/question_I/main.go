package main

import "fmt"

func main() {
	numbers := []int{10, 15, 16, 14, 55, 99}
	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * 2
	}
	fmt.Println("The updated slice is : ", numbers)
}
