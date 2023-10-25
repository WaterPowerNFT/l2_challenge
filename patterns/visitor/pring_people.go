package main

import (
	"fmt"
)

type PrintPeople struct{}

func (this PrintPeople) visitForHome(cur_home Home) {
	fmt.Println(fmt.Sprintf("Peoples at home: residents %v", cur_home.Residents))
}

func (this PrintPeople) visitForShop(cur_shop Shop) {
	fmt.Println(fmt.Sprintf("Peoples at shop: %v customers, %v workers", cur_shop.customers, cur_shop.workers))
}

func (this PrintPeople) visitForWork(cur_work Work) {
	fmt.Println(fmt.Sprintf("At work we can expect %v peoples", cur_work.Workers))
}
