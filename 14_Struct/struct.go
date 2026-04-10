package main

import (
	"fmt"
	//"go/types"
	"time"
)

//  order struct

type customer struct {
	name string
	phone string
}

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time					// nanosecond precision
	customer						   // struct embedding
}

// hardcoded constructor
// go inherently has no constructor

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup...
// 	myOrder := order {
// 		id : id,
// 		amount : amount,
// 		status : status,
// 	}
// 	return &myOrder
// }


// // receiver function

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }




func main() {

	// newCustomer := customer {
	// 	name : "John",
	// 	phone: "12342346546",
	// }

	newOrder := order {
		id :     "101",
		amount:  34,
		status:  "paid",
		customer: customer {
			name : "John",
			phone: "12342346546",
		},
	}

	newOrder.customer.name = "Robbin"
	fmt.Println(newOrder)

	// myOrder := order {
	// 	id : "1",
	// 	amount : 234.34,
	// 	status : "Paid",
	// }

	//myOrder.changeStatus("Delivered")
	//fmt.Println(myOrder)

	// myOrder.createdAt = time.Now()

	// fmt.Println("Order struct", myOrder)

	// myOrder := newOrder("123", 34.23, "Delivered")
	// fmt.Println(myOrder)

	// language := struct {
	// 	name string
	// 	isGood bool
	// } {"Golang", true}

	// fmt.Println(language)
	//fmt.Println(go/types(language))
	//fmt.Println(len(language))

}
