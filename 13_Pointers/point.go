package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	x := 7
	y := &x						// address of x # Reference
	fmt.Println(x,y)
	*y = 8           			// dereferencing     
	fmt.Println(x,y)

	/*num := 1
	changeNum(&num)

	fmt.Println("After changeNum", num)*/

	//fmt.Println("Memory address of num:", &num)

}
