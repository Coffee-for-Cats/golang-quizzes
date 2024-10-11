package db

import "errors"

func CreateQuestion(
	question, correct_option, wrong_option string,
	quiz_id int,
) (int, error) {

	result, err := Use().Query(`
		INSERT INTO questions
		(question, correct_option, wrong_option, quiz_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, question, correct_option, wrong_option, quiz_id,
	)

	if err != nil {
		return 0, err
	}

	var id int
	if !result.Next() {
		return 0, errors.New("Error while returning question's id.")
	}
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CreateQuiz(creator_id int, title string) (int, error) {
	rows, err := Use().Query(`
		INSERT INTO quizes
		(title, owner_id)
		VALUES ($1, $2) 
		RETURNING owner_id
	`, title, creator_id,
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
