package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNum := rand.Intn(5)

	var num int
	for {
		fmt.Println("Enter the number : ")
		fmt.Scan(&num)
		if num == randomNum {
			fmt.Println("You have guessed it right")
			return
		} else if num > randomNum {
			fmt.Println("Guess a lesser number")
		} else {
			fmt.Println("Guess a larger number")
		}

	}

}
