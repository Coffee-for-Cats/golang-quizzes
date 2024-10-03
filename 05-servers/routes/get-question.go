package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type question struct {
	Id      int    `json:"id"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
}

func GetQuetion(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Use().Query(`SELECT  id, option1, option2 
	FROM questions 
	ORDER BY RANDOM()
	LIMIT 1;
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int
	var option1, option2 string

	for rows.Next() {
		err := rows.Scan(&id, &option1, &option2)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(id)
	fmt.Println(option1)
	fmt.Println(option2)

	response, err := json.MarshalIndent(question{
		Id:      id,
		Option1: option1,
		Option2: option2,
	}, "", "  ")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response))
}
