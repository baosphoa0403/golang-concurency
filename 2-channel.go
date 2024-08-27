package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//unbufferedExample()
	buffered()
}

func buffered() {
	ch := make(chan string, 4) // Buffered channel with a capacity of 2

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "send data " + strconv.Itoa(i)
			fmt.Println("send data", i)
		}
		close(ch) // Close the channel after sending all data
	}()

	// Receiver goroutine
	for val := range ch {
		fmt.Println("received = ", val)
		time.Sleep(time.Second)
	}

	fmt.Println("Buffered Channel Example Done")
}

func unbufferedExample() {
	ch := make(chan string, 0) // Buffered channel with a capacity of 2

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "send data " + strconv.Itoa(i)
			fmt.Println("send data", i)
		}
		close(ch) // Close the channel after sending all data
	}()

	// Receiver goroutine
	for val := range ch {
		fmt.Println("received = ", val)
		time.Sleep(time.Second)
	}

	fmt.Println("Buffered Channel Example Done")
}
