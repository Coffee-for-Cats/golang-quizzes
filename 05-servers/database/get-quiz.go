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
