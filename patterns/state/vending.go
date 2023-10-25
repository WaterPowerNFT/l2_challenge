package main

type vendingMachine struct {
	hasItem       state
	noItem        state
	hasMoney      state
	itemRequested state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(count int, price int) *vendingMachine {
	vending_unit := &vendingMachine{itemCount: count, itemPrice: price}
	hasItemStateUnit := &hasItemState{vendingMachine: vending_unit}
	itemRequestedStateUnit := &itemRequestedState{vendingMachine: vending_unit}
	noItemStateUnit := &noItemState{vendingMachine: vending_unit}
	hasMoneyStateUnit := &hasMoneyState{vendingMachine: vending_unit}
	vending_unit.setState(hasItemStateUnit)
	vending_unit.hasItem = hasItemStateUnit
	vending_unit.noItem = noItemStateUnit
	vending_unit.itemRequested = itemRequestedStateUnit
	vending_unit.hasMoney = hasMoneyStateUnit
	return vending_unit
}

func (this *vendingMachine) requestItem() error {
	return this.currentState.requestItem()
}

func (this *vendingMachine) addItem(count int) error {
	return this.currentState.addItem(count)
}

func (this *vendingMachine) insertMoney(amount int) error {
	return this.currentState.insertMoney(amount)
}

func (this *vendingMachine) dispenseItem() error {
	return this.currentState.dispenseItem()
}

func (this *vendingMachine) setState(new_state state) error {
	this.currentState = new_state
	return nil
}

func (this *vendingMachine) incrementItemCount(count int) error {
	this.itemCount += count
	return nil
}
