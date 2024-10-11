package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type answer_request struct {
	Option int `json:"option"`
	Id     int `json:"id"`
}

type answer_response struct {
	Option1_count int `json:"option1_count"`
	Option2_count int `json:"option2_count"`
}

func PostResponse(w http.ResponseWriter, req *http.Request) {

	var content answer_request
	err := json.NewDecoder(req.Body).Decode(&content)
	if err != nil {
		http.Error(w, "Bad Request, invalid json", 400)
		return
	}

	// no SQL Inject since Option is always an int.
	column := fmt.Sprintf("option%d_count", content.Option)
	query := fmt.Sprintf(`
		UPDATE questions
		SET %s = %s + 1
		WHERE id = $1
	`, column, column)

	_, err2 := db.Use().Exec(query, content.Id)
	if err2 != nil {
		http.Error(w, "Invalid ID or Option", 500)
		return
	}

	// send the counts so the frontend calculates the %'s.
	row := db.Use().QueryRow(`
		SELECT option1_count, option2_count
		FROM questions
		WHERE id = $1
	`, content.Id)

	var option1_count, option2_count int
	err3 := row.Scan(&option1_count, &option2_count)
	if err3 != nil {
		panic(err3)
	}

	response_str, err4 := json.MarshalIndent(answer_response{
		Option1_count: option1_count,
		Option2_count: option2_count,
	}, "", "	")

	if err4 != nil {
		http.Error(w, "Conection Failed", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(response_str))
}
