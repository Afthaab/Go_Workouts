package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := []int{10, 20, 30, 40}

	if ok := isSliceEqual(slice1, slice2); !ok {
		fmt.Println("The Slices are not Equal")
		return
	}

	fmt.Println("The Slices are equal")

}

func isSliceEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
