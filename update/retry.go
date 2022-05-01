package update

import (
	"fmt"
	"individu/database"
	"individu/inputs"

	"gorm.io/gorm"
)

func Retry(selected database.User, db *gorm.DB) {
	for {
		fmt.Printf("voulez vous modifier un autre champ \n 1- Oui \n 2-Non \n")
		res, err := inputs.ConfirmNewUpdate()
		if err != nil {
			fmt.Println(err)
		}
		if res == 1 {
			UpdateColums(selected, db)
		}
		if res == 2 {
			fmt.Printf("Merci ")

			break
		}
	}
}

func UpdateColums(selected database.User, db *gorm.DB) {

	fieldName := ShowToSelectColumn(int(selected.ID))

	UpdateSelectedColumn(db, selected, fieldName)

	fmt.Printf("le champ %s bien été modifié \n", fieldName)

}
