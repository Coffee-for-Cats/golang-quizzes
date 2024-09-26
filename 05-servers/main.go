package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHome)

	fmt.Println("🎉 Sever up! (on :8080) 🎉")
	http.ListenAndServe(":8080", nil)
}
