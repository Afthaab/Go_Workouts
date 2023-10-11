package main

import "fmt"

func main() {
	fmt.Println("The array values:")
	var arr = [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
