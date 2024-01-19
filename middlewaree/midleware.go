package middlewaree

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sustract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func Middleware() {
	fmt.Println("Initial request processing....")

	result := operateMidd(sum)(4, 4)
	fmt.Println(result)

	result = operateMidd(sustract)(4, 4)
	fmt.Println(result)

	result = operateMidd(multiply)(4, 4)
	fmt.Println(result)

	fmt.Println("End of the request processing...")
}

func operateMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("Begin the operation")
		return f(i1, i2)
	}
}
