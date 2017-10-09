package main

import (
	"fmt"
	"math"
)

// First create the circle and rectangle classes
type Rectangle struct {
	x1, x2, y1, y2 float64
}

type Circle struct {
	x, y, r float64
}

type Shape interface {
	area() float64 
	perimeter() float64
}

func distance(a, b float64) float64 {
	return (b - a)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.x2)
	w := distance(r.y1, r.y2)
	return (l * w)
}

func (r Rectangle) perimeter() float64 {
	l := distance(r.x1, r.x2)
	w := distance(r.y1, r.y2)
	return (l*2 + w*2)
}

func (c Circle) area() float64 {
	return (math.Pi * c.r*c.r)
}

func (c Circle) perimeter() float64 {
	return (2 * math.Pi * c.r)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	r := Rectangle{0, 20, 0, 10}
	fmt.Println(r.area())
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println("\n-- Total Area --")
	fmt.Println(totalArea(&r, &c))

	fmt.Println("\n-- Perimeter --")
	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())
	fmt.Println(totalPerimeter(&r, &c))
}