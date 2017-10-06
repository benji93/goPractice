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

func main() {
	fmt.Println("-- Output a string --")
	outString()
	fmt.Println("-- String Equality --")
	stringEquality()
	fmt.Println("-- Name a var --")
	nameVar()
	fmt.Println("-- Name Constant --")
	constTest()
	fmt.Println("-- Define multi vars --")
	defineMultiVars()
}