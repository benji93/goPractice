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

func sliceAppend() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice2)
}

func sliceCopy() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2) // Makes an array with a length of two
	copy(slice2, slice1) // Copy slice1 into slice 2
	fmt.Println(slice1, slice2)
}

func simpleMap() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	y[2] = 11
	y[3] = 13
	fmt.Println(y)
	
	delete(y, 2)
	fmt.Println(y)
}

func multiMap() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hyrdrogen",
			"state": "Gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "Gas",
		},
		"Li": map[string]string{
			"name": "Lithium",
			"State": "Solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el)
	}
	
	// fmt.Println(elements)
}

func getSmallestNum() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	var lowest int = x[0] // Set lowest to the first value in the array
	for _, value := range x {
		if value < lowest {
			lowest = value
		}
	}

	fmt.Println(total)
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
	fmt.Println("-- Slice Append --")
	sliceAppend()
	fmt.Println("-- Slice Copy --")
	sliceCopy()
	fmt.Println("-- SimpleMap --")
	simpleMap()
	fmt.Println("-- MultiMap --")
	multiMap()
	fmt.Println("-- Get smallest value --")
	getSmallestNum()
}