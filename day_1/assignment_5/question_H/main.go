package main

import "fmt"

func main() {
	evenSum := 0
	oddSum := 0
	numbers := []int{10, 11, 12, 13, 15, 14, 999, 888, 969}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenSum++
		} else {
			oddSum++
		}
	}
	fmt.Println("The Odd numbers are : ", oddSum)
	fmt.Println("The even numbers are : ", evenSum)
}
