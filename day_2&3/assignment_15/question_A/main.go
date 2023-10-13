package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch chan int) {
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
}

func main() {
	ch := make(chan int)

	go sender(ch)
	go receiver(ch)

	time.Sleep(2 * time.Second)
}
