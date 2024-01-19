package defer_panic

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("This is the first line..")
	defer fmt.Println("This is the final line...")

	fmt.Println("This is the second line..")
}

func ShowPanic() {
	a := 1

	if a == 1 {
		panic("The value 1 was found..")
	}
}

func ShowRecover() {
	/*It is require to use the recover function inside when defer
	otherwise it won't fallback on the panic event*/
	defer func() {
		recovery := recover()
		if recovery != nil {
			log.Fatalf("There was a fatal error that caused a Panic: \n %v", recovery)
		}
	}() //This parenthesis is required when using function on a defer

	a := 1

	if a == 1 {
		panic("The value 1 was found..")
	}
}
