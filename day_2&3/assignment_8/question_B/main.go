package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))

}

func Factorial(num int) int {
	if num < 1 {
		return 1
	}
	return num * Factorial(num-1)
}
