package outputs

import (
	"fmt"
	"individu/database"
	"individu/inputs"
)

func QuestionResponse(index int) (string, error) {
	DisplayQuestion(index)
	return inputs.GetAnswer()
}

func DisplayCreateScreen() {

	var values []string

	loop := 0

	for loop <= 2 {
		value, err := QuestionResponse(loop)

		if err != nil {
			fmt.Println(err)
			continue
		}

		values = append(values, value)
		loop++
	}

	db, err := database.DbConnect("individu.sqlite")

	if err != nil {
		fmt.Println(err)
	}

	database.LaunchMigrations(db)

	newUser := database.User{
		Name:     values[0],
		Birthday: values[1],
		Tel:      values[2],
	}

	tx := newUser.Create(db)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	if tx.RowsAffected == 1 {
		fmt.Println("Utilisateur ajouté avec succès")
	}
}
