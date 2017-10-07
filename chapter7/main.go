package main

import "fmt"

func total(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func sum(xs ...float64) float64 {
	sum := 0.0
	for _, v := range xs {
		sum += v
	}
	return sum
}

func average(xs []float64) float64 {
	return total(xs) / float64(len(xs))
}

func multi(xs []float64) (float64, float64) {
	return total(xs), average(xs)
}

func main() {
	xs := []float64{98, 32, 77, 82, 89}
	fmt.Println("-- Total and average --")
	fmt.Println(total(xs))
	fmt.Println(average(xs))
	fmt.Println("-- Return multiple values --")
	fmt.Println(multi(xs))
	fmt.Println("-- Variadic Function --")
	fmt.Println(sum(xs...))
}