package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// simulate
		time.Sleep(time.Second)
		results <- task * 2
	}
}

func main() {
	//numWorkers := 3
	//numJobs := 10
	//
	//tasks := make(chan int, numJobs)
	//results := make(chan int, numJobs)
	//
	//// Create workers
	//for i := range numWorkers {
	//	go worker(i, tasks, results)
	//}
	//
	//for i := range numJobs {
	//	tasks <- i
	//}
	//close(tasks)
	//
	//for range numJobs {
	//	result := <-results
	//	fmt.Println("Result:", result)
	//}

	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket request
	for i := range numRequests {
		ticketRequests <- ticketRequest{
			personId:   i + 1,
			numTickets: (i + 1) * 2,
			cost:       price * (i + 1) * 2,
		}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("ticket for person id %d proccessed success\n", <-ticketResults)
	}
}

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}

func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for request := range requests {
		fmt.Printf("Processing %d ticke(s) of personID %d with total cost %d\n", request.numTickets, request.personId, request.cost)
		time.Sleep(time.Second)
		results <- request.personId
	}
}
