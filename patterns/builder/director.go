package main

type Director struct {
	CountryBuilder iBuilder
}

func CreateDirector(builder iBuilder) Director {
	return Director{CountryBuilder: builder}
}

func (this *Director) ChangeCoutryBuilder(new_coutry_builder iBuilder) {
	this.CountryBuilder = new_coutry_builder
}
func (this *Director) BuildHouse() house {
	this.CountryBuilder.setDoorType()
	this.CountryBuilder.setNumFloor()
	this.CountryBuilder.setWindowType()
	return this.CountryBuilder.getHouse()
}
