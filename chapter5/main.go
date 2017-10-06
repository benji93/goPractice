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

func main() {
	fmt.Println("-- For1 --")
	for1()
	fmt.Println("\n-- For2 --")
	for2()
}