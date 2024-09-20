package main

import "fmt"

type s1 struct {
	x int
}

func main() {
	c := s1{x: 5}
	c1 := s1{x: 6}
	c2 := s1{x: 5}
	if c != c1 {
		fmt.Println("c and c1 have different values")
	}
	if c == c2 {
		fmt.Println("c is the same as c2")
	}
}
