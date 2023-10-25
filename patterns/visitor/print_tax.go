package main

import (
	"fmt"
)

type PrintTax struct{}

func (this PrintTax) visitForHome(cur_home Home) {
	fmt.Println("You need to pay for home 0,1% anyally")
}

func (this PrintTax) visitForShop(cur_shop Shop) {
	fmt.Println("You need to pay 10-20% nds for your purchases")
}

func (this PrintTax) visitForWork(cur_work Work) {
	fmt.Println("You need to pay 13% any time you get salary ")
}
