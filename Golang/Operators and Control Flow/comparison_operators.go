package main
import "fmt"

func main(){
	var city string = "Kolkata"
	var city_2 string = "Calcutta"
	fmt.Println(city == city_2)

	fmt.Println(city != city_2)

	var a, b int = 5, 10
	fmt.Println(a < b)

	var c int = 10
	fmt.Println(b <= c)

	fmt.Println(a > b)

	fmt.Println(a >= b)
}