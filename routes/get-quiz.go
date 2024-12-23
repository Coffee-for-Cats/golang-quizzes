package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GET /quiz/:id/
func GetQuiz(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Path
	pathParts := strings.Split(query, "/")
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "Invalid id path parameter.", 500)
		return
	}

	quiz, err := db.GetQuiz(id)
	if err != nil {
		http.Error(w, "Quiz not found.", 404)
		return
	}

	resp_text, err := json.MarshalIndent(quiz, "", "	")
	if err != nil {
		http.Error(w, "Error in database: invalid values.", 503)
		return
	}

	fmt.Fprintln(w, string(resp_text))
}

func GetRandomQuiz(w http.ResponseWriter, req *http.Request) {
	quiz, err := db.GetRandomQuiz()
	if err != nil {
		panic(err)
	}

	json_b, err := json.Marshal(quiz)
	if err != nil {
		http.Error(w, "Something went wrong. Give feedback to the dev.", 500)
	}

	fmt.Fprintln(w, string(json_b))
}
