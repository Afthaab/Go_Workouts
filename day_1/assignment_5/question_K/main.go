package main

import "fmt"

func main() {
	numbers := []int{10, 10, 30, 40, 60, 77, 10}
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				numbers = deletetion(numbers, j)
			}
		}

	}
	fmt.Println("The resultant slice is : ", numbers)

}

func deletetion(numbers []int, index int) []int {
	numbers = append(numbers[:index], numbers[index+1:]...)
	return numbers
}
