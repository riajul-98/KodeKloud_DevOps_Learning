package main

import "fmt"

func addNumbers(a int, b int) int {
	sum := a + b
	return sum
}

func printGreeting(str string){
	fmt.Println("Hey there,", str)
}


func main(){
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)

	printGreeting("Riajul")
}