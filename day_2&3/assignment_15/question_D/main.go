package main

import (
	"fmt"
	"sync"
)

func Sender(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	if i == 2 {
		close(ch)
	}
}
func Receiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}

}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 5)

	for i := 0; i <= 2; i++ {
		wg.Add(1)
		// defer wg.Done()
		go Sender(ch, &wg, i)
	}
	for i := 0; i <= 1; i++ {
		wg.Add(1)
		// defer wg.Done()
		go Receiver(ch, &wg)
	}
	// time.Sleep(5 * time.Second)
	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func sender1(c1 chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 5; i++ {
// 		c1 <- i
// 	}
// 	close(c1)
// }

// func sender2(c2 chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 6; i <= 10; i++ {
// 		c2 <- i
// 	}
// 	close(c2)
// }

// func sender3(c3 chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 11; i <= 15; i++ {
// 		c3 <- i
// 	}
// 	close(c3)
// }

// func receiver1(c1 chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range c1 {
// 		fmt.Printf("Received: %d\n", num)
// 	}
// }

// func receiver2(c2 chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num := range c2 {
// 		fmt.Printf("Received: %d\n", num)
// 	}
// }

// func main() {

// 	wg := &sync.WaitGroup{}

// 	wg.Add(5)

// 	c1 := make(chan int, 5)
// 	c2 := make(chan int, 5)
// 	c3 := make(chan int, 5)

// 	go sender1(c1, wg)
// 	go sender2(c2, wg)
// 	go sender3(c3, wg)

// 	go receiver1(c1, wg)
// 	go receiver2(c2, wg)

// 	fmt.Println("End of the main")

// 	wg.Wait()
// }
