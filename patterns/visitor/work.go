package main

type Work struct {
	Workers uint
}

func (this Work) GetBuildingType() string {
	return "work"
}

func (this Work) accept(v visitor) {
	v.visitForWork(this)
}
