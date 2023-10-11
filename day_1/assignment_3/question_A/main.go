package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {

	fmt.Println("The Square of the number : ", calculator.SquareOfANumber(10))
	fmt.Println("The Sum of the numbers : ", calculator.SumOfNumbers(10, 20))
	fmt.Println("The Difference of the numbers : ", calculator.DifferenceOfNumbers(20, 10))
	fmt.Println("The Product of the numbers : ", calculator.ProductOfNumbers(10, 20))
	quo, rem := calculator.DivisionOfNumbers(60, 3)
	fmt.Println("The Quotientof of the number : ", quo)
	fmt.Println("The Remainder of the number : ", rem)
}
