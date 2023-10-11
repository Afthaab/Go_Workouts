package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("The number is EVEN")
	} else {
		fmt.Println("The number is ODD")
	}
}
