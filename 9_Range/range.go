package main

import "fmt"

func main() {
	// iterating over the data structure

	// nums := []int {1,2,3,4,5}

	// for i := 0; i<len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// nums := []int {1,2,3,4,5}
	// sum := 0

	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)

	// nums := []int {1,2,3,4,5}

	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// }

	// m := map[string]string {"fname":"John", "lname":"Doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// i is starting byte of rune 
	// c is unicode code point {"rune" data structure in go}
	
	for i, c := range "Golang v-24.5" {
		//fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
