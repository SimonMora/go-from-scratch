package slices

import "fmt"

var array []int = []int{10, 2, 3}
var arrangement [10]int = [10]int{7, 8, 23, 4, 56, 7, 8, 3, 12}

func ShowSlices() {
	fmt.Println(array)

	first := arrangement[:5]
	second := arrangement[3:7]
	third := arrangement[6:]

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}

func Capacity() {
	elements := make([]int, 5, 20)
	fmt.Printf("Length: %d and Capacity: %d\n", len(elements), cap(elements))

	new := make([]int, 0, 0)

	for i := 0; i <= 100; i++ {
		new = append(new, i)
		//fmt.Println(i % 10)
		if i%10 == 0 {
			fmt.Printf("Length: %d and Capacity: %d and i: %d\n", len(new), cap(new), i)
		}
	}

}
