package main

import "fmt"

func boolean() {
	fmt.Println(true && true)
}

func strandint() {
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[1])
	fmt.Println("Hello " + "world")
}

func main() {
	strandint()
	boolean()
}