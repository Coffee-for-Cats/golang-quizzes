package main

import (
	db "coffee-server/database"
	"fmt"
	"net/http"
)

func getHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello web!")

	rows, err := db.Use().Query("SELECT VERSION();")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var version string
		rows.Scan(&version)
		fmt.Fprintln(w, rows)
	} else {
		fmt.Println("versão não encontrada.")
	}
}
