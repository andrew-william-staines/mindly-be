package config

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func ConnectDB () *sql.DB {
	connStr := "user=postgres password=Andrew dbname=mindy_db port=8080 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if(err != nil ) {
		log.Fatal("DB Connection Failed:",err)
	}
	
	return db
}
