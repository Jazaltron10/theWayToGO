package main

import (
	"fmt"
)

/*
What are Pointers?
Pointers are special types in go.
A pointer variable is a variable that stores memory location rather than values like integers or floats for example.
A pointer in Go is like a note that tells you where something is. It doesn't contain the actual data; it just points to where the data is stored in memory. Pointers are often used to work with data more efficiently and to share data between different parts of your code.

Pointer variable
var p = 0x140000b2020

Normal boring variable
var i = 123
var f = 2.3
To declare a pointer we use the * symbol, very similar to C++

NOTE
A pointer holds a memory address and never a value, it just holds the address where a value is stored.
It's like having a treasure map that shows you where the treasure (the value) is buried, but the map itself is not the treasure; it's just a guide to help you find it.
*/
func workingWithPointers() {
	//The variable p would hold the memory address of an int32 value.
	var num int = 42
	fmt.Printf("\nThis is the value of num b4 it gets pointed to -->- %d", num)
	var ptr *int = &num // ptr points to the memory location where num is stored
	*ptr = 100          // Changes the value of num to 100 through the pointer
	fmt.Printf("\nThis is the value of num after it gets pointed to and has it's value modified and then it's address -->- \n%d--> \n%v--> ", num, &num)
	fmt.Printf("\n%v-->\n%v-->\n%v-->\n\n", *ptr, ptr, &ptr)

	//A check to understand
	for i := 0; i < 200; i += 50 {
		num += i
		if num == *ptr {
			fmt.Printf("\nThey are the same value i.e num(%d) == *ptr(%d)\n", num, *ptr)
		} else {
			fmt.Printf("\nThey are not the same thing")
		}

		if &num == ptr {
			fmt.Printf("These are essentially the same address\n&num(%v) == ptr(%v)\n", &num, ptr)
			fmt.Println("ptr and num point to the same memory address")
		} else {
			fmt.Println("ptr and num point to different memory addresses")
		}

	}

}

func square(aray2 [5]float64) [5]float64 {
	fmt.Printf("\nThe Memory location of aray2 in the normal function is: %p", &aray2)
	for i := range aray2 {
		aray2[i] = aray2[i] * aray2[i]
	}
	return aray2
}
func squarePointer(aray2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe Memory location of aray2 in the pointer function is: %p", aray2)
	for i := range aray2 {
		aray2[i] = aray2[i] * aray2[i]
	}
	return *aray2
}

func main() {
	fmt.Println("Hello, Universe!")
	workingWithPointers()
	/*
		The * symbol has two functions
		The first is to initialize a new pointer and
		The second is used to reference the value of a pointer
	*/
	//Asterisk
	var p *int32 = new(int32)                          // Create a pointer p pointing to a new int32 with value 0
	var i int32                                        // Declare an int32 variable i with value 0
	fmt.Printf("\nThe value p points to is:-> %v", *p) // Prints 0 (value pointed to by p)
	fmt.Printf("\nThe value of i is:-> %v", i)         // Prints 0 (value of i)

	// Ampersand (&): Assign the address of i to p
	p = &i                                               // Now p points to i
	*p = 10                                              // Modify the value of i through p
	fmt.Printf("\n\nThe value p points to is:-> %v", *p) // Prints 10 (value pointed to by p)
	fmt.Printf("\nThe value of i is:-> %v", i)           // Prints 10 (value of i)

	var k int32 = 2
	i = k // Copy the value of k to i

	// SLICES
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Printf("\n%v", slice)
	fmt.Printf("\n%v", sliceCopy)

	//Functions
	var aray1 = [5]float64{1, 2, 3, 4, 5}
	// This Method uses double the memory and is therefore inefficient because it passes a copy of the aray1 to the function
	fmt.Printf("\n\nPass by value (Copying) in Functions")
	fmt.Printf("\nThe Memory location of aray1 is: %p", &aray1)
	var result1 [5]float64 = square(aray1)
	fmt.Printf("\nThe result1 is: %v", result1)
	fmt.Printf("\nThe value of aray1 is: %v", aray1)

	// This Method uses a pointer, which is more efficient as aray1 can directly be modified or worked on
	fmt.Printf("\n\nPass by reference (referencing with pointers) in Functions")
	fmt.Printf("\nThe Memory location of aray1 is: %p", &aray1)
	var result2 [5]float64 = squarePointer(&aray1)
	fmt.Printf("\nThe result2 is: %v", result2)
	fmt.Printf("\nThe value of aray1 is: %v", aray1)
}
