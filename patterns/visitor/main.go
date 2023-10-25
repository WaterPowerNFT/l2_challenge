package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	shop := &Shop{customers: 100, workers: 10}
	work := &Work{Workers: 25}
	home := &Home{Residents: 5}

	people_printer := &PrintPeople{}
	people_printer.visitForHome(*home)
	people_printer.visitForShop(*shop)
	people_printer.visitForWork(*work)
	fmt.Println()
	tax_printer := *&PrintTax{}
	tax_printer.visitForHome(*home)
	tax_printer.visitForShop(*shop)
	tax_printer.visitForWork(*work)
	fmt.Println()
}
