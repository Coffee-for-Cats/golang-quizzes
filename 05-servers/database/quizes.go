package db

import "errors"

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

func GetQuiz(quiz_id int) (quiz_resp, error) {
	// get all questions
	questions_rows, err := Use().Query(`SELECT id, question, correct_option, wrong_option
		FROM questions
		WHERE quiz_id = $1 
	`, quiz_id)

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

	var questions []question_resp

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
