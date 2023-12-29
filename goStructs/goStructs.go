package main

import "fmt"

type car struct {
	Make       string
	Model      string
	Height     int // from the ground
	Width      int
	Price      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

type truck struct {
	// 	"car" is embedded, so the definition of a
	// 	"truck" now also additionally contains all
	// 	of the fields of the car struct
	car
	power string
}

func canDriveOffroad(myCar car) bool {
	if myCar.FrontWheel.Radius < 20 || myCar.BackWheel.Radius < 21 {
		return false
	}
	if myCar.FrontWheel.Material != "offroad" || myCar.BackWheel.Material != "offroad" {
		return false
	}
	return true
}

func test(myCar car) {
	fmt.Printf("Checking car status...\nHeight: %d\nFront Wheel Radius: %d\nBack Wheel Radius: %d\nFront Wheel Material: %s\nBack Wheel Material: %s\n",
		myCar.Height,
		myCar.FrontWheel.Radius,
		myCar.BackWheel.Radius,
		myCar.FrontWheel.Material,
		myCar.BackWheel.Material,
	)
	fmt.Println()
	if canDriveOffroad(myCar) {
		fmt.Println("This car is ready for off-road driving")
	} else {
		fmt.Println("Only suitable for road usage")
	}
	fmt.Println("====================================")
}

// A rectangle struct
type rect struct {
	width  int
	height int
}

// Area has a receiver of (r rect)
func (r rect) area() int {
	fmt.Printf("\nThe area of this rectangle of width %d and height %d is -> ", r.width, r.height)
	return r.width * r.height
}

func main() {
	// Welcome Message
	sugar()

	test(car{
		Height: 15,
		FrontWheel: Wheel{
			Radius:   23,
			Material: "offroad",
		},
		BackWheel: Wheel{
			Radius:   24,
			Material: "offroad",
		},
	})

	test(car{
		Height: 5,
		FrontWheel: Wheel{
			Radius:   19,
			Material: "sports",
		},
		BackWheel: Wheel{
			Radius:   20,
			Material: "sports",
		},
	})

	test(car{
		Height: 10,
		FrontWheel: Wheel{
			Radius:   23,
			Material: "luxury",
		},
		BackWheel: Wheel{
			Radius:   23,
			Material: "luxury",
		},
	})

	test(car{
		Height: 10,
		FrontWheel: Wheel{
			Radius:   22,
			Material: "sedan-suv",
		},
		BackWheel: Wheel{
			Radius:   22,
			Material: "sedan-suv",
		},
	})

	jTruck := truck{
		power: "Electric",
		car: car{
			Make:   "Tesla",
			Model:  "Cybertruck",
			Height: 15, //in inches
			Width:  8,  // in feet
			Price:  80000,
		},
	}
	fmt.Println(jTruck.power)
	// embedded fields promoted to the top-level
	// instead of jTruck.car.make
	fmt.Println(jTruck.Make)
	fmt.Println(jTruck.Model)

	//Methods In Struct
	r := rect{
		width:  5,
		height: 10,
	}

	fmt.Println(r.area())
	// prints 50

}

func sugar() {
	fmt.Println("Welcome to the Car Test")
}
