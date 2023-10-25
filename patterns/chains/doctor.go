package main

import "fmt"

type doctor struct {
	next department
}

func (this *doctor) execute(p *patient) {
	if p.wasDoctor {
		fmt.Println("Patient already was at doctor")
	} else {
		p.wasDoctor = true
		fmt.Println("Patient is done with doctor")
	}
	if isNotNull(this.next) {
		this.next.execute(p)
	} else {
		panic("next for doctor is nil")
	}
}

func (this *doctor) setNext(next_dep department) {
	this.next = next_dep
}
