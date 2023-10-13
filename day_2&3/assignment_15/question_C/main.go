package main

import (
	"fmt"
	"time"
)

func sender1(c1 chan int) {
	for i := 1; i <= 5; i++ {
		c1 <- i
	}
	close(c1)
}

func sender2(c2 chan int) {
	for i := 6; i <= 10; i++ {
		c2 <- i
	}
	close(c2)
}

func sender3(c3 chan int) {
	for i := 11; i <= 15; i++ {
		c3 <- i
	}
	close(c3)
}

func receiver1(c1 chan int) {
	for num := range c1 {
		fmt.Printf("Received: %d\n", num)
	}
}

func receiver2(c2 chan int, c3 chan int) {
	for num := range c2 {
		fmt.Printf("Received: %d\n", num)
	}

	for num := range c3 {
		fmt.Printf("Received: %d\n", num)
	}
}

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)
	c3 := make(chan int, 5)

	go sender1(c1)
	go sender2(c2)
	go sender3(c3)

	go receiver1(c1)
	go receiver2(c2, c3)

	time.Sleep(4 * time.Second)
}
