package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetAnswer() (value string, err error) {
	reader := bufio.NewReader(os.Stdin)

	val, _, er := reader.ReadLine()

	if err != nil {
		err = er
		return
	}

	value = string(val)

	if len(value) == 0 {
		err = errors.New("Veuillez enter une valeur")
	}

	return
}
func GetChoiceAnswer() (value int, err error) {

	// fmt.Scanf("%d", &value)
	reader := bufio.NewReader(os.Stdin)

	val, _, er := reader.ReadLine()

	if err != nil {
		err = er
		return
	}
	value, _ = strconv.Atoi(string(val))

	if !(value >= 1 && value <= 4) {
		err = errors.New("Veuillez enter une valeur comprise entre 0 et 4")
	}

	return
}
func GetChoiceUpdateAnswer(min int, max int) (value int, err error) {

	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadLine()

	// fmt.Scanf("%d", &value)
	reader := bufio.NewReader(os.Stdin)

	val, _, er := reader.ReadLine()

	if err != nil {
		err = er
		return
	}
	value, _ = strconv.Atoi(string(val))

	if !(value >= min && value <= max) {
		// err = errors.New(fmt.Sprintf("Veuillez enter une valeur comprise entre %v et %v", min, max))
		err = errors.New(fmt.Sprintf("Veuillez enter une valeur comprise entre %v et %v", min, max))
	}

	return
}
func GetResponse() (value string, err error) {

	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadLine()

	// fmt.Scanf("%d", &value)

	reader := bufio.NewReader(os.Stdin)

	val, _, er := reader.ReadLine()

	if err != nil {
		err = er
		// return
	}
	value = string(val)

	if len(value) == 0 {
		err = errors.New("Veuillez enter une valeur")
	}

	return
}
func ConfirmNewUpdate() (value int, err error) {

	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadLine()

	// fmt.Scanf("%d", &value)

	reader := bufio.NewReader(os.Stdin)

	val, _, er := reader.ReadLine()

	if err != nil {
		err = er
		// return
	}
	value, _ = strconv.Atoi(string(val))

	if !(value == 1 || value == 2) {
		err = errors.New("Veuillez enter une valeur")
	}

	return
}
