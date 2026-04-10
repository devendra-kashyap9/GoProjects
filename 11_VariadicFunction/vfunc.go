package main

import "fmt"

func sum(nums ...int)int {                    // any type return we use "interface{}"

	total := 0

	for _, num := range nums {
		total += num
		//fmt.Println(total, num)
	}
	return  total
}

func main() {

	nums := []int{1,2,4,3,21}					  // packing
	results := sum(nums...)    					 // unpacking
	fmt.Println(results)
}
