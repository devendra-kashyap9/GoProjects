package main

import (
	"fmt"
	//"math/rand"
	//"time"
)


// sending data
// func processNum(numChan chan int) {

// 	for num := range numChan{
// 		fmt.Println("Processing numChan", num )
// 		time.Sleep(time.Second)
// 	}
	
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }


// go routine synchronizer 

func task (done chan bool){
	defer func() {done <- true}()
	fmt.Println("Processing...")
}


func main(){

	// buffered channel or the unblocking
	emailChain := make(chan string, 100)

	emailChain <- "1@example.com"
	emailChain <- "2@example.com"

	fmt.Println(<-emailChain)
	fmt.Println(<-emailChain)


	// done := make(chan bool)
	// go task(done)

	// <- done    // blocking

	// receive
	// result := make(chan int)

	// go sum(result, 4, 5)
	// res := <- result
	// fmt.Println(res)

	//sending data 
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }
	

	//time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "ping"    // channels are the  blocking

	// msg := <- messageChan
	// fmt.Println(msg)

}
