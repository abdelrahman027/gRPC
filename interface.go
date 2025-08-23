package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	pirm() float64
}

type rect struct {
	width, hight float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.hight
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rect) pirm() float64 {
	return 2 * (r.width + r.hight)
}

func (c circle) pirm() float64 {
	return 2 * math.Pi * c.radius
}
func measure(s geometry) {
	fmt.Println("AREAM of shape is", s.area())
	fmt.Println("PIRM of shape is", s.pirm())
}

func getType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("The type of this input is INT")
	case string:
		fmt.Println("The type of this inpyt is STRING")
	default:
		fmt.Println("Unknown Type")
	}
}
func main() {
	rectangle := rect{width: 10, hight: 5}

	circ := circle{radius: 10}
	println("Rectangle Measures ============================")
	measure(rectangle)
	println("Circle Measures ===============================")
	measure(circ)
	fmt.Println("***************************************************")
	getType(13)
	getType("Hello")
	getType(rectangle)
}
