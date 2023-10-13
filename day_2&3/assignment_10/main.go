package main

import "fmt"

func main() {
	fmt.Println("The sum of numbers : ", CalculateSum(10, 50, 66, 1))
}

func CalculateSum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
