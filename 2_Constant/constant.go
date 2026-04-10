package main

import "fmt"

const age = 30

func main() {

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
