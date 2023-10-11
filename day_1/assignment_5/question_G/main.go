package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	fmt.Println(numbers)
}
