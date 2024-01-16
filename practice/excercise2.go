package practice

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var err error

func MultiplicationTable() {
	scn := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Insert the number to evaluate the table: ")
		if scn.Scan() {
			number, err = strconv.Atoi(scn.Text())
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}
			break
		}
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}

}
