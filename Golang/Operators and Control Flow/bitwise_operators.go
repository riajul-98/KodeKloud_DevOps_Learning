package main
import "fmt"

func main(){
	// Work at bit level. The operators are & (bitwise and), | (bitwise or), ^ (bitwise xor), >> (right shift), << (left shift)

	/* Bitwise AND: takes 2 numbers as operands and does AND on every bit of two numbers. e.g. the binary value of 10 is 1010 and the
	binary value of 5 is 0011. Using bitwise AND will check if both the values are 1 and if they are, the result will be one. In this
	case, the binary value of the result will be 0010, which is 2.
	*/

	var x, y int = 12, 25
	z := x & y
	fmt.Println(z)

	/* Bitwise OR: Similar to bitwise AND but the value will be 1 if either value is 1. e.g. the binary value of 10 is 1010 and the
	binary value of 5 is 0011. Using bitwise OR will check if either value is 1 and if they are, the result will be one. In this case,
	the binary value for the result is 1011, which is 11.
	*/

	a := x | y
	fmt.Println(a)

	/* Bitwise XOR: The resulting output will be 1 if the two numbers are opposite e.g. the binary value of 10 is 1010 and the
	binary value of 5 is 0011. Using bitwise XOR will check if value for one of the values is 1 and the other is 0 and if they are, 
	the result will be one. In this case,the binary value for the result is 1001, which is 9.
	*/

	b := x ^ y
	fmt.Println(b)

	/* Bitwise Left shift: The left shift operator shifts all the binary values to the left by a specified number of bits. e.g. if we
	have 10 whose binary number is 1010 and we use the left shift operator to shift by one bit, the resulting value will be 10100 
	which is 20.
	*/

	var c int = 212
	d := c << 1
	fmt.Println(d)

	/* Bitwise Right shift: The left shift operator shifts all the binary values to the right by a specified number of bits. e.g. if we
	have 10 whose binary number is 1010 and we use the right shift operator to shift by one bit, the resulting value will be 0101 
	which is 5.
	*/

	var e int = 212
	f := e >> 1
	fmt.Println(f)


}