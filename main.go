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
	}

	count, error := user.QueryCountUser(connectDb)
	const (
		size = 100
	)
	var (
		wg sync.WaitGroup
		ch = make(chan []user.User, size)
	)

	page := int(math.Floor(float64(count) / float64(size)))

	// work pool
	for i := 1; i <= page; i++ {
		wg.Add(1)
		offset := (i - 1) * size
		go func(page int, offset int) {
			defer func() {
				wg.Done()
			}()
			data, error := user.QueryUser(connectDb, size, offset)
			if error != nil {
				return
			}
			ch <- data
		}(i, offset)
	}

	if error != nil {
		fmt.Printf("query error")
	}

	// Close channel after all goroutines are done
	go func() {
		fmt.Printf("close channel\n")
		wg.Wait()
		close(ch) // nên close channel khi tất cả data đã dc send
	}()

	file, err := os.Create("output-1.sql")
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf(err.Error())
		}
	}(file)

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
	duration := time.Since(start) // Calculate the duration
	fmt.Printf("Execution time: %v\n", duration)
	//73.5985ms

}
