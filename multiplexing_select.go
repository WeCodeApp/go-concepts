package main

import (
	"fmt"
	"time"
)

func main() {
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	//time.Sleep(2 * time.Second)
	//	ch1 <- 1
	//}()
	//
	//go func() {
	//	//time.Sleep(1 * time.Second)
	//	ch2 <- 1
	//}()
	//
	//time.Sleep(2 * time.Second)
	//
	//for range 2 {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println("Received from ch1:", msg1)
	//	case msg2 := <-ch2:
	//		fmt.Println("Received from ch2:", msg2)
	//	default:
	//		fmt.Println("No channels ready")
	//	}
	//}

	ch := make(chan int)

	go func() {
		//time.Sleep(3 * time.Second)
		ch <- 1
		close(ch)
	}()

	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Println("Channel closed")
			// clean up activities
			return
		}
		fmt.Println("Received from ch:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}

}
