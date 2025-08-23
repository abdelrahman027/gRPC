package main

import "fmt"

func main() {

	// nums := []int{15, 4, 6}
	// for index, value := range nums {
	// 	if index == 1 {
	// 		fmt.Println("skip index 1")
	// 		continue
	// 	}
	// 	fmt.Printf("Index is %d and value is %d \n", index, value)
	// }
	// rows := 5
	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print("_")
	// 	}

	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	for i := range 10 {
		fmt.Println(i)
	}
}
