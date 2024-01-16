package iterate

import "fmt"

func Iterate() {
	for i := 10; i > 0; i-- {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
