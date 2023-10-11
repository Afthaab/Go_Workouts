package main

import "fmt"

func main() {
	numbers1 := []int{10, 30, 40, 60}
	numbers2 := []int{44, 65, 88, 777, 96}
	numbers2 = append(numbers2, numbers1...)
	fmt.Println("The concatenated slice : ", numbers2)
}
