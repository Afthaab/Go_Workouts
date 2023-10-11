package main

import (
	"fmt"

	"converter/temperature"
)

func main() {
	var temp float64
	fmt.Println("Enter the Fahrenheit value : ")
	fmt.Scanln(&temp)
	fmt.Println("The Celsius Value is : ", temperature.TempratureConverter(temp))

}
