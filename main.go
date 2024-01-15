package main

import (
	"fmt"

	"github.com/SimonMora/go-from-scratch/variables"
)

func main() {
	var result, text = variables.ReturnString(1400)
	fmt.Println(result)
	fmt.Println(text)
}
