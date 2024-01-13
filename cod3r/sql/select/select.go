package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "bruno:test123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var usuarios []usuario
	rows, _ := db.Query("select * from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		usuarios = append(usuarios, u)
	}

	fmt.Println(usuarios)
}
