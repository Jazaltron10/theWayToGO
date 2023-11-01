package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func basicTypes() {
	// Great that you are doing this,
	// `var` is commonly defined public, unless its for a specific reason otherwise
	// like checking a nil type
	var username string = "wagslane"
	var password string = "20947382822"

	// don't edit below this line
	// why not use fmt.Printf(...), This would mean you can use /n (new line) or /t (tab)
	fmt.Printf("Authorization: Basic --> %s : %s\n", username, password) // should there be a space after `Basic`

	congrats := "Happy birthday!"
	fmt.Println(congrats)

	// Working with constants
	// pn - Nice this is also commonly used globally, but also can be used locally for a specific reason
	// like a internal infinate loop that checks a const value
	const premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)

	const firstName = "James"
	const lastName = "Hunger"
	// You can also use fmt.Sprintf("%s %s", firstName, lastName) this would be more ideal
	const fullName = firstName + " " + lastName

	fmt.Println("His name is", fullName)
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 12345678.9
	var intNum32 int32 = 10
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var uintNum1 uint = 3
	var uintNum2 uint = 2
	fmt.Println(uintNum1 / uintNum2)
	fmt.Println(uintNum1 % uintNum2)

	var myString string = `Hello Universe, i'm back \n`
	fmt.Println(myString)

	//Printing the length of a string
	fmt.Println(len("Microscopic"))                                                                      //This prints the number of bytes in this string
	fmt.Printf("\nThe number of characters in this string is %v", utf8.RuneCountInString("Microscopic")) //This prints the number of characters in this string
}

func workingWithStrings() {
	// fmt.Printf() - Prints a formatted string to standard out
	// fmt.Sprintf() - Returns the formatted string

	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("\nHi %s, your open rate is %.2f percent",
		name,
		openRate,
	)

	// pn - Nice when you have time look at the strings.Atio and others

	fmt.Println("\n", msg)

	//Strings, Runes, Bytes
	fmt.Printf("\n--->>Go Fast\n")

	var myString = "résumé"
	// var indexed = myString[0]
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	/*
		Understanding utf-8 encoding would help with getting a better grasp on how strings in golang work under the hood
		When you are dealing with strings In golang your dealing with a value whose underlying representation is an array of bytes.
	*/
	// This is why taking the length of a string is it's length in the number of bytes and not the number of characters

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	fmt.Printf("\n\n--->>- RUNES -<<---\n\n")
	// The easier way to deal with index and iterating a string is to cast it into an array of runes
	//They are just unicode point numbers which just represent the character.
	//They are also just an alias for int32

	var oString = []rune("résumé")
	var idx = oString[1]
	fmt.Printf("%v, %T\n", idx, idx)
	for i, v := range oString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nThe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	//Using StringBuilder
	fmt.Printf("\n\n--->>String Builder")
	var strSlice2 = []string{"j", "a", "c", "o", "b"}
	var strBuilder strings.Builder
	// The range keyword allows you to iterate over a string and access the index and character value in the string.
	for i := range strSlice2 {
		strBuilder.WriteString(strSlice2[i])
	}

	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)

	fmt.Printf("\n\n--->>Go Fast\n")

}

func workingWithConditionals() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	} else {
		fmt.Println("Message not Sent")
	}

	// Shortened if statement hack
	if age := 21; age > 18 {
		fmt.Println("\nYou are an adult!")
	}
}

func workingWithFunctions(x, y int) int {
	return x + y
}

func workingWithFuncStrings(s1 string, s2 string) string {
	// You can also use fmt.Sprintf("%s %s", s1, s2) this would be more ideal
	// fullName := s1 + " " + s2
	fullName := fmt.Sprintf("%s <-> %s", s1, s2)

	return fullName
}

func workingWithValues(x int) int {

	x++

	return x
}

func workingWithMultipleReturnValues() (string, string) {
	// A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore: _

	return "James", "Baldwin"
}

