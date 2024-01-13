package main

import (
	"database/sql"
	"log"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "bruno:test123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
