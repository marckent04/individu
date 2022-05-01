package main

import (
	"fmt"
	"individu/choice"
)

func main() {

	for {
		response, error := choice.Mychoice()

		if error != nil {
			fmt.Println(error)
			continue
		}
		choice.Selectable(response)

		break
	}

}
