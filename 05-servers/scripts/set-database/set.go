package main

import (
	db "coffee-server/database"
	"fmt"
)

func main() {
	_, err := db.Use().Query(`DROP TABLE IF EXISTS questions;`)
	if err != nil {
		panic(err)
	}

	_, err2 := db.Use().Query(`CREATE TABLE questions(
		id SERIAL PRIMARY KEY,
		option1 TEXT NOT NULL,
		option2 TEXT NOT NULL,
		option1_count INT DEFAULT 0,
		option2_count INT DEFAULT 0
	);`)
	if err2 != nil {
		panic(err2)
	}

	_, err3 := db.Use().Query(`INSERT INTO questions (option1, option2)
	VALUES (
		'Unlimited coffee suply (you cant sell it).',
		'Unlimited tea suply (you cant sell it).'
	);
	`)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println("Exiting after success.")
}
