package main
import "fmt"

func main(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hi there, %s", name)
}