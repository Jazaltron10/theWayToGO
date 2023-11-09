package main

import "fmt"

func Sum[T int | float32 | float64](numbers []T) T {
	var result T
	for _, num := range numbers {
		result += num
	}
	return result
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

type gasEngine struct {
	gallons float32
	mpg     float32
}
type electricEngine struct {
	kwh   float32
	mpkwh float32
}
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}
// func (c car[T]) carSpecifications() string{
	
// 	return fmt.Sprintf("Hello: %v", c.carMake)
// }
func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	fmt.Println(Sum(intSlice))   // Prints 15
	fmt.Println(Sum(floatSlice)) // Prints 16.5

	// var inSlice = []int{}
	fmt.Println(isEmpty[int](intSlice))
	fmt.Println(isEmpty[float64](floatSlice))

	// Generics With Structs
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Accord",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}
	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}
	fmt.Printf("\n->:%v",gasCar.carMake)
	fmt.Printf("\n->:%v",electricCar.carMake)
	
}
