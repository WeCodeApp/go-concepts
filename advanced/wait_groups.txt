package main

import (
	"fmt"
	"sync"
	"time"
)

// example without channel
//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	fmt.Printf("Worker %d starting\n", id)
//	time.Sleep(time.Second)
//	fmt.Printf("Worker %d done\n", id)
//}
//
//func main() {
//	var wg sync.WaitGroup
//	numWorkers := 3
//
//	wg.Add(numWorkers)
//
//	for i := range numWorkers {
//		go worker(i, &wg)
//	}
//
//	wg.Wait()
//	fmt.Println("All workers done")
//}

func worker(id int, task <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	for task := range task {
		results <- task * 2
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("result:", result)
	}
}
