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
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    )
    `

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Could not create users table: %v\n", err)
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMPTZ NOT NULL,
        user_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
    )
    `

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v\n", err)
	}

	fmt.Println("Tables created or already exist.")
}
