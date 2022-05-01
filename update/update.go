package update

import (
	"fmt"
	"individu/database"
)

func DisplayUpdateScreen() {

	db, _ := database.DbConnect("individu.sqlite")

	var allUsers []database.User

	var user database.User

	db.Find(&allUsers)

	selected := ShowToSelect(allUsers)

	db.First(&user)

	fmt.Printf("je veux modifier le numero %d \n", selected.ID)

	UpdateColums(selected, db)

	Retry(selected, db)

}
