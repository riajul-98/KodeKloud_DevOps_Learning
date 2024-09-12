package main
// package are Go's way of organising and reusing code. All go programs must start with a package declaration. Main is used to
// make the program into an executable program. 

import "fmt"

// Similar to import in Python. Used to import external packages. fmt is the format package which allows inputs and outputs. Package
// names need to be encapsulated in double quotes

func main(){
	// The building blocks of a Go program. The main function is a special type of function and it is the entrypoint of executable
	// program.
	fmt.Println("Hello World!")
}

// This is a single line comment

/* This is a
multi-line
comment
*/