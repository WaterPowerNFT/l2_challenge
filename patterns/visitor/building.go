package main

type building interface {
	GetBuildingType() string
	accept(visitor)
}
