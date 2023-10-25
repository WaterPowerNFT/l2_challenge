package main

type Home struct {
	Residents uint
}

func (this Home) GetBuildingType() string {
	return "home"
}

func (this Home) accept(v visitor) {
	v.visitForHome(this)
}
