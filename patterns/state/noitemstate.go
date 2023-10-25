package main

import (
	"fmt"
)

type noItemState struct {
	vendingMachine *vendingMachine
}

func (this *noItemState) requestItem() error {
	return fmt.Errorf("no items in vending machine")
}

func (this *noItemState) addItem(count int) error {
	this.vendingMachine.incrementItemCount(count)
	this.vendingMachine.setState(this.vendingMachine.hasItem)
	return nil
}

func (this *noItemState) insertMoney(amount int) error {
	return fmt.Errorf("no items in vending machine")

}

func (this *noItemState) dispenseItem() error {
	return fmt.Errorf("no items in vending machine")

}
