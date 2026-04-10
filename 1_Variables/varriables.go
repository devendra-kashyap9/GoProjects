package main

import "fmt"

func main() {

	//var age int = 30
	//var name = "Golang"
	//var isAdult bool = true
	//name := "Golang"

	// var name string
	// name = "Golang"

	//var price float32 = 50.5

	//fmt.Println(price)
	//fmt.Printf("Type of variable price is %T\n", price)   // Printf and %T for getting the type of the var

	fmt.Printf("Hello %T %v\n", 10, 10)
	fmt.Printf("%e\n", 22341.434)
	fmt.Printf("%b\n", 234)

	fmt.Printf("%f\n", 234.435324345435346)
	fmt.Printf("%g\n", 234.435324345435346)   // %g for large float value DOUBLES

	fmt.Printf("%s\n", "Golang")
	fmt.Printf("%q\n", "Golang")              // %q for double quoted string

	fmt.Printf("%.2f\n", 234.435324345435346)
	fmt.Printf("%.f\n", 234.435324345435346)

	fmt.Printf("%12q is cool\n", "Golang")
	fmt.Printf("%-12q is cool\n", "Golang")

	fmt.Printf("%07d\n", 45)
	
	
}
