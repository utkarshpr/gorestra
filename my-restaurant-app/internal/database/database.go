package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB creates a connection to the PostgreSQL database.
func ConnectDB() (*sql.DB, error) {
	connStr := "root:root1234@tcp(localhost:3306)/mydb"

	// Open the connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open a DB connection: %v", err)
	}

	// Ping the database to verify that the connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the DB: %v", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")
	return db, nil
}
