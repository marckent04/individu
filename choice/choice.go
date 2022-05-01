package choice

import (
	"fmt"
	"individu/inputs"
	"individu/outputs"
	"individu/update"
)

func Mychoice() (int, error) {
	fmt.Printf("========================\n        INDIVIDU\n ========================\n ")
	outputs.DisplayActionChoice()
	return inputs.GetChoiceAnswer()
}

func Selectable(value int) {
	if value == 1 {
		outputs.DisplayCreateScreen()
	}
	if value == 2 {
		update.DisplayUpdateScreen()
	}
}
