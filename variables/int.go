package variables

import "fmt"

func ShowInts() {
	var normalInt int
	entero32 := int32(10)
	entero64 := int64(100)

	fmt.Println("Normal int = ", normalInt)
	fmt.Println("32 bit int = ", entero32)
	fmt.Println("64 bit int = ", entero64)
}
