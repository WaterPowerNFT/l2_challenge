package main

func isNotNull(d department) bool {
	if d == nil {
		return false
	}
	return true
}

func main() {
	doc := &doctor{}
	cash := &cashier{}
	reg := &registration{}

	vasya := &patient{name: "Vasiliy"}

	reg.setNext(doc)
	doc.setNext(cash)
	reg.execute(vasya)
}
