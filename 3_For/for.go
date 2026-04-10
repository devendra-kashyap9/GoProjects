package main

import "fmt"

// for -> only loop avaiable in go
// for -> only construct available for looping

func main() {

	// while loop
	// i := 1
	// for i<=3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// Infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop
	// for i := 0; i<3; i++ {

	// 	fmt.Println(i)
	// }

	// for i := 0; i<=3; i++ {

	// 	if i == 2 {

	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range from 1.22

	for i := range 5 {

		fmt.Println(i)
	}

}
