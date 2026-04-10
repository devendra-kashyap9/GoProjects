package main

import (
	"fmt"
	"sync"
	//"time"
)

// go routines are kind of multithreading

func task(id int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("Doing task", id)
}

func main() {
	// creating waitgroup instead of sleep
	var wg sync.WaitGroup
	for i := 0; i<=10; i++ {
		wg.Add(1)
		go task(i, &wg)					// go routine

		// go func(i int){
		// 	fmt.Println(i)
		// }(i)
	}

	wg.Wait()

	//time.Sleep(time.Second * 2)
}
