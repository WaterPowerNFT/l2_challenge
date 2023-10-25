package main

type onCommand struct {
	device device
}

func (this *onCommand) perform() {
	this.device.on()
}
