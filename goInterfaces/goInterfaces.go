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

// Another take on structs and Interfaces
type gasEngine struct {
	mpg     uint32 //Miles per gallon
	gallons uint32 //How many gallons of fuel are left
	// ownerInfo owner or
	owner
}
type electricEngine struct {
	mpkwh uint32 //miles per kilowatt hour
	kwh   uint32 //specifies how much charge is left in the battery
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint32 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint32 {
	return e.kwh * e.mpkwh
}

/*
Currently the canMakeIt function can only take in the gasEngine struct type, to make it general we'll have to use an interface, this way both the gasEngine and electricEngine would be supported
*/
type engine interface {
	milesLeft() uint32
}

// Test function to check the functionality of the method
func canMakeIt(e engine, miles uint32) {
	if miles <= e.milesLeft() {
		fmt.Println("\nYou can make it there!")
	} else {
		fmt.Println("\nNeed to fuel up first!")
	}
}

func main() {
	fmt.Println("Hello to you")
	// Engines
	var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	fmt.Printf("\nmyEngine.milesLeft()-->- %d\n\n", myEngine.milesLeft())

	//Anonymous structs
	//These types of structs are not reusable
	var myEngine2 = struct {
		mpg     uint32
		gallons uint32
	}{25, 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	//interface
	var myEngine3 gasEngine = gasEngine{25, 20, owner{"James"}}
	var myEngine4 electricEngine = electricEngine{25, 2}

	canMakeIt(myEngine3, 50)
	canMakeIt(myEngine4, 80)
}
