package inputs

import (
	"bufio"
	"errors"
	"os"
)


func GetAnswer()  (value string, err error ) {
reader := 	bufio.NewReader(os.Stdin)

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