package main

import "fmt"

type gas_stove struct {
	is_working bool
}

func (this *gas_stove) on() {
	this.is_working = true
	fmt.Println("Gas stove turned on")
}

func (this *gas_stove) off() {
	this.is_working = false
	fmt.Println("Gas stove turned off")
}

func (this *gas_stove) stove_state() {
	if this.is_working {
		fmt.Println("Status checked. Gas stove is on")
	} else {
		fmt.Println("Status checked. Gas stove is off")
	}
}
