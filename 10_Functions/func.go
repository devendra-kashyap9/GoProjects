package main

import "fmt"

// func add(a int, b int)int {
// 	return a+b
// }

func getLanguages() (string, string, string) {
	return "Golang", "Python", "Swift"
}

func add(x, y int) (z, z1 int) {

	// defer keyword
	defer fmt.Println("Hello!!") // it executes after the return hit
	z = x + y
	z1 = x - y
	fmt.Println("Before Return...")

	return
}

func main() {

	// result := add(5,3)
	// fmt.Println(result)

	// ouptput := getLanguages()
	//fmt.Println(getLanguages())

	// a, b := add(7, 9)
	// fmt.Println(a, b)


	// anonymous function
	/*test := func() {
		fmt.Println("Hello!! This is anonymous function")
	}
	test()*/

	// anothere way of calling and making it

	test := func(x int) int  {
		return x * -1
	}(9)
	fmt.Println(test)
}
