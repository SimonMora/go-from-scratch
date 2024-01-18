package slices

import "fmt"

var table [10]int = [10]int{1, 2, 3, 4, 5, 6, 7}

func ShowArrays() {
	table[2] = 40
	table[5] = 22

	fmt.Println(table)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}

	fmt.Printf("Array length: %d, Array capacity: %d", len(table), cap(table))
}
