package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "bruno:test123@/")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("create database if not exists bruno_test")
	db.Exec("use bruno_test")
	db.Exec(`create table if not exists users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY (id)
	)`)

	return db, err
}
