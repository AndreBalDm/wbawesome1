package wbawesome1

// Реализовать паттерн адаптер на любом примере.

import "fmt"

type Human struct{}

type Place interface {
	PeeInToilet()
}

func (a *Human) Pee(p Place) {
	p.PeeInToilet()
}

type Toilet struct{}

func (t *Toilet) PeeInToilet() {
	fmt.Println("aaaahhh")
}

type Park struct{}

func (t *Park) PeeOnTree() {
	fmt.Println("aaaahhh")
}

type ParkAdapter struct {
	park *Park
}

func (ma *ParkAdapter) PeeInToilet() {
	ma.park.PeeOnTree()
}

func Wb21() {
	human := &Human{}
	toilet := &Toilet{}

	human.Pee(toilet)

	park := &Park{}

	parkAdapter := &ParkAdapter{
		park: park,
	}

	human.Pee(parkAdapter)
}
