package main

type switcher struct {
	action command
}

func (this switcher) press() {
	this.action.perform()
}
