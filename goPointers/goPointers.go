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
func main() {
	fmt.Println("Hello, Universe!")
	workingWithPointers()
}
