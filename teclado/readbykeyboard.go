package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var leyenda string
var err error

func ReadNumbersAndMultiply() {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type the first number: ")
	if scn.Scan() {
		number1, err = strconv.Atoi(scn.Text())
		if err != nil {
			panic("Please review the data as it failed to parse because: " + err.Error())
		}
	}

	fmt.Println("Please type the second number: ")
	if scn.Scan() {
		number2, err = strconv.Atoi(scn.Text())
		if err != nil {
			panic("Please review the data as it failed to parse because: " + err.Error())
		}
	}

	fmt.Println("Please type the leyend: ")
	if scn.Scan() {
		leyenda = scn.Text()
	}

	fmt.Println(leyenda, number1*number2)
}
