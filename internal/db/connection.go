package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteConnection struct {
	DB       *sql.DB
	FileName string
}

func NewSqliteConnection(filename string) *SqliteConnection {
	return &SqliteConnection{
		FileName: filename,
	}
}

func (m *SqliteConnection) Connect() error {
	var err error
	m.DB, err = sql.Open("sqlite3", m.FileName)
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
		return fmt.Errorf("connection failed: %w", err)
	}

	// Verify connection with ping
	if err = m.DB.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	return nil
}

func (m *SqliteConnection) Close() error {
	if m.DB != nil {
		return m.DB.Close()
	}
	return nil
}
