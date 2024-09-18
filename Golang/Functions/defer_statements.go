package main

import "fmt"

// Defer statements delay the execution of a function until the surrounding function returns

func printName(str string) {
	fmt.Println(str)
}

func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {
	printName("Joe")
	defer printRollNo(23)
	printAddress("Street-32")
}
