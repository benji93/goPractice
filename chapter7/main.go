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

func closure() {
	add := func(x, y int) int {
		return x + y
	} 
	fmt.Println(add(10, 12))
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func getEvenMain() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
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
	fmt.Println("-- Closure / Func within a Func --")
	closure()
	fmt.Println("-- Get Even Number --")
	getEvenMain()
	fmt.Println("-- Factorial --")
	fmt.Println(factorial(3))
}