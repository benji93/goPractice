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

func first() {
	fmt.Println("This is the first func")
}

func second() {
	fmt.Println("This is the second func")
}

func deferEx() {
	defer second()
	first()
}

func panicEx() {
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	panic("Panic")
}

func half(x int) (int, bool) {
	y := false
	if (x%2 == 0) {
		y = true
	} else {
		y = false
	}
	return x / 2, y
}

func highestNum(x ...int) int {
	heighest := 0
	for _, v := range x {
		if v > heighest {
			heighest = v
		}
	}
	return heighest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fibonacci(n uint) uint {
	if n < 2 {
		return n	
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func problems() {
	fmt.Println("-- Half --")
	fmt.Println(half(3))
	fmt.Println("-- Variadic --")
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	fmt.Println(highestNum(x...))
	fmt.Println("-- Make Odd Generator --")
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())	
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("-- Fibonacci --")

	// Loop through the fibonacci sequence
	var n uint
	for n = 0; n < 50; n++ {
		fmt.Println(fibonacci(n))
	}
	
}

func main() {
	// xs := []float64{98, 32, 77, 82, 89}
	// fmt.Println("-- Total and average --")
	// fmt.Println(total(xs))
	// fmt.Println(average(xs))
	// fmt.Println("-- Return multiple values --")
	// fmt.Println(multi(xs))
	// fmt.Println("-- Variadic Function --")
	// fmt.Println(sum(xs...))
	// fmt.Println("-- Closure / Func within a Func --")
	// closure()
	// fmt.Println("-- Get Even Number --")
	// getEvenMain()
	// fmt.Println("-- Factorial --")
	// fmt.Println(factorial(3))
	// fmt.Println("-- Defer --")
	// deferEx()
	// fmt.Println("-- Panic --")
	// panicEx()
	problems()
}