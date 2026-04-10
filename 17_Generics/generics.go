package main

import "fmt"

// func printSlice[T any](items []T) {    // [T any] is generics // instead of any we can use interface{}

// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T int | string | bool](items []T) {    // [T any] is generics // instead of any we can use interface{}

	for _, item := range items {
		fmt.Println(item)
	}
}

// to solve the duplicacy we will use the "GENERICS"

// func printStringSlice(items []string) {
// 	// continuous repetition of logic
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }


func main() {
	//nums := []int {1,2,3,4}
	//names := []string {"Golang", "Python", "Swift", "Ruby"}
	values := []bool {true, false, true, true}
	//printStringSlice(names)
	printSlice(values)
}
