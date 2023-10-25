package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func GetConcreteBuilder(coutry_name string) iBuilder {
	if coutry_name == "USA" {
		return &AmericaHouse{}
	} else if coutry_name == "Africa" {
		return &AfricaHouse{}
	}
	return nil
}
