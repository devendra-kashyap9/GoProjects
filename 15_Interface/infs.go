package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

type payment struct {
	gateway paymenter       // using gateway property
}

// open- close principle
// which states code should be open for expansion but closed for modificatuon
func (p payment) makePayment(amount float32) {
	//razorpayPaymentGW := razorpay{}
	//stripePaymentGw := stripe{}
	//razorpayPaymentGW.pay(amount)
	//stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
} 

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	// logic to make payment 
	fmt.Println("Making payment using razorpay..", amount)
}

// type stripe struct {}

// func (s stripe) pay(amount float32){
// 	fmt.Println("Making oayment using Stripe", amount)
// }

// type fakepayment struct {}

// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("payment through fakePayment for testing purpose")
// }

type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using paypal", amount)
}

func main() {
	// stripePaymentGw := stripe{}
	//razorpayPaymentGW := razorpay{}
	//fakeGW := fakepayment{}
	paypalGW := paypal{}
	newPayment := payment{
		gateway: paypalGW,
	}
	newPayment.makePayment(100)

}
