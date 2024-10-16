package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type post_question_request struct {
	Question string `json:"question"`
	Correct  string `json:"correct"`
	Wrong    string `json:"wrong"`
	Id       int    `json:"quiz_id"`
}

type id_response struct {
	Id int `json:"id"`
}

func PostQuestion(w http.ResponseWriter, req *http.Request) {
	var request post_question_request
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", 401)
		return
	}

	// verify user
	user, pass, ok := req.BasicAuth()
	if !ok || !db.VerifyUser(user, pass) {
		http.Error(w, "Unauthorized", 403)
		return
	}

	// verify quiz owner
	quiz_owner, err := db.GetQuizOwner(request.Id)
	if err != nil {
		http.Error(w, "Unauthorized or wrong quiz_id", 403)
		return
	}

	if user != quiz_owner {
		http.Error(w, "Unauthorized.", 403)
		return
	}

	// create record
	id, err := db.CreateQuestion(request.Question, request.Correct, request.Wrong, request.Id)
	if err != nil {
		http.Error(w, "Question already exists.", 409)
		return
	}

	response_str, err := json.MarshalIndent(id_response{
		Id: id,
	}, "", "	")

	if err != nil {
		http.Error(w, "Error returning id.", 500)
		fmt.Fprintln(w, "Id should be: ", id)
	}

	fmt.Fprint(w, string(response_str))
}
