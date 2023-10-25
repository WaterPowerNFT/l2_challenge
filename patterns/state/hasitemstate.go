package main

import "fmt"

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (this *hasItemState) requestItem() error {
	if this.vendingMachine.itemCount == 0 {
		this.vendingMachine.setState(this.vendingMachine.noItem)
		return fmt.Errorf("No items in vending machine")
	}
	fmt.Printf("Item requested")
	this.vendingMachine.setState(this.vendingMachine.itemRequested)
	return nil
}

func (this *hasItemState) addItem(count int) error {
	this.vendingMachine.incrementItemCount(count)
	return nil
}

func (this *hasItemState) insertMoney(amount int) error {
	return fmt.Errorf("Please specify item at first")
}

func (this *hasItemState) dispenseItem() error {
	return fmt.Errorf("Please specify item at first")

}
