package main

import (
	"fmt"
	"sync"
)

type Result struct {
	data     int
	workerId int
}

func worker1(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	var (
		wg sync.WaitGroup
	)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(jobs <-chan int) {
			fmt.Printf("%v", wg)
			defer func() {
				wg.Done()
			}()
			for j := range jobs {
				fmt.Printf("data %d\n", j*2)
			}
		}(jobs)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()

	// for {
	// 	select {
	// 	case val, ok := <-jobs:
	// 		if !ok { // Channel closed
	// 			fmt.Println("No more data. Exiting loop.")
	// 			return
	// 		}
	// 		// Process the received data
	// 		fmt.Printf("Data: %d\n", val)
	// 	default:
	// 		// Optionally do something while waiting for new data
	// 		time.Sleep(100 * time.Millisecond) // Avoid busy-waiting
	// 	}
	// }
}
