package main

import (
	"fmt"
)

func main() {
	my_vending := newVendingMachine(0, 10)
	fmt.Println("**********noItemState*************")
	err := my_vending.insertMoney(20)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.dispenseItem()
	if err != nil {
		fmt.Println(err)
	}

	err = my_vending.requestItem()
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.addItem(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("**********hasItemState*************")

	err = my_vending.addItem(1)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.insertMoney(20)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.dispenseItem()
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.requestItem()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("**********itemRequested*************")
	err = my_vending.addItem(1)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.dispenseItem()
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.requestItem()
	if err != nil {
		fmt.Println(err)
	}

	err = my_vending.insertMoney(5)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.insertMoney(25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("**********hasMoney*************")

	err = my_vending.addItem(1)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.insertMoney(20)
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.requestItem()
	if err != nil {
		fmt.Println(err)
	}
	err = my_vending.dispenseItem()
	if err != nil {
		fmt.Println(err)
	}
}
