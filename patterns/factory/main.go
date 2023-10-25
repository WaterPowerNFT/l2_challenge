package main

import (
	"fmt"
)

func main() {
	african := GetConcreteBuilding("Africa")
	american := GetConcreteBuilding("USA")

	fmt.Println(african)
	fmt.Println(american)
}
