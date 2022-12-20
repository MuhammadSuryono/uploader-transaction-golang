package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const dbConnString = "postgres://transactions_service:5p3S9POusrYXlH1oOw4@34.101.89.115:8435/edc_prod?sslmode=disable"
const dbMaxIdleConns = 4
const dbMaxConns = 100

func DBOpenConnection() (*sql.DB, error) {
	log.Println("=> open db connection")

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	log.Println("=> Success connect to DB", dbConnString)

	return db, nil
}
