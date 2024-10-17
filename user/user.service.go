package user

import (
	"database/sql"
	"fmt"
	"interview/db"
)

type User struct {
	ID      string
	NAME    string
	AGE     int
	ADDRESS string
	EMAIL   string
}

func QueryUser(db *db.DB, limit int, offset int) ([]User, error) {
	fmt.Printf("QueryUser limit: %d - offset : %d\n", limit, offset)
	var users []User

	rows, err := db.QueryStatement("SELECT id, name, age, address, email FROM users LIMIT ? OFFSET ? ", limit, offset)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.NAME, &user.AGE, &user.ADDRESS, &user.EMAIL); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil { // Check for any errors encountered during iteration
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return users, nil
}

func QueryCountUser(db *db.DB) (int, error) {
	var count int
	if err := db.DB.QueryRow("SELECT count(*) FROM users").Scan(&count); err != nil {
		fmt.Println("error", err)
		return 0, err
	}

	return count, nil
}
