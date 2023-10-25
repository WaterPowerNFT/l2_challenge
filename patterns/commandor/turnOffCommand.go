package main

type offCommand struct {
	device device
}

func (this *offCommand) perform() {
	this.device.off()
}
