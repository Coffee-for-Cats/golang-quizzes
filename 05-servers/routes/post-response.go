package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Option int `json:"option"`
	Id     int `json:"id"`
}

func PostResponse(w http.ResponseWriter, req *http.Request) {

	var content response
	// var json_bytes string
	err := json.NewDecoder(req.Body).Decode(&content)
	if err != nil {
		http.Error(w, "Bad Request, invalid json", 400)
		return
	}

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

	fmt.Fprintln(w, "Ok!")
}
