package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func OtherVariables() {
	Nombre = "Simon"
	Estado = true
	Sueldo = 2334.76
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ReturnString(num int) (bool, string) {
	var text string
	text = strconv.Itoa(num)
	return true, text
}
