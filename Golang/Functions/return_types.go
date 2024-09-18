package main
import "fmt"

// func operation(a int, b int) (int, int){
// 	sum := a + b
// 	diff := a - b
// 	return sum, diff
// }

func operation(a int, b int) (sum int, diff int){
	// named return values: If we place the returned values in the parameter section above, we do not need to add it to the 
	// return statment below and we do not need to use the shorthand declaration to declare a variable.
	sum = a + b
	diff = a - b
	return
}

func sumNumbers(numbers ...int) int {
	// Variadic functions allow you to add as many arguments as you want. These are initiated by using the elipses in the 
	// parameter field of the function declaration as above.
	sum := 0
	for _, value := range numbers {		
		// The underscore is known as the blank identifier an is used to ignore the values returned by functions
		sum += value
	}
	return sum
}

func main(){
	// Returning multiple values
	sum, difference := operation(20, 10)
	fmt.Println(sum, " ", difference)

	
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10, 20))
	fmt.Println(sumNumbers(10, 20, 30, 40, 50))
}