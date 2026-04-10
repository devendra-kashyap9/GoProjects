package main

import (
	"fmt"
)

// slices are the dynamic array of Go
// most used concept
// + useful method

func main() {

	// unintialized slices is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println((nums==nil))

	// making non-nill slice using "make" fn
	//var nums = make([]int, 0, 5)   // 2 and 5 are the initial size and the initial capacity of the slice
	// initial size also initialises the 'given number of elemnts(size) with "0"'
	// here 2 will always keep first 2 elements starting with 0
	// to prevent this we give the size "0"

	// capacity -> max number of elements can fit
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)

	// fmt.Println(nums)
	// fmt.Println(cap(nums))      // capacity of the slice
	// fmt.Println(len(nums))
	//fmt.Println(nums==nil)

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)

	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// slicing in slice using slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:])

	// comparing the slice
	// num1 := []int{1,2}
	// var num2 = []int{1,2}

	// fmt.Println(slices.Equal(num1, num2))

	// 2d slices
	var mum = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(mum)
}
