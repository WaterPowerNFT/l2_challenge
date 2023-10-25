package main

func main() {
	gas_stove := &gas_stove{}
	on_commando := &onCommand{gas_stove}
	off_commando := &offCommand{gas_stove}

	on_switcher := &switcher{on_commando}
	off_switcher := &switcher{off_commando}

	on_switcher.press()
	off_switcher.press()

}
