package main

import (
	"fmt"
)

func main() {
	var numbers = []int{-1, -22, -10005, -140, -150}

	num := numbers[1]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > num {
			num = numbers[i]
		}
	}

	fmt.Println(num)

}
