package main

import "fmt"

func main() {
	num := 0
	sliceA := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Enter the number to searched")
	fmt.Scan(&num)
	for i := 0; i < len(sliceA); i++ {
		if num == sliceA[i] {
			fmt.Printf("The element is found in %d of the slice", i)
			return
		}
	}
	fmt.Println("The number is not found")
}
