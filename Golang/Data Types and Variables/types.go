package main
import (
	"fmt"
	"reflect"
)

func main(){
	var name string = "Hello world"
	var age int = 26
	var isGraduated bool = true
	var height float64 = 1.64

	fmt.Printf("variable name = %v is of type %T \n", name, name)
	fmt.Printf("variable age = %v is of type %T \n", age, age)
	fmt.Printf("variable isGraduated = %v is of type %T \n", isGraduated, isGraduated)
	fmt.Printf("variable height = %v is of type %T \n", height, height)

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Riajul"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000.3))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	fmt.Printf("variable age = %v is of type %v \n", age, reflect.TypeOf(age))
	fmt.Printf("variable name = %v is of type %v \n", name, reflect.TypeOf(name))
}