package main

import (
	"fmt"
)

func main() {
	var numbers = []int{100, 20, 10005, 40, 50}

	num := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > num {
			num = numbers[i]
		}
	}

	fmt.Println(num)

}
