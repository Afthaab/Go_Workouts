package main

import "fmt"

func main() {
	evenSum := 0
	numbers := []int{10, 11, 10, 13, 15, 10, 999, 20, 969}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenSum = evenSum + numbers[i]
		}
	}
	fmt.Println("The sum of even numbers are : ", evenSum)
}
