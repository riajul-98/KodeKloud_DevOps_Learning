package main
import "fmt"

func main(){
	var i int = 100
	switch i {
		case 10:
			fmt.Println("i is 10")
		case 100, 200:
			fmt.Println("i is either 100 or 200")
			fallthrough		// Goes through the through next case rather than skipping.
		default:
			fmt.Println("i is neither 0, 100 or 200")
	}


	var a, b int = 10, 20
	switch{
	case a + b == 30:
		fmt.Println("Equal to 30")
	case a + b <= 30:
		fmt.Println("Less than or equal to 30")
	default:
		fmt.Println("Greater than 30")
	}
}