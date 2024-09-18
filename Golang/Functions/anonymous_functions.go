package main

import "fmt"

// Anonymous functions are functions which are declared without any named identifiers to refer to it. They can accept inputs and return
// outputs

func main() {
	x := func(l int, b int) int {
		return l * b
	}

	fmt.Printf("%T \n", x)
	fmt.Println(x(20, 30))

}
