package wbawesome1

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Set(x float64, y float64) {
	p.x, p.y = x, y
}

func (p *Point) Get() (x float64, y float64) {
	return p.x, p.y
}

func (p1 *Point) Range(p2 *Point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func Wb24() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(10, 10)
	fmt.Println(p1.Range(p2))
	p3 := NewPoint(10, 10)
	p4 := NewPoint(0, 0)
	fmt.Println(p3.Range(p4))
}
