package main

import "fmt"

//import "time"

func main() {

	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekdays")

	// }

	// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer dataType")
		case string:
			fmt.Println("String Type")
		case bool:
			fmt.Println("Bool dataType")
		default:
			fmt.Println("Others", t)
		}
	}

	whoAmI("GOLANG")
}
