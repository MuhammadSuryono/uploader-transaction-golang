package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

const dbConnString = "postgres://postgres:password@localhost:5433/majestic?sslmode=disable"
const dbConnStringMysql = "dbuser:5FSX65VtBj#sb__G@tcp(ec2-18-142-248-178.ap-southeast-1.compute.amazonaws.com:3306)/db_messenger?charset=utf8mb4&parseTime=True&loc=Local"
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

func DBOpenConnectionMysql() (*sql.DB, error) {
	log.Println("=> open db connection msql")

	db, err := sql.Open("mysql", dbConnStringMysql)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)

	log.Println("=> Success connect to DB MYSQL", dbConnStringMysql)

	return db, nil
}
