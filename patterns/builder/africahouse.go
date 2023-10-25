package main

type AfricaHouse struct {
	WindowsType string
	DoorType    string
	NumFloors   int
}

func (this *AfricaHouse) setWindowType() {
	this.WindowsType = "Hole"
}
func (this *AfricaHouse) setDoorType() {
	this.DoorType = "Wood"
}
func (this *AfricaHouse) setNumFloor() {
	this.NumFloors = 1
}
func (this *AfricaHouse) getHouse() house {
	return house{
		WindowsType: this.WindowsType,
		DoorType:    this.DoorType,
		NumFloors:   this.NumFloors,
	}
}
