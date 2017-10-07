package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)
	return math.Sqrt(l * w)
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return (a*a + b*b)
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hello my is", p.Name)
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.r) // Get the value
	fmt.Println(c.area())

	r := Rectangle{0, 10, 0, 10}
	fmt.Println(r.area())

	fmt.Println("\n-- Android/Person --")
	a := &Android{
		Person: Person{
			Name: "Terminator",
		},
		Model: "T-1000",
	}
	a.Person.Talk()
}