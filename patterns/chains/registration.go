package main

import (
	"fmt"
)

type registration struct {
	next department
}

func (this *registration) execute(person *patient) {
	if person.wasRegistration {
		fmt.Println("this person already was at registration")
	} else {
		person.wasRegistration = true
		fmt.Print("Registration for person done.")
		fmt.Println(person.name)
	}
	if isNotNull(this.next) {
		this.next.execute(person)
	} else {
		panic("next for registration is nil")
	}
}

func (this *registration) setNext(next_dep department) {
	this.next = next_dep
}
