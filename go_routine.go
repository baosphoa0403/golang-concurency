package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	//go func() {
	//	fmt.Println("Hello, world.")
	//}()
	//time.Sleep(time.Second) // using time main function call time pause
	// Start a new goroutine
	//go printMessage("Hello from Goroutine 1")

	// Start another goroutine
	//go printMessage("Hello from Goroutine 2")

	// Wait for 2 seconds to let the goroutines finish
	//time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	// Add 3 goroutines to the wait group
	//wg.Add(4) the 4th call to wg.Done() will never happen, and the program will deadlock.
	//wg.Add(2) and then start 3 goroutines, it will not cause a deadlock.
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go doWork(&wg, "Task "+strconv.Itoa(i))
	}
	// Start the goroutines

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All tasks completed!")
}

func doWork(wg *sync.WaitGroup, task string) {
	// Defer the Done() call to signal that the task is completed
	defer wg.Done()

	fmt.Println("Starting", task)
	// Simulate some work
	time.Sleep(2 * time.Second)
	fmt.Println("Completed", task)
}

//func printMessage(message string) {
//	fmt.Println(message)
//}