// This function demonstrates implicit and explicit return
func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {

	// pn - do you think this could also be a switch case?

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func test(age int) {
	fmt.Println("\nAge:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("==========================")
}

// Named return parameters are particularly important in longer functions with many return values.
// It is also advisable to only use implicit returns for small functions
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// Guard Clauses {Early Returns}
/*Guard Clauses Facilitate Multiple returns in a function at the different levels, helping to avoid nested conditional logic*/
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

// Errors
func workingWithErrors(num, den int) (res, rem int, err error) {

	if den == 0 {
		err = errors.New("cannot Divide by Zero")
		return 0, 0, err
	}
	res = num / den
	rem = num % den
	return res, rem, err
}

/*
Arrays, Slices, Maps and Looping Control Structures
.Fixed Length
.Same Type
.Indexable
.Contiguous in Memory
*/
func workingWithSlices() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Printf("->%d\n", intArr[0])
	//indexing(Accessing) elements(@index) 1 and 2 in the array
	fmt.Println(intArr[1:3])

	//Printing the Memory address of the elements in the array
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	//shortcut initialization
	var onArr [5]int32 = [5]int32{3, 4, 5, 6, 7}
	fmt.Printf("\n%v\n", onArr)
	//OR
	arr := [...]int32{13, 24, 35, 26, 17}
	fmt.Printf("\nAn alternative method to declaring and initializing arrays in go -> %v\n", arr)

	//SLICES
	/*
		Slices wrap arrays to give a more general,
		powerful, and convenient interface to sequences of data.
		They are simply arrays with added functionalities
		I think this is golang's version of dynamic arrays
	*/

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("\nThe array %v has a length of %v with capacity %v\n", intSlice, len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7, 3)
	fmt.Printf("\n%v", intSlice)
	fmt.Printf("\nThe array %v has a length of %v with capacity %v\n", intSlice, len(intSlice), cap(intSlice))

	//The Spread Operator
	// intSlice2 :=[...]int32{8,9}
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\ngo -> %v\n", intSlice)
	//OR use the make function, where you get to specify the length(3) and capacity(8) of the slice
	var intSlice3 []int32 = make([]int32, 3, 8)
	intSlice3 = append(intSlice3, intSlice...)
	fmt.Printf("\ngo -> %v\n", intSlice3)

}

func workingWithMaps() {
	/*
		It works on the same concept as like other programming languages, i.e key, value pairs ]
	*/
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	/*
		A map will always return something, even if the key doesn't exist
		Maps in go, also return an optional second value, which is a boolean
		e.g it returns true if the key(value) is in the map and false otherwise.
	*/
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}
	//To delete a value you, you use the delete function -> delete(map,"key")
	// delete(myMap2,"Adam")

	// Iterating over a map
	for name, age := range myMap2 {
		fmt.Printf("Name: %v\n", name)
		fmt.Printf("Name: %v, Age:%d \n", name, age)
	}

	//Iterating over Arrays and Slices
	intArr := []int32{2, 3, 5, 4, 6, 5, 3, 2, 4, 6}
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value:%d \n", i, v)
	}

	// While Loop variant
	var i, j int = 0, 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
	for j < 10 {
		if j >= 15 {
			break
		}
		fmt.Println(j)
		j = j + 1
	}

	// Traditional for loop
	for i := 0; i < 100; i += 10 {
		fmt.Printf("\n--->>> %d", i)
	}

}

func main() {
	basicTypes()
	workingWithStrings()
	workingWithConditionals()
	res_a := workingWithFunctions(2, 3)
	fmt.Println(res_a)
	res_b := workingWithFuncStrings("John", "Gab")
	fmt.Println(res_b)

	x := 5
	x = workingWithValues(x)
	fmt.Println(x)
	//prints 6 because the increment function received a copy of x, incremented it and then returned it.
	//Multiple Returns (Implicit and Explicit)
	firstName, _ := workingWithMultipleReturnValues()
	fmt.Println("Welcome to the Family Mi amor,", firstName)

	test(6)
	test(10)
	test(19)
	test(26)

	mul, div, err := calculator(23, 15)
	fmt.Printf("\nThe Multiplication of the two numbers produced is %d", mul)

	fmt.Printf("\nThe Division of the two numbers produced is %d", div)

	fmt.Println("\nThis is the error", err)

	//Guard Clauses
	res_c, err := divide(86, 3)
	fmt.Printf("\nWhen Divided the result is %d and the error is %v", res_c, err)

	//Errors
	fmt.Printf("\n\nThe Original Error Codes")
	var res_d, rem_d, erri = workingWithErrors(200, 25)
	if erri != nil {
		fmt.Printf("\n%v\n", erri.Error())
	} else if rem_d == 0 {
		fmt.Printf("\nThe result of the integer division is %v", res_d)
	} else {
		fmt.Printf("\nThe result of the integer division is %v with remainder %v", res_d, rem_d)
	}

	//Switch ALternative
	fmt.Printf("\n\nThe Switch Alternative")
	switch {
	case erri != nil:
		fmt.Println(erri.Error())
	case rem_d == 0:
		fmt.Printf("\nThe result of the integer division is %d\n", res_d)
		// fmt.Printf("\n%v", res_d)
	default:
		fmt.Printf("\nThe result of the integer division is %v with remainder %v", res_d, rem_d)
	}

	//Arrays and Slices
	workingWithSlices()

	//Maps
	workingWithMaps()
}
