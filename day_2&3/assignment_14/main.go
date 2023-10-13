package main

import (
	"fmt"
	"log"
)

func Divide(numerator int, denominator int) int {
	if denominator == 0 {
		panic("the denominator is zero; cannot proceed")
	}
	return numerator / denominator
}

func SafeDivide(numerator int, denominator int) {
	defer recoveryFunction()
	fmt.Println("the answer is : ", Divide(numerator, denominator))
}

func recoveryFunction() {
	msg := recover()
	if msg != nil {
		log.Println(msg)
	}
	// fmt.Println(string(debug.Stack())) to print the debug stack
}

func main() {
	SafeDivide(0, 10)
	fmt.Println("The main landed safetly")
}
