package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create ctx cancellable ctx
	ctx, cancel := context.WithCancel(context.Background())

	// launch goroutine that will listen to ctx cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 canceled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 1 working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 2 canceled", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 2 working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	fmt.Println("main function is working")
	time.Sleep(5 * time.Second)

	// cancel the ctx, which will signal all goroutine stop
	fmt.Println("Cancelling ctx")
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("main function is done")
}
