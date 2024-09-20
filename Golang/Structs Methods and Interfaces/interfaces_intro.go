package main

import "fmt"

// Interface specifies a method set and is a powerful way to introduce modularity in Go. Like a blueprint for a method set.
// Describe all the methods of a method set by providing the function signature for each method.
// They specifiy a set of methods but do not implement them.

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

type rect struct {
	length, width float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.width
}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rect{length: 3, width: 4}
	c := square{side: 5}
	printData(r)
	printData(c)
}
