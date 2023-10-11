package main

import "fmt"

func main() {
	sum := 0
	var numbers = []int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	fmt.Println("The sum of the numbers : ", sum)
	fmt.Println("The averegare of the numbers : ", sum/len(numbers))

}
