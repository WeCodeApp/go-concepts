package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)

	time.Sleep(5 * time.Second)
}

// Send only
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}
