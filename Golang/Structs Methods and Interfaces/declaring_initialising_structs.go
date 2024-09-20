package main

import "fmt"

// user defined data type. A structure that groups together data elements. Provide a way to reference a series of grouped values
// through a single variable name. Used when it makes sense to group or associate two or more data variables.

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

func main() {
	var s Student
	fmt.Printf("%+v", s)

	st := new(Student)
	fmt.Printf("%+v", st)

	student_1 := Student{
		name:   "Joe",
		rollNo: 12,
	}
	fmt.Printf("%+v", student_1)
}
