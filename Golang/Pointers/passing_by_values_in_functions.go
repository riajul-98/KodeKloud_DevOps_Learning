package main

import "fmt"

// Passing by Value: Function is called by directly passing the value of the variable as an argument. Parameter is copied into another
// location of memory. Only copy is modified.

func modify(s string) {
	s = "World"
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}
