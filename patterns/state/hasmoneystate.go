package main

import (
	"fmt"
)

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (this *hasMoneyState) requestItem() error {
	return fmt.Errorf("Dispense item in progress")
}

func (this *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("Dispense item in progress")

}

func (this *hasMoneyState) insertMoney(amount int) error {
	return fmt.Errorf("Dispense item in progress")

}

func (this *hasMoneyState) dispenseItem() error {
	this.vendingMachine.itemCount -= 1
	fmt.Println("Item dispensed succesfully")
	if this.vendingMachine.itemCount == 0 {
		this.vendingMachine.setState(this.vendingMachine.noItem)
	} else {
		this.vendingMachine.setState(this.vendingMachine.hasItem)

	}
	return nil
}
