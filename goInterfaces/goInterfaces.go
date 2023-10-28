package main

import (
	"fmt"
	"math"
)

/*
Interfaces in GO is just a collection of method signatures
*/
type shape interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	fmt.Println("Hello to you")
}
