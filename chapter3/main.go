package main

import (
	"fmt" // Format: used to output code
	"math" // Math used to find the power of two numbers
)

func boolean() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}

func strandint() {
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[1])
	fmt.Println("Hello " + "world")
}

// Find the largest number
func largestNum(num float64) {
	num = math.Pow(10, num)
	fmt.Println("You get:", num - 1)
}

func computeTwoNum(x, y int32) {
	var num = x * y
	fmt.Println(num);
}

func getStringLen(userString string) {
	var out = len(userString)
	fmt.Println(out);
}

func main() {
	strandint()
	fmt.Println("-- Boolean Values --")
	boolean()
	fmt.Println("-- Largest Number --")
	largestNum(2)
	fmt.Println("-- Compute two numbers --")
	computeTwoNum(32132, 42452)
	fmt.Println("-- Get string length --")
	getStringLen("The quick brown fox jummed over the lazy dog")
}