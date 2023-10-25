package shop

import (
	"errors"
	"example/user"
)

func (this Shop) CheckAvailability(item_name string, quantity int) error {
	val, ok := this._goods_remains[item_name]
	if !ok {
		return errors.New("Havent this item in shop")
	}
	if val < quantity {
		return errors.New("Not enough items in shop")
	}
	return nil
}

func (this *Shop) BuyItem(item_name string, quantity int, buyer *user.User) error {
	//check availability
	err := this.CheckAvailability(item_name, quantity)
	if err != nil {
		return err
	}
	//get amount of money needed
	total_cost := float64(quantity) * this._goods_price[item_name]
	//check user balance
	err = buyer.CheckBalance(total_cost)
	if err != nil {
		return err
	}
	//minus money, add
	buyer.Items[item_name] += quantity
	buyer.Money -= total_cost
	this._goods_remains[item_name] -= quantity
	return nil
}
func CreateShop(name string, goods_list []string, goods_prices []float64, remains []int) Shop {
	if len(goods_list) != len(goods_prices) || len(goods_list) != len(remains) {
		panic("wrong length of items")
	}
	to_ret := new(Shop)
	to_ret._name = name
	to_ret._goods_price = make(map[string]float64)
	to_ret._goods_remains = make(map[string]int)
	for i := 0; i < len(goods_list); i += 1 {
		to_ret._goods_price[goods_list[i]] = goods_prices[i]
		to_ret._goods_remains[goods_list[i]] = remains[i]
	}
	return *to_ret
}

type Shop struct {
	_name          string
	_goods_remains map[string]int
	_goods_price   map[string]float64
}
