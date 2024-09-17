package main
import "fmt"

func main(){
	/* Slices are a portion of data taken from an array. A slice has 3 major components, the pointer, the length and the capacity.
	The pointer is used to point to the index of the list for which you want the slice to start from. The length of a slice is the
	number of elements it contains. The capacity is the number of elements in the underlying array, starting from the first element
	of the slice.
	*/

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	sub_slice := slice[0:3]
	fmt.Println(sub_slice)

	slice_2 := make([]int, 5, 8)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

	// Changing the value of an index in a slice will also change the value in the underlying array.
	arr2 := [5]int{10, 20, 30, 40, 50}
	slice_3 := arr2[:3]
	fmt.Println(arr2)
	fmt.Println(slice_3)

	slice_3[1] = 9000
	fmt.Println(arr2)
	fmt.Println(slice_3)

	// Appending to a slice: If you append more values than that of the capacity, a new array will be created.
	new_array := [4]int{10, 20, 30, 40}
	new_slice := new_array[1:3]

	fmt.Println(new_slice)
	fmt.Println(len(new_slice))
	fmt.Println(cap(new_slice))

	new_slice = append(new_slice, 900, -80, 50)
	fmt.Println(new_slice)
	fmt.Println(len(new_slice))
	fmt.Println(cap(new_slice))

	// Appending a slice to another slice
	new_slice = append(new_slice, slice_3...)
	fmt.Println(new_slice)

	// Deleting element from a slice
	arr3 := [5]int{10, 20, 30, 40, 50}
	i := 2
	fmt.Println(arr3)
	slice1 := arr[:i]
	slice2 := arr[i+1:]
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)


	// Copying from a slice (must be the same data type)
	src_slice := []int{10, 20, 30, 40, 50}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("Number of elements copied: ", num)

	// Looping through a slice
	looping_arr := []int{10, 20, 30, 40, 50}
	for index, value := range looping_arr{		// If you do not need index, replace "index" with "_"
		fmt.Println(index, "=>", value)
	}
}