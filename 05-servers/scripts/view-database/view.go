package main

import (
	db "coffee-server/database"
	"fmt"
)

func main() {
	rows, err := db.Use().Query(`SELECT option1, option2, option1_count, option2_count
	FROM questions LIMIT 10;
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		var option1, option2 string
		var option1_count, option2_count int
		err := rows.Scan(&option1, &option2, &option1_count, &option2_count)
		if err != nil {
			panic(err)
		}

		fmt.Printf("{\n"+
			"  Option 1 (%d): %s \n"+
			"  Option 2 (%d): %s \n"+
			"}\n",
			option1_count, option1,
			option2_count, option2,
		)
	}
	fmt.Println("Exiting after success.")
}
