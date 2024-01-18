package interface_practice

import (
	"fmt"

	IDENTinterface "github.com/SimonMora/go-from-scratch/interface"
)

func HumanBreathing(h IDENTinterface.Humano) {
	h.Breath()
	fmt.Printf("I'm a %s and I'm breathing.\n", h.Sex())
}
