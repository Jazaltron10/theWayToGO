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