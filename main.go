package main

import (
	"coffee-server/routes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /quiz/random", routes.GetRandomQuiz)
	http.HandleFunc("GET  /quiz/{quizID}/", routes.GetQuiz)

	http.HandleFunc("POST /register", routes.Register)
	http.HandleFunc("POST /quiz", routes.PostQuiz)
	http.HandleFunc("POST /question", routes.PostQuestion)
	http.HandleFunc("POST /answer/{questionID}/", routes.PostResponse)

	fmt.Println("🎉 Sever up! (on :8080) 🎉")
	http.ListenAndServe(":8080", nil)
}
