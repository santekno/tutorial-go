package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// The function to connect database (Postgres)
func OpenDB(dsn string, setLimits bool) (*sql.DB, error) {
	// Connect the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Set limiter how many connection that can connect to database
	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	// Define 5 seconds timeout when connect to the database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the database to make sure the database connected
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// The function to close database connection
func CloseDB(db *sql.DB) {
	// Close the database connection
	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
