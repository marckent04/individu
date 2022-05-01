package outputs

import "fmt"

func DisplayQuestion(index int) {
	questions := []string{
		"Quel est votre nom ?",
		"Quel est votre date de naissance ?",
		"Quel est votre numéro de téléphone ?",
	}

	fmt.Println(questions[index])
}
func DisplayActionChoice() {
	questions := "Que voulez vous faire :\n 1- create \n 2- update \n 3- delete  \n 4- show \n"

	fmt.Println(questions)
}
