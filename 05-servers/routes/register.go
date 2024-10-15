package routes

import (
	db "coffee-server/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type register_req struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

func Register(w http.ResponseWriter, req *http.Request) {
	var body register_req
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request.", 400)
		return
	}

	err = db.RegisterUser(body.User, body.Pass)
	if err != nil {
		http.Error(w, "Invalid or already taken username.", 500)
		return
	}

	// TODO: save Authorization in a header.

	fmt.Fprintln(w, "Ok")
}
