package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type post_response_req struct {
	Question int    `json:"question_id"`
	Answer   string `json:"answer_text"`
}

func PostResponse(w http.ResponseWriter, req *http.Request) {
	var response post_response_req
	err := json.NewDecoder(req.Body).Decode(&response)
	if err != nil {
		http.Error(w, "Bad request.", 400)
		return
	}

	correct, err := db.Answer(response.Question, response.Answer)
	if err != nil {
		http.Error(w, "Bad request.", 400)
		return
	}

	// saves scores for the player
	user, pass, ok := req.BasicAuth()
	if ok {
		// someone is logged
		if !db.VerifyUser(user, pass) {
			// someone is trying to mess other people score.
			http.Error(w, "Unauthorized", 405)
			return
		}

		// logged and authenticated!
		err = db.Answered(user, correct)
		if err != nil {
			// this should never happen ( >->)
			http.Error(w, "Unauthorized", 405)
			return
		}

	} // no else: maybe the guy is just not logged.

	if correct {
		w.WriteHeader(201)
		fmt.Fprintln(w, "✅ You're a Jedi!")
	} else {
		w.WriteHeader(202)
		fmt.Fprintln(w, "❌ Missed the shot.")
	}
}
