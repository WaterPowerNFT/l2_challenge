package main

import (
	"fmt"
)

func main() {
	african_builder := GetConcreteBuilder("Africa")
	american_builder := GetConcreteBuilder("USA")

	my_director := CreateDirector(african_builder)
	first_house := my_director.BuildHouse()
	fmt.Println(first_house)
	my_director.ChangeCoutryBuilder(american_builder)
	second_house := my_director.BuildHouse()
	fmt.Println(second_house)
}
