package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SimonMora/go-from-scratch/practice"
)

var fileName = "./files/txt/table.txt"

func CreateTableFile() {
	var text string = practice.MultiplicationTable()
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating the file: " + err.Error())
		return
	}

	fmt.Fprintln(file, text)
	file.Close()
}

func SaveInTableFile() {
	var text string = practice.MultiplicationTable()
	if !Append(fileName, text) {
		fmt.Println("Error on concatenation in the file.")
	}
}

func Append(fileName string, text string) bool {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erro trying to open the file: " + err.Error())
		return false
	}

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Erro trying to write the file: " + err.Error())
		return false
	}

	file.Close()
	return true
}

func FileReader() {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading the file: " + err.Error())
		return
	}

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		line := scn.Text()
		fmt.Println("> " + line)
	}
	file.Close()
}
