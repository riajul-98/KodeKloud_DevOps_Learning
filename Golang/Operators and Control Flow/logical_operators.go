package main
import "fmt"

func main(){
	var a int = 10
	fmt.Println((a < 100) && (a < 200))
	fmt.Println((a < 300) && (a < 0))

	fmt.Println((a < 100) || (a < 200))
	fmt.Println((a < 0) || (a > 200))

	var x, y int = 10, 20
	fmt.Println(!(x > y))
	fmt.Println(!(true))
	fmt.Println(!(false))
}