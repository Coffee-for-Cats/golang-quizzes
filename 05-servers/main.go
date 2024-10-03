package main

import (
	"coffee-server/routes"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("POST /response", routes.PostResponse)
	http.HandleFunc("/", routes.GetQuetion)

	fmt.Println("🎉 Sever up! (on :8080) 🎉")
	http.ListenAndServe(":8080", nil)
}
