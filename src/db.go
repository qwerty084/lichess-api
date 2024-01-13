package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// DB represents a database connection.
type DB struct {
	conn *pgx.Conn
}

// NewDB creates a new instance of DB and establishes a connection to the database.
// It retrieves the necessary database connection details from environment variables.
// Returns a pointer to the DB struct and an error if the connection fails.
func NewDB() (*DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

// CreateTable creates a table in the database using the SQL query specified in the file at the given file path.
// It reads the query from the file, executes it using the database connection, and returns an error if any.
func (db *DB) CreateTable(filePath string) error {
	query, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if _, err := db.conn.Query(context.Background(), string(query)); err != nil {
		fmt.Printf("There was an error creating a table: %s", err)
		return err
	}

	return nil
}
