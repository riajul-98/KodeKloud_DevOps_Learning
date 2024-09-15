package main
import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n", f)

	var new_f float64 = 45.24
	var new_i int = int(new_f)
	fmt.Printf("%v\n", new_i)

	/* strconv package

	Itoa():
	    Converts integer to string
		Returns one value - String formed with the given integer

	Atoi()
	    Converts string to integer
		Returns two values - The corresponding integer and error (if any)
	*/

	var int_to_convert int = 42
	var new_string string = strconv.Itoa(int_to_convert)
	fmt.Printf("%q", new_string)

	var string_to_convert string = "200"
	new_int, err := strconv.Atoi(string_to_convert)
	fmt.Printf("%v, %T \n", new_int, new_int)
	fmt.Printf("%v, %T", err, err)
}