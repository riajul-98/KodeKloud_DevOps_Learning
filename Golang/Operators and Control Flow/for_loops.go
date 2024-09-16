package main
import "fmt"

func main(){
	for i := 1; i <= 5; i++{
		fmt.Println(i * i)
	}


	for i := 1; i <= 5; i++{
		if i == 3 {
			break
		}
		fmt.Println(i)
	}


	for i := 1; i <= 5; i++{
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

}