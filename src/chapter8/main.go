package main

import "fmt"

func zero(x int) {
	x = 0
}

func example1() {
	x := 5
	zero(x)
	fmt.Println(x) // Will stay as 5
}

func zeroPoint(xPtr *int) {
	*xPtr = 0
}

func example2() {
	x := 5
	zeroPoint(&x)
	fmt.Println(x) // Will change to zero
}

func one(xPtr *int) {
	*xPtr = 1
}

func newEx() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func getAddress(x *int) {
	fmt.Println(x)
}

func main() {
	// fmt.Println("-- Pointers --")
	// fmt.Println("Example1")
	// example1()
	// fmt.Println("Example2")
	// example2()
	// fmt.Println("\n-- New --")
	// newEx()
	fmt.Println("\n-- Swap --")
	x := 12
	y := 76
	fmt.Println("X =", x, "Y =", y)
	swap(&x, &y)
	fmt.Println("X =", x, "Y =", y)
	fmt.Println("\n-- Get pointer value --")
	px := 93
	getAddress(&px)
}