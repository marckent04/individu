package update

import (
	"fmt"
	"individu/database"
	"individu/inputs"

	"gorm.io/gorm"
)

func ShowToSelect(users []database.User) (choiceToUpdate database.User) {

	for {

		fmt.Printf("Quel compte voulez vous modifier \n")
		for _, result := range users {
			fmt.Printf("%d - %s \n", result.ID, result.Name)
		}
		val, err := inputs.GetChoiceUpdateAnswer(int(users[0].ID), int(users[len(users)-1].ID))
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {

			for _, row := range users {
				if int(row.ID) == val {

					choiceToUpdate = row
				}
			}

			break
		}

	}
	return choiceToUpdate
}
func ShowToSelectColumn(id int) (nameField string) {
	columns := []string{"Name", "Birthday", "Tel"}

	for {

		fmt.Printf("Quel champ voulez vous modifier \n")

		for index, column := range columns {
			fmt.Printf("%d - %s \n", index, column)
		}
		val, err := inputs.GetChoiceUpdateAnswer(0, len(columns)-1)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			nameField = columns[val]
			break
		}

	}
	return
}
func UpdateSelectedColumn(db *gorm.DB, user database.User, fieldName string) {
	for {

		fmt.Printf("la nouvelle valeur de %s  sera : \n", fieldName)
		res, err := inputs.GetResponse()
		if err != nil {
			fmt.Println(err)
			// break
		}

		errordb := db.Model(&user).Update(fieldName, res).Error
		if errordb != nil {
			fmt.Println(errordb)
		}
		if errordb == nil {
			break
		}

	}

}
