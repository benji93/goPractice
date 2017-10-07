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

func main() {
	fmt.Println("-- Pointers --")
	fmt.Println("Example1")
	example1()
	fmt.Println("Example2")
	example2()
}