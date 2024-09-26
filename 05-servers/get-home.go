package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func getHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello web!")

	db, err := sql.Open("postgres", "postgresql://fluffycat:s3cret@172.19.0.2:5432/heavycake?sslmode=disable")
	if err != nil {
		panic("Erro ao se conectar com o banco!")
	}
	defer db.Close()

	rows, err := db.Query("SELECT VERSION();")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Fprintln(w, rows)
	} else {
		fmt.Println("Nenhuma linha foi retornada/versão não encontrada.")
	}
}
