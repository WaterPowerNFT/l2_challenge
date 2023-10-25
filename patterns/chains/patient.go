package main

import (
	"fmt"
)

type patient struct {
	name            string
	wasRegistration bool
	wasDoctor       bool
	wasCashier      bool
}

func (this *patient) printName() {
	fmt.Println(this.name)
}
