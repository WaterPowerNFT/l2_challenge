package main

type AmericaHouse struct {
	WindowsType string
	DoorType    string
	NumFloors   int
}

func (this *AmericaHouse) setWindowType() {
	this.WindowsType = "Plastic window"
}
func (this *AmericaHouse) setDoorType() {
	this.DoorType = "Iron"
}
func (this *AmericaHouse) setNumFloor() {
	this.NumFloors = 2
}
func (this *AmericaHouse) getHouse() house {
	return house{
		WindowsType: this.WindowsType,
		DoorType:    this.DoorType,
		NumFloors:   this.NumFloors,
	}
}
