# My Go Notes

## Table of Contents
- [Introduction](#introduction)
- [Syntax Basics](#syntax-basics)
- [Data Types](#data-types)
- [Control Flow](#control-flow)
- [Functions](#functions)
- [Packages and Modules](#packages-and-modules)
- [Error Handling](#error-handling)

## Introduction

- Welcome to my Go programming notes!
- This document serves as a reference for various Go topics.

## Syntax Basics
    


### Variables

- Declaring variables: `var x int`
- Assigning values: `x = 10`
- Short variable declaration: `y := 20`

### Data Types
<!-- The Number attached to the type is the number of bits memory allocated for that particular type -->
- Integer: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned Integer: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` <!-- can only represent positive numbers with this data type -->
- Byte: `byte` <!--alias for uint8 -->
- Rune: `rune` <!--alias for int32 represents a Unicode code point-->
- Float: `float32`, `float64`
- String: `string`
- Boolean: `bool`
- Complex Numbers: `complex64`, `complex128`

## Control Flow

### Conditional Statements

- `if`, `else if`, `else`
- `switch`, `case`, `default`
<!-- - The initial statement of an IF BLOCK -->

        if INITIAL_STATEMENT; CONDITION{

        }
<!-- Now note, the initial variable can only be acessed within the scope of this if block -->
### Loops

- `for` loops
- `range` for iterating over slices and maps

## Functions

- Declaring functions
- Function parameters and return values
- Function Signatures
- Anonymous functions (closures)
- Recursion

## Packages and Modules

- Importing packages
- Creating and organizing your own packages
- Working with Go modules

## Error Handling

- Using `error` for handling errors
- `panic` and `recover`

## Useful Resources

- [Official Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/welcome/1)


## General Notes
- `%T` : This tells the fmt.Printf function that the type of data type should be printed out
        
            fmt.Printf("The type of penniesPerText is %T\n", penniedPerText)


- `Guard clause` : An early return from a function when a given condition is met
---

> "Concurrency is not parallelism." - Rob Pike

<!-- git config --global init.defaultBranch <name> -->


## Pointers
        From what I understand pointers are used to manipulate data efficiently in go. They are used to store memory address of a variable that contains a value. 

        I also came across the concept of pass by value and pass by reference. This describes how data is transferred to functions.

        Pass by value is basically an inefficient method of managing data as here the data is being copied into a function with the original being unaffected by any operation carried out. It's like giving someone a photocopy of a book to make changes but the original book remains unchanged.

        Pass by reference is basically giving the original book to the person to make changes directly.

        Most basic data types, like numbers and strings are passed by value.
        Slices and Maps, on the other hand, are passed by reference, which means changes to these data structures inside a function affect the original data outside the function. 

        It is very important to understand pass by value and pass by reference as it can help in controlling how data is manipulated in programs and help manage memory efficiently.

 
 # Asterisk and Ampersand 
        * and &
 
        Creating a pointer to a variable is done using the & operator, and dereferencing a pointer (getting the value it points to) is done using *.

#####



## **What are `*` and `&` in Go?**

- In Go, `*` and `&` are special symbols used with pointers.
  - `*` is like a "magic arrow" that helps you see what's inside a box (variable).
  - `&` is like a "magic wand" that helps you find the exact location of a box (memory address).

## **Why Use Pointers in Go?**

- Pointers are used to:
  1. Make changes to data in one place and see those changes everywhere.
  2. Save memory when working with big data.
  3. Share data between different parts of your program.
  4. Avoid making copies of data, which can be slow and use up memory.

## **How to Use `*` and `&` in Go?**

1. **Creating Pointers with `&`:**
   - Use `&` to make a pointer that points to a variable.
   - Example:

   ```go
   var num int = 42
   var ptr *int = &num
   ```

   Here, `ptr` is like a "magic arrow" pointing to the number 42.

2. **Dereferencing with `*`:**
   - Use `*` to see what's inside a box that a pointer points to.
   - Example:

   ```go
   *ptr = 100
   ```

   This is like opening the box that `ptr` points to and changing the number inside.

## **Where to Use Pointers in Go?**

- Use pointers when you want to:
   - Make changes in one place and have those changes appear everywhere.
   - Work with big data without using up too much memory.

## **When to Use Pointers in Go?**

- Use pointers when you:
   - Need to change data inside a function and want those changes to show up outside the function.
   - Want to share data between different parts of your program.
   - Are dealing with big data and want to be efficient with memory.

## **Key Differences Between `*` and `&`:**

- `*` is like a "magic arrow" to see what's inside a box (dereference).
- `&` is like a "magic wand" to find the exact location of a box (get memory address).

## **Example with `*` and `&` in Go:**

```go
var num int = 42
var ptr *int = &num

*ptr = 100

// Now, num is 100, and ptr points to it!
```

In this example, `ptr` is like a "magic arrow" pointing to `num`, and `*ptr` lets us change the number inside `num`.
#####


## Pointers with Structs, Functions and Slices
        The difference between `p := &Person{Name: "Alice", Age: 30}` and `p := Person{Name: "Alice", Age: 30}` is the use of a pointer (`&`) in the first one and the absence of a pointer in the second one.

- `p := &Person{Name: "Alice", Age: 30}` creates a variable `p` that is a pointer to a `Person` object. In other words, `p` doesn't directly hold a `Person` but rather the memory address of a `Person` object.

- `p := Person{Name: "Alice", Age: 30}` creates a variable `p` that directly holds a `Person` object. It's not a pointer; it's the actual `Person` data.

In simple terms, the first one is like having a piece of paper with the address of a house written on it, and the second one is like having the house itself. With the pointer, you have a reference to the object, and without the pointer, you have the actual object.



Certainly! Here's the explanation of pointers with functions in Go, presented in a stylized Markdown format:

### Pointers with Functions in Go

Pointers are used in Go to pass data by reference, allowing you to modify the original data directly within functions. Understanding how `*` and `&` operators are used in different scenarios with functions is crucial for working with pointers.

1. **Passing Variables by Reference**

   When you need to modify a variable's value in a function and have those changes reflected outside the function, use pointers. 

   ```go
   func modifyData(data *int) {
       *data = *data * 2
   }
   
   func main() {
       value := 10
       modifyData(&value)
       fmt.Println(value) // Prints 20
   }
   ```

   In this example, `modifyData` takes a pointer to an `int`, and the `*` operator is used to dereference the pointer and modify the original `value`.

2. **Returning a Pointer from a Function**

   You can return a pointer from a function to expose data that was modified within it.

   ```go
   func createAndReturnPointer() *int {
       value := 42
       return &value
   }
   
   func main() {
       ptr := createAndReturnPointer()
       fmt.Println(*ptr) // Prints 42
   }
   ```

   The `createAndReturnPointer` function creates an `int` variable, returns a pointer to it, and that pointer is used to access the value in `main`.

3. **Modifying Structs with Functions**

   Pointers are commonly used with structs to pass them efficiently between functions and modify their values in place.

   ```go
   type Person struct {
       Name string
       Age  int
   }
   
   func changeName(p *Person, newName string) {
       p.Name = newName
   }
   
   func main() {
       p := &Person{Name: "Alice", Age: 30}
       changeName(p, "Bob")
       fmt.Println(p.Name) // Prints "Bob"
   }
   ```

   In this example, the `changeName` function takes a pointer to a `Person` and directly modifies the `Name` field within the original struct.

4. **Using Pointers with Slices**

   Slices are already references to arrays in Go, so you usually don't need pointers for slices. Modifying the slice itself doesn't require a pointer.

   ```go
   func modifySlice(slice []int) {
       slice[0] = 100
   }
   
   func main() {
       data := []int{1, 2, 3}
       modifySlice(data)
       fmt.Println(data[0]) // Prints 100
   }
   ```

   In this example, the `modifySlice` function modifies the original data, and there's no need for a pointer.

Pointers with functions are a powerful way to work with data in Go, allowing you to efficiently modify variables, structs, and other data structures. The `*` and `&` operators help you create, dereference, and pass pointers as needed for different use cases.



































































































