package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	var secuencia int
	return func() int {
		secuencia++
		return secuencia * numero
	}
}

func MyTabla(number int) {
	Closure := tabla(number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", number, i, Closure())
	}
}
