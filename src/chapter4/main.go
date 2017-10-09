package main

// Variables

import "fmt"

var globalVar string = "Global"

func outString() {
	var x string 
	x = "Hello "
	fmt.Println(x)
	x += "second"
	fmt.Println(x)
	fmt.Println(globalVar)
}

func stringEquality() {
	var x, y string
	x = "hello"
	y = "world"
	fmt.Println(x == y)
	fmt.Println(globalVar)
}

func nameVar() {
	dogsName := "Max"
	fmt.Println("My dogs name is:", dogsName)
	fmt.Println(globalVar)
}

func constTest() {
	const myName string = "Ben"
	fmt.Println(myName)
}

func defineMultiVars() {
	var (
		a = 5
		b = 15
		c = 6
	)
	fmt.Println(a, b, c)
}

func multiplyByTwo() {
	var input float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)
	
	output := input * 2

	fmt.Println(output)
}

func assignPlus() {
	x := 5
	x += 1
	fmt.Println(x)
}

func fToC() {
	var input float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9
	fmt.Println(output) 
}

func feetToMeters() {
	var input float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &input)

	output := input * 0.3048
	fmt.Println(output)
}

func main() {
	fmt.Println("\n-- Output a string --")
	outString()
	fmt.Println("\n-- String Equality --")
	stringEquality()
	fmt.Println("\n-- Name a var --")
	nameVar()
	fmt.Println("\n-- Name Constant --")
	constTest()
	fmt.Println("\n-- Define multi vars --")
	defineMultiVars()
	fmt.Println("\n-- Multiply by two --")
	// multiplyByTwo()
	fmt.Println("\n-- Assign Plus --")
	assignPlus()
	fmt.Println("\n-- Fahrenheight to Celsius --")
	fToC()
	fmt.Println("\n-- Feet to meters --")
	feetToMeters()
}