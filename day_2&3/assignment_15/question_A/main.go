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
	v := <-ch
	fmt.Println("Recieved", v)
	p := <-ch
	fmt.Println("Recieved", p)
	q := <-ch
	fmt.Println("Recieved", q)
	r := <-ch
	fmt.Println("Recieved", r)
	s := <-ch
	fmt.Println("Recieved", s)
}

func main() {
	ch := make(chan int, 5)

	go sender(ch)
	go receiver(ch)

	time.Sleep(10 * time.Second)
}
