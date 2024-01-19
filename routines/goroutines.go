package routines

import (
	"fmt"
	"strings"
	"time"
)

func SlowlyShowedName(name string, chn1 chan bool) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}

	//channel was added to send the signal for ending the execution
	//async sencond class
	chn1 <- true
}
