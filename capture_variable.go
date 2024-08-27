package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i) // Capturing loop variable i
		}(i)
	}

	time.Sleep(1 * time.Second) // Wait for all goroutines to finish
}
