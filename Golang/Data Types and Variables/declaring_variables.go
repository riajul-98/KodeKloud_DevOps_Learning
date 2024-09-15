package main
import "fmt"

func main(){
	var name string		// You are able to declare a variable without assigning it a value at the time of declaring.
	name = "Riajul"		// You can then assign the value later on.
	fmt.Println(name)


	var containers, language string = "docker", "Python"	// You can declare two variables of the same data type at once.

	var (
		university string = "QMUL"				// For different data types, you can declare them like this.
		grad_year int = 2020
	)

	fave_anime := "Naruto"		// Shorthand way of declaring a variable without needing to specify type.

}