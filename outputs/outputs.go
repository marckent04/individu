package outputs

import "fmt"


func DisplayQuestion(index int)  {
	questions := []string{
		"Quel est votre nom ?",
		"Quel est votre date de naissance ?",
		"Quel est votre numéro de téléphone ?",
	}

	fmt.Println(questions[index])
}
