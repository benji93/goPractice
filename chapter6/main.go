package main

import "fmt"

func array1() {
	var x [5]int
	x[0] = 66
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])
}

func arrayWithFor() {
	var x [5]float64
	x[0] = 10
	x[1] = 20
	x[2] = 30
	x[3] = 40
	x[4] = 50

	var total float64 = 0
	for i := 0; i < len(x); i ++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func arrayWithForSpecial() {
	x := [5]float64{
		10, 
		20, 
		30, 
		40, 
		50,
	}

	var total float64 = 0
	for _, value := range x { // Using the underscore means we don't need to set an iterator
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func simpleSlice() {
	// var x []float64 // []
	// x := make([]float64, 5) // [0, 0, 0, 0, 0]
	// x := make([]float64, 5, 10) // [0, 0, 0, 0, 0]
	// fmt.Println(x)

	arr := [5]float64{1,2,3,4,5} // [1, 2, 3, 4, 5]
	x := arr[0:3] // [1, 2, 3]
	fmt.Println(x)
}

func main() {
	fmt.Println("-- Array1 --")
	array1()
	fmt.Println("-- Array with Func --")
	arrayWithFor()
	fmt.Println("-- Special for --")
	arrayWithForSpecial()
	fmt.Println("-- Simple Slice --")
	simpleSlice()
}