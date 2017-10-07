package main

import (
	"fmt"
	// "math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c)
}