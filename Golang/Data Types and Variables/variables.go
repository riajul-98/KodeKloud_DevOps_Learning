package main
import "fmt"

func main(){
	var name string = "Riajul"
	var age int = 26
	var height float64 = 1.64
	var tired bool = true

	fmt.Println(name) // Prints and then goes to a new line, ready for next command.
	fmt.Print(age, "\n") // Prints on the same line. You can print to a new line using the new line character; "\n"
	fmt.Println(height)
	fmt.Println(tired)

	fmt.Println("Hello, my name is ",name,", I am", age, "years old. I am ",height, "m tall. Am I always tired? ",tired)
	// String concatination can be done using this method.


	fmt.Printf("Hello, my name is %s and I am %d years old.", name, age)
	/* You can also concatinate strings using this method which is similar to f-strings in Python. %s is a format specifier. Other 
	format specifiers are:
	%v: formats the value in a default format
	%d: formats decimal integers
	%T: type of the value
	%s: Plain String
	%f: floaring number
	%t: true or false
	In the above example, the value of the name variable, replaces %s.
	See below for another example.
	*/

	fmt.Printf("Hi, my name is %v", name)
}