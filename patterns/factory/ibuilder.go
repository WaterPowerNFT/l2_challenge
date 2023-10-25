package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse()
}

func GetConcreteBuilding(coutry_name string) iBuilder {
	if coutry_name == "USA" {
		america_house := &AmericaHouse{}
		america_house.getHouse()
		return america_house
	} else if coutry_name == "Africa" {
		africa_house := &AfricaHouse{}
		africa_house.getHouse()
		return africa_house
	}
	return nil
}
