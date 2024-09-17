package main
import "fmt"

func main(){
	var grades [3] int = [3]int{10,20,30}
	/* You can also use the shorthand methods: 
	grades := [3]int{10, 20, 30}
	grades := [...]int{10, 20, 30} */
	fmt.Println(grades)


	var fruits [2] string = [2]string{"apples", "oranges"}
	fmt.Println(len(fruits))

	marks := [3]int{10,20, 30}
	fmt.Println(marks[0])
	marks[1] = 100
	fmt.Println(marks)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	// Looping through an Array
	for i := 0; i < len(grades); i++{
		fmt.Println(grades[i])
	}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}


	// Multi-dimensional Arrays: nested arrays
	arr := [3][2]int{{2, 4}, {4, 16}, {8, 64}}
	fmt.Println(arr[2][1])
}