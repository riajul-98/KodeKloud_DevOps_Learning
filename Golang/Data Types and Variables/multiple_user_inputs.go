package main
import "fmt"

func main(){
	var name string
	var muggle bool

	fmt.Println("What is your name and are you a muggle: ")
	fmt.Scanf("s% %t", &name, &muggle)
	fmt.Println(name, muggle)
}