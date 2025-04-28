package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func DB_Connection() *DB {
	dbConn, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db: dbConn}
}

func (db *DB) CreateTable () error {

	sqlQuery :=


	if err := 
}