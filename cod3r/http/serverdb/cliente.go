package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SuccessMsg struct {
	Status_Code int
	Message     string
}

// Usuario :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	// 111 é workaround para uma rota especifica
	case r.Method == "GET" && id == 111:
		ultimoCadastrado(w, r)
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "POST":
		var u Usuario
		json.NewDecoder(r.Body).Decode(&u)
		criarUsuario(w, r, u)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, _ := initDB()
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, _ := initDB()
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func criarUsuario(w http.ResponseWriter, r *http.Request, u Usuario) {
	db, _ := initDB()
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec(u.Nome)
	stmt.Close()

	jsonMsg, _ := json.Marshal(SuccessMsg{Status_Code: 201, Message: "User created successfully!"})

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonMsg))
}

func ultimoCadastrado(w http.ResponseWriter, r *http.Request) {
	db, _ := initDB()
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios order by id desc").Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
