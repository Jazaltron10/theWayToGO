package main

import (
	"errors"
	"fmt"
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
func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {

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
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}

// Guard Clauses {Early Returns}
/*Guard Clauses Facilitate Multiple returns in a function at the different levels, helping to avoid nested conditional logic*/
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
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
}
