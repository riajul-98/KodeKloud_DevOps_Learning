package main
import "fmt"

func main(){
	var a string = "happy"
	if a == "happy" {
		fmt.Println(a)
	} else if a == "Okay" {
		fmt.Println("Somewhat happy")
	} else {
		fmt.Println("Sad")
	}
}