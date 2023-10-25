package main

type Shop struct {
	workers   uint
	customers uint
}

func (this Shop) GetBuildingType() string {
	return "shop"
}

func (this Shop) accept(v visitor) {
	v.visitForShop(this)
}
