package main

import (
	"coffee-server/routes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/register", routes.Register)

	http.HandleFunc("/quiz/{quizID}/", routes.GetQuiz)
	http.HandleFunc("/quiz/random/", routes.GetRandomQuiz)

	fmt.Println("🎉 Sever up! (on :8080) 🎉")
	http.ListenAndServe(":8080", nil)
}
