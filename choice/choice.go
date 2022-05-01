package choice

import (
	"individu/inputs"
	"individu/outputs"
	"individu/update"
)

func Mychoice() (int, error) {
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
