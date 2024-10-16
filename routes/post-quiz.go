package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type post_quiz_request struct {
	Title string `json:"title"`
}

func PostQuiz(w http.ResponseWriter, req *http.Request) {
	user, pass, ok := req.BasicAuth()
	if !ok {
		http.Error(w, "Invalid Basic Auth Format.", 401)
	}

	if !db.VerifyUser(user, pass) {
		http.Error(w, "Incorrect username or password.", 401)
		return
	}

	var title_struct post_quiz_request
	err := json.NewDecoder(req.Body).Decode(&title_struct)
	if err != nil {
		panic(err)
	}

	quiz_id, err := db.CreateQuiz(user, title_struct.Title)
	if err != nil {
		http.Error(w, "Error creating the quiz", 500)
	}

	fmt.Fprintf(w, "quiz_id: %x", quiz_id)
}
