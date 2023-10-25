package main

import (
	"example/shop"
	"example/user"
	"fmt"
)

// based on https://github.com/MaksimDzhangirov/go-patterns/blob/main/patterns/structural/facade/facade.md
func main() {
	Vasya := user.CreateUser("Vasya", 100000)
	Eldorado := shop.CreateShop("Eldorado", []string{"microwave", "laptop"}, []float64{100.0, 1000.0}, []int{100, 100})
	fmt.Println("Vasyas items before buying")
	Vasya.PrintUserItems()
	fmt.Println("Vasyas items after buying")
	Eldorado.BuyItem("microwave", 2, &Vasya)
	Vasya.PrintUserItems()
}
