package main

import (
	db "coffee-server/database"
	"fmt"
)

func main() {
	rows, err := db.Use().Query("SELECT option1, option2 FROM questions LIMIT 10;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var option1, option2 string
		err := rows.Scan(&option1, &option2)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s \n<or>\n%s\n\n", option1, option2)
	}
	fmt.Println("Exiting after success.")
}
