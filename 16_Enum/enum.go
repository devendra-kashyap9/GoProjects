package main

import "fmt"

// there is no inbuild enum type in Golang
// but we can implement using const

// type MyType string      custom defined dataTypes

//type OrderStatus int
type OrderStatus string

// const (
// 	Received OrderStatus = iota 
// 	Confirmed
// 	Prepared
// 	Delivered
// )

const (
	Received OrderStatus = "Received"       // if we will not give the type to all the var it will take
											//  the type of first one 
	Confirmed OrderStatus = "Confirmed"
	Prepared OrderStatus = "Prepared"
	Delivered OrderStatus = "Delivered"
)

func chnageOrderStatus(status OrderStatus) {
	fmt.Println("Changing orderStatus to ", status)
}

func main() {
	chnageOrderStatus(Confirmed)
}
