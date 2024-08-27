package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// create channel
// ch := make(chan int)
// ch <- 42 // Send the value 42 to the channel
// value := <-ch // Receive a value from the channel and store it in the 'value' variable
func worker(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate some work and send the worker's id to the channel
	fmt.Printf("Worker %d is sending data\n", id)
	ch <- id
}

func main() {
	//const numWorkers = 5
	//ch := make(chan int)
	//var wg sync.WaitGroup
	//for i := 0; i < numWorkers; i++ {
	//	wg.Add(1)
	//	go worker(i, ch, &wg)
	//}
	//ch := make(chan int)

	// Start a goroutine to send values to the channel
	//go func() {
	//	ch <- 1
	//	ch <- 2
	//	ch <- 3
	//	ch <- 4
	//	close(ch)
	//	//forget close channel then deadlock
	//}()

	// Receive values from the channel
	//for value := range ch {
	//	fmt.Println(value)
	//}

	//myChan := make(chan int) //unbuffered channel
	//myChan <- 1              // deadlock here
	// the sending goroutine will block, waiting for a receiving goroutine to be ready to receive the value.
	// no goroutine receive data
	//the main goroutine blocked on send
	//myChan := make(chan int)

	// Start a goroutine to receive the value
	//go func() {
	//	value := <-myChan
	//	println(value) // This will print "1"
	//}()

	// Send the value to the channel
	//myChan <- 1
	myChan := make(chan string, 10)
	// go routine dù chạy bất đồng bộ nhưng khi vào channel thì cũng phải đợi channel
	go func() {
		for i := 0; i < 10; i++ {
			myChan <- "go routine 1 " + strconv.Itoa(i)
			time.Sleep(2 * time.Second)
		}
		close(myChan)
	}()

	// using  for-range
	// Select handle multiple channel và WaitGroup
	for {
		val, ok := <-myChan
		if ok == false {
			fmt.Println("Channel closed ", val)
			break
		}
		fmt.Println("data ne = ", val)
	}

}
