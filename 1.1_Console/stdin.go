package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	// taking the user input in console 

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Type something for the Input")
	// scanner.Scan()
	// input := scanner.Text()   // storing the input value
	// fmt.Printf("You typed : %q\n", input)

	// by default scanner.Scan() takes the input in strin format so we need to type cast for any arithemetic
	//   op

	convScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your year of Birth:")
	convScanner.Scan()
	newInput, _ := strconv.ParseInt(convScanner.Text(), 10, 64)
	fmt.Printf("You are %d years old at the end of 2025 \n", 2025 - newInput)


}