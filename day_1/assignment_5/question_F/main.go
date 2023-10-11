package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("The slice is : ")
	fmt.Println(numbers)

	num := 0
	fmt.Println("Enter the number to be searched for : ")
	fmt.Scan(&num)

	for i := 0; i < len(numbers); i++ {
		if num == numbers[i] {
			fmt.Printf("The element is found at %d of the slice", i)
			return
		}
	}
	fmt.Println("The element is not found in the slice")
}
