package main

import (
	"fmt"
)

type payment interface {
	Pay()
}

type card struct {
	number, cvc string
}

func (this *card) Pay() {
	fmt.Println("Pay for order by card")
}

func NewCard(number, cvc string) *card {
	return &card{number: number, cvc: cvc}
}

type sbp struct {
	number, bank string
}

func (this *sbp) Pay() {
	fmt.Println("Pay for order by sbp")
}

func NewSbp(number, bank string) *sbp {
	return &sbp{number: number, bank: bank}
}

type cash struct {
}

func (this *cash) Pay() {
	fmt.Println("Pay for order by cash")
}

func NewCash() *cash {
	return &cash{}
}

func MakePay(order_num string, payment_method payment) {
	fmt.Println("Processing payment")
	fmt.Printf("order num is %v\n", order_num)
	payment_method.Pay()
	fmt.Println()
}

func main() {
	my_card := NewCard("1111", "111")
	my_sbp := NewSbp("2222", "222")
	my_cash := NewCash()

	MakePay("first order", my_card)
	MakePay("second order", my_sbp)
	MakePay("third order", my_cash)

}
