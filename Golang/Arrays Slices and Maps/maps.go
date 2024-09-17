package main

import "fmt"

func main(){
	// Unordered collection of key/value pairs (similar to dictionaries in Python)
	// var code map[string]string		// This is a nil map, no data can be added to it
	codes := map[string]string{"en": "English", "fr": "French"}
	fmt.Println(codes)

	// declaring and initialising a map using the make() function
	codes_1 := make(map[string]int)
	fmt.Println(codes_1)

	// Length of Map
	fmt.Println(len(codes))

	// Accessing the value of a key in a map
	fmt.Println(codes["en"])

	value, found := codes["de"]
	fmt.Println(found, value)

	// Adding key-value pairs
	codes["it"] = "Italian"
	fmt.Println(codes)

	// Iterating over a map
	for key, value := range codes{
		fmt.Println(key, "=>", value)
	}

	// Updating key-pair values
	codes["en"] = "English Language"
	fmt.Println(codes)


	// Delete key-value pairs
	delete(codes, "en")
	fmt.Println(codes)

	// Truncate (clear) map
	for key, value := range codes{
		delete(codes, key)
	}
	fmt.Println(codes)
}