package main

import (
	"fmt"
	"maps"
)

// maps -> hashable datatypes like; hash,dict

func main() {

	// creating the map
	// m := make(map[string]string) // [key]value

	// // initialisation
	// m["name"] = "Golang"
	// m["area"] = "Back-End"

	// fmt.Println(m)
	// fmt.Println(m["name"], m["area"])
	// fmt.Println(len(m))
	// //fmt.Println(cap(m))

	// // if key doesnot exist it return the zero value

	// // deleting the values
	// delete(m, "area")
	// fmt.Println(m)

	// // clearing the map
	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"Brand":1, "Price":12, "value":100}
	// fmt.Println(m)

	// // checking for the key existence
	// // we can use "_" underscore instead of a var name iif we are not using it

	// v, ok := m["Python"]
	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("OKAY!")
	// } else {
	// 	fmt.Println("Not Okay!")
	// }

	// comparing two maps
	m1 := map[string]int{"Brand": 1, "Price": 12, "value": 100}
	m2 := map[string]int{"Brand": 1, "Price": 12, "value": 100}

	fmt.Println(maps.Equal(m1, m2))
}
