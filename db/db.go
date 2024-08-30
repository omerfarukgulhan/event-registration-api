package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connString := "postgres://postgres:153515@localhost:5432/workshops"

	DB, err = sql.Open("pgx", connString)
	if err != nil {
		log.Fatalf("Could not connect to the PostgreSQL database: %v\n", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		log.Fatalf("Could not ping the PostgreSQL database: %v\n", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	createTables()
}

func createTables() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMPTZ NOT NULL,
        user_id INTEGER
    )
    `

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v\n", err)
	}

	fmt.Println("Events table created or already exists.")
}
