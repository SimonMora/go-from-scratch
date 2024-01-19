package main

import (
	//"github.com/SimonMora/go-from-scratch/iterate"

	d "github.com/SimonMora/go-from-scratch/defer_panic"
	//"github.com/SimonMora/go-from-scratch/variables"
)

func main() {
	/*var result, text = variables.ReturnString(1400)
	fmt.Println(result)
	fmt.Println(text)*/

	/*if os := runtime.GOOS; os == "Linux." || os == "darwin" {
		fmt.Print("This is not Windows, this is ", os)
	} else {
		fmt.Print("This is Windows.")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("This is not Windows, this is Linux.")
	case "windows":
		fmt.Println("This is not Windows, this is Windows.")
	default:
		fmt.Printf("\nThis is: %s \n", os)
	}*/

	/*integer, message := practice.EvaluateInteger("15")
	fmt.Printf("The integer passed was: %d\n", integer)
	fmt.Println(message)*/

	/*teclado.ReadNumbersAndMultiply()*/

	//iterate.Iterate()

	//fmt.Println(practice.MultiplicationTable())

	//files.CreateTableFile()
	//files.SaveInTableFile()

	//files.FileReader()

	//funciones.Calculos()
	//funciones.MyTabla(7)

	//fmt.Println(funciones.Pow(4, 4))

	//slices.Capacity()
	//maps.ShowsMaps()

	/*u := new(model.User)
	u.User(1, "Simon Mora", time.Now(), true)

	fmt.Println(u)*/

	/*Juan := new(model.Man)
	interface_practice.HumanBreathing(Juan)

	Maria := new(model.Woman)
	interface_practice.HumanBreathing(Maria)*/

	//Using alias when import the defer_panic package
	//d.ShowDefer()
	//d.ShowPanic()
	d.ShowRecover()
}
