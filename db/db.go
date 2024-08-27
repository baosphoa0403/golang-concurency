package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func ConnectDb() (*DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/sandbox")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		// Close the database connection if the ping fails
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) QueryStatement(statement string, args ...interface{}) (*sql.Rows, error) {
	rows, error := db.DB.Query(statement, args...)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	return rows, nil
}
