package main

import (
	"fmt"
)

type cashier struct {
	next department
}

func (this *cashier) execute(p *patient) {
	if p.wasCashier {
		fmt.Println("Person already was at cashier")
	} else {
		p.wasCashier = true
		fmt.Println("Paid")
	}
}

func (this *cashier) setNext(next_dep department) {
	this.next = next_dep
}
