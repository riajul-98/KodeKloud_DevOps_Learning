package main

import "fmt"

// pointers are variables which hold memory addresses of another variable
// Address operator: You can find the address of a variable by preceding the name of the variable with an ampersand (&)
// Dereference operator: When placed before an address, it returns the value stored at that address. The operator is asteric (*)

func main() {
	x := 77
	fmt.Println(&x)
	fmt.Println(*(&x))
}
