package main

import (
	"fmt"
	"time"
)

func main() {
	//numGoroutines := 3
	//done := make(chan int, 3)
	//
	//for i := range numGoroutines {
	//	go func(id int) {
	//		fmt.Printf("Goroutine %d working...\n", id)
	//		time.Sleep(time.Second)
	//		done <- id // sending signal of complete
	//	}(i)
	//}
	//
	//for range numGoroutines {
	//	<-done
	//}
	//
	//fmt.Println("All goroutines completed")

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}
}
