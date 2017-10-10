package main

import "fmt"

type testValues struct {
	values []float64
}

type testAverage struct {
	testValues
	average float64
}

var values = []testValues{
	{ []float64{1,2} },
	{ []float64{1,1,1,1,1,1} },
	{ []float64{-1,1} },
}



func main() {
	fmt.Println(values)
}