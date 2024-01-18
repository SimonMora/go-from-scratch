package funciones

import "fmt"

func Calculos() {
	var num3 = 30
	var num4 = 10

	sum := func(number1 int, number2 int) int {
		return number1 + number2 + num3 + num4
	}

	fmt.Println(sum(14, 36))

	sum = func(number1, number2 int) int {
		return number1*number2 + num3*num4
	}

	fmt.Println(sum(14, 36))
}
