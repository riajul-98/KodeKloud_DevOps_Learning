package main

import "fmt"

// declaring a pointer:	var <pointer_name> *<data_type>
// Initialising a pointter:	var <pointer_name> *<data_type> = &<variable_name>		You do not need to add the data type (optional),
// You can also use the shorthand operator (meaning you do not need to add the "var" keyword)

func main() {
	var i *int
	var s *string
	fmt.Println(i)
	fmt.Println(s)

	j := 10
	var ptr_j *int = &j
	fmt.Println(ptr_j)
}
