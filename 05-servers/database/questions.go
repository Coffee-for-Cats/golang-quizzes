package db

import (
	"errors"
)

func CreateQuestion(
	question, correct_option, wrong_option string,
	quiz_id int,
) (int, error) {

	if correct_option == wrong_option {
		return 0, errors.New("Options should be different.")
	}

	result, err := Use().Query(`INSERT INTO questions
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

func Answer(question_id int, answer string) (bool, error) {
	row := Use().QueryRow(`SELECT correct_option, wrong_option
		FROM questions
		WHERE id = $1
	`, question_id)

	var correct, wrong string
	err := row.Scan(&correct, &wrong)
	if err != nil {
		return false, err
	}

	increment := 0
	if answer == correct {
		increment = 1
	} else if answer != wrong {
		// user sent random text instead of the question's text
		return false, errors.New("Invalid option/answer.")
	}

	_, err = Use().Exec(`UPDATE questions
		SET hit_count = hit_count + $1,
				guess_count = guess_count + 1
		WHERE id = $2
		`, increment, question_id,
	)

	if err != nil {
		return false, errors.New("Error updating the question's score.")
	}

	return increment == 1, nil
}
