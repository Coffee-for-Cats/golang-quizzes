package db

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type question_resp struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Option1  string `json:"option_1"`
	Option2  string `json:"option_2"`
}

type quiz_resp struct {
	Title     string          `json:"title"`
	Questions []question_resp `json:"questions"`
}

// This could be done as a single query
// I made 2 so I work a bit of my knowledge
// in golang's data structures.
func GetQuiz(quiz_id int) (quiz_resp, error) {
	// get all questions
	questions_rows, err := Use().Query(`SELECT id, question, correct_option, wrong_option
		FROM questions
		WHERE quiz_id = $1 
	`, quiz_id)
	defer questions_rows.Close()

	if err != nil {
		return quiz_resp{}, errors.New("Quiz not found.")
	}

	// get the quiz title
	row := Use().QueryRow(`SELECT title
		FROM quizes
		WHERE id = $1
	`, quiz_id)

	var title string
	err = row.Scan(&title)
	if err != nil {
		return quiz_resp{}, errors.New("Quiz not found.")
	}

	// needs initialization to avoid null in the json response
	questions := []question_resp{}

	for questions_rows.Next() {
		var id int
		var question, correct_option, wrong_option string
		questions_rows.Scan(&id, &question, &correct_option, &wrong_option)

		questions = append(questions, question_resp{
			Id:       id,
			Question: question,
			Option1:  correct_option,
			Option2:  wrong_option,
		})
	}

	return quiz_resp{
		Title:     title,
		Questions: questions,
	}, nil
}

// Crude and overengineered example of single query's approach.
func GetRandomQuiz() (quiz_resp, error) {
	row := Use().QueryRow(`SELECT title,
		-- title and a json object:
		(
			SELECT json_agg(
				json_build_object(
					'id', 			questions.id,
					'question', questions.question,
					'option_1', questions.correct_option,
					'option_2', questions.wrong_option
				)
			)
			FROM questions
			WHERE questions.quiz_id = quizes.id
		) AS questions -- json as question
		FROM quizes
		ORDER BY random()
		LIMIT 1;
	`)

	var title string
	var null_question sql.NullString

	err := row.Scan(&title, &null_question)
	if err != nil {
		panic(err)
	}

	question_json := "[]"
	if null_question.Valid {
		question_json = null_question.String
	}

	var questions_slc []question_resp
	err = json.Unmarshal([]byte(question_json), &questions_slc)
	if err != nil {
		return quiz_resp{}, errors.New("Error building json for response.")
	}

	return quiz_resp{
		Title:     title,
		Questions: questions_slc,
	}, nil
}

func GetQuizOwner(quiz_id int) (creator_name string, err error) {
	row := Use().QueryRow(`SELECT name
		FROM users
		WHERE id = (
			SELECT owner_id FROM quizes
			WHERE id = $1
		)
	`, quiz_id)

	var name string
	err = row.Scan(&name)

	return name, err
}

func CreateQuiz(creator string, title string) (quiz_id int, err error) {
	rows, err := Use().Query(`
		INSERT INTO quizes
		(title, owner_id)
		VALUES (
			$1,
			(SELECT id FROM users WHERE name = $2)
		) 
		RETURNING id
	`, title, creator,
	)
	defer rows.Close()

	if err != nil {
		return 0, err
	}

	var id int
	if !rows.Next() {
		return 0, errors.New("Error while returning quiz's id")
	}
	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}
