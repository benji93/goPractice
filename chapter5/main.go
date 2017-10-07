package main

import "fmt"

func for1() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func for2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i);
	}
}

func ifState() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
}

func fizzBuzz() {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		switch {
			case (i % 3 == 0) && (i % 5 == 0): fmt.Println(i, "FizzBuzz") 
			case i % 3 == 0: fmt.Println(i, "Fizz")
			case i % 5 == 0: fmt.Println(i, "Buzz")
			default: fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("-- For1 --")
	for1()
	fmt.Println("\n-- For2 --")
	for2()
	fmt.Println("\n-- Odd/Even --")
	ifState()
	fmt.Println("\n-- FizzBuzz --")
	fizzBuzz()
}