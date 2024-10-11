package main

import (
	db "coffee-server/database"
	"fmt"
)

func main() {
	_, err := db.Use().Exec(`
		DROP TABLE IF EXISTS questions CASCADE;
		DROP TABLE IF EXISTS quizes CASCADE;
		DROP TABLE IF EXISTS users CASCADE;
	`)
	if err != nil {
		panic(err)
	}

	// users table
	_, err = db.Use().Exec(`CREATE TABLE users(
		id SERIAL PRIMARY KEY,
		-- necessary fields:
		name TEXT NOT NULL UNIQUE,
		salt TEXT NOT NULL,
		hash TEXT NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	// quizes table
	_, err = db.Use().Exec(`CREATE TABLE quizes(
		id SERIAL PRIMARY KEY,
		-- necessary fields:
		title TEXT NOT NULL,
		owner_id INT NOT NULL REFERENCES users(id)
	)`)

	// questions table
	_, err = db.Use().Exec(`CREATE TABLE questions(
		id SERIAL PRIMARY KEY,
		-- necessary fields:
		question TEXT NOT NULL UNIQUE,
		correct_option TEXT NOT NULL,
		wrong_option TEXT NOT NULL,
		quiz_id INT NOT NULL REFERENCES quizes(id),
		--
		guess_count INT DEFAULT 0,  -- how many times it got guessed
		hit_count INT DEFAULT 0		  -- how many times it was right
	)`)
	if err != nil {
		panic(err)
	}

	// # random content

	err = db.RegisterUser("User 1", "cafe-babe")
	if err != nil {
		panic(err)
	}

	_, err = db.CreateQuiz(1, "Opinions matter")
	if err != nil {
		panic(err)
	}

	_, err = db.CreateQuestion(
		"What's better?", // question
		"Coffee", "Tea",  // correct | wrong
		1, // creator id
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("🎉 Success 🎉")
}
