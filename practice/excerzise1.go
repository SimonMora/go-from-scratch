package practice

import (
	"strconv"
)

func EvaluateInteger(num string) (int, string) {
	var text string
	integer, error := strconv.Atoi(num)

	if error != nil {
		return 0, string("There was an error: " + error.Error())
	}

	if integer > 100 {
		text = "The number is bigger than 100."
	} else {
		text = "The number is lower than 100."
	}

	return integer, text
}
