package main

type visitor interface {
	visitForShop(building)
	visitForHome(building)
	visitForWork(building)
}
