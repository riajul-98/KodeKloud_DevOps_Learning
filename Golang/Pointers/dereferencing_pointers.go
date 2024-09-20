package main

import "fmt"

func main() {
	s := "Hello"
	ps := &s
	fmt.Printf("%T %v \n", s, s)
	*ps = "World"
	fmt.Printf("%T %v \n", s, s)
}
