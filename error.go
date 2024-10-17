package main

import "fmt"

func main() {
	// dataChan := make(chan string) //adds data to the channel

	// go func() {
	// 	dataChan <- "Hey Champ!" // gets data from the channel
	// }()
	// c := <-dataChan

	// fmt.Println("Main function finished.")
	// fmt.Println(c)
	// time.Sleep(time.Second * 10)
	// deadlock
	// because unbuffered block until send and receiver nhận dc message

	// solution 1
	// dataChan := make(chan string, 1) //buffered channel
	// solution 2
	// use go routine run background
	// -> do unbuffered blocking program

	// chứng minh unbuffered block

	// dataIntChan := make(chan int)

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		dataIntChan <- i
	// 		fmt.Print("send data index = ", i, "\n")
	// 	}
	// 	close(dataIntChan) // forget deadlock
	// 	fmt.Print("send all success")
	// }()

	// for i := 0; i < 10; i++ {
	// 	data := <-dataIntChan
	// 	fmt.Print("nhận data index = ", data, "\n")
	// }

	// Using range to receive from channel until it's closed
	// for data := range dataIntChan {
	// 	fmt.Print("received data index = ", data, "\n")
	// }

	// for {
	// 	select {
	// 	case val, done := <-dataIntChan:
	// 		if !done {
	// 			fmt.Println("Channel closed, exiting.")
	// 			// goto END_LOOP
	// 			return
	// 		}
	// 		fmt.Print("data ne index = ", val, "\n")
	// 	}
	// }
	// END_LOOP:
	// 	fmt.Printf("run complete")
	// 	time.Sleep(1 * time.Second)
	// ok == true: The channel is open, and the value was successfully received.
	// ok == false: The channel has been closed, and no more values can be received.

	// run background
	// go func() {
	// 	fmt.Print("hello")
	// }()

	//  must use time for main thread wait
	// time.Sleep(time.Second * 5)

	// solution use
	// var wg sync.WaitGroup

	// wg.Add(1)

	// // Launch a goroutine to run in the background
	// go func() {
	// 	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	// 	fmt.Println("hello")
	// }()

	// // Wait for the goroutine to finish
	// wg.Wait()
	// fmt.Println("Goroutine has finished, main function exiting.")

	dataChan := make(chan int, 1) // Unbuffered channel

	// Goroutine 1: Tries to send data
	go func() {
		fmt.Println("Sending data...")
		dataChan <- 1 // This will block because no one is receiving yet
		fmt.Println("Data sent.")
	}()

	// Goroutine 2: Tries to receive data
	go func() {
		fmt.Println("Receiving data...")
		data := <-dataChan // This will block indefinitely because Goroutine 1 is stuck
		fmt.Println("Data received:", data)
	}()

	// The main goroutine exits immediately, causing the goroutines to deadlock
	// because Goroutine 1 is waiting to send data, and Goroutine 2 is waiting to receive it.
	fmt.Println("Main function exiting.")

}
