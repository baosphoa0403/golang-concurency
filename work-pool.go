package main

import (
	"fmt"
	"interview/db"
	"interview/user"
	"math"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now() // Capture the start time
	connectDb, err := db.ConnectDb()
	if err != nil {
		fmt.Printf("error connect database")
		return
	}

	count, err := user.QueryCountUser(connectDb)
	if err != nil {
		fmt.Printf("error querying user count")
		return
	}

	const (
		size       = 100
		numWorkers = 5 // Fixed number of workers in the pool
	)
	var (
		wg sync.WaitGroup
		ch = make(chan []user.User, size)
	)

	page := int(math.Ceil(float64(count) / float64(size)))
	taskCh := make(chan int, page)

	// Start the worker pool
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for offset := range taskCh {
				data, err := user.QueryUser(connectDb, size, offset)
				fmt.Printf("Worker %d: querying users", workerId)
				if err != nil {
					fmt.Printf("Worker %d: error querying users: %v\n", workerId, err)
					continue
				}
				ch <- data
			}
		}(w)
	}

	// Send tasks to the workers
	go func() {
		for i := 1; i <= page; i++ {
			offset := (i - 1) * size
			taskCh <- offset
		}
		close(taskCh) // Close the task channel after sending all tasks
	}()

	// Close the data channel after all workers are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Process results
	// Open a file to write the SQL insert statements
	file, err := os.Create("output.sql")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer file.Close()

	// Process results and generate SQL insert statements
	for val := range ch {
		for _, user := range val {
			// Generate SQL insert statement
			sql := fmt.Sprintf(
				"INSERT INTO users (id, name, age, address, email) VALUES (%s, '%s', %s, '%s', '%s');\n",
				user.ID, user.NAME, user.AGE, user.ADDRESS, user.EMAIL,
			)
			// Write the SQL statement to the file
			_, err := file.WriteString(sql)
			if err != nil {
				fmt.Printf("Failed to write to file: %v\n", err)
				return
			}
		}
	}

	fmt.Println("SQL script has been generated in output.sql")

	fmt.Println("All tasks completed")
	duration := time.Since(start) // Calculate the duration
	fmt.Printf("Execution time: %v\n", duration)
	//Execution time: 29.039083ms

}
