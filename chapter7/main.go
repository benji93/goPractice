package main

import "fmt"

func total(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func average(xs []float64) float64 {
	return total(xs) / float64(len(xs))
}

func main() {
	xs := []float64{98, 32, 77, 82, 89}

	fmt.Println(total(xs))
	fmt.Println(average(xs))
}