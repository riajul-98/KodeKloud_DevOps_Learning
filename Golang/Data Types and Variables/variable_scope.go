package main
import "fmt"
var name string = "Lisa"  // Global variable. Can be accessed anywhere in code.

func main(){
	// Outer block
	city := "London"
	{
		// Inner block
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
		// Inner blocks can access variables from outer blocks
	}
	// fmt.Println(country)
	// Outer blocks cannot access variables from inner blocks.



	// Local vs Global variables
	name := "Lisa" // Local variable as it is created inside a function and can only be accessed in the function.
	
}