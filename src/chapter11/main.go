package main

import "fmt"
import "chapter11/math"

func main() {
	xs := []float64{10, 20, 30, 40}
	avg := math.Average(xs)
	fmt.Println(avg)
	min := math.Min(xs)
	fmt.Println(min)
	max := math.Max(xs)
	fmt.Println(max)
}