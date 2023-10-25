package main

import (
	"fmt"
)

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (this *itemRequestedState) requestItem() error {
	return fmt.Errorf("request already sent")
}

func (this *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("Error. Item dispensing in progress")

}

func (this *itemRequestedState) insertMoney(amount int) error {
	if amount < this.vendingMachine.itemPrice {
		return fmt.Errorf("please, insert %d money at least", this.vendingMachine.itemPrice)
	}
	fmt.Println("Money check is ok")
	this.vendingMachine.setState(this.vendingMachine.hasMoney)
	return nil
}

func (this *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please, insert de money")

}
