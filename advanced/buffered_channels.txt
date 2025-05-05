package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("Buffered channel", <-ch)

	ch <- 1
	ch <- 2

	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		fmt.Println("Receive:", <-ch)
	}()

	fmt.Println("Blocking")
	ch <- 3
	fmt.Println("Blocking ends")
	fmt.Println("Receive:", <-ch)
	fmt.Println("Receive:", <-ch)

	ch <- 4

	go func() {
		time.Sleep(2 * time.Second)
	}()
	fmt.Println("Value:", <-ch)
	fmt.Println("End")
}
