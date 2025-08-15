package basics

import (
	"fmt"
)

func basics() {
	var a, b float64 = 10, 3
	var result float64

	result = a + b
	fmt.Println("Add", result)

	result = a - b
	fmt.Println("Subs", result)

	result = a * b
	fmt.Println("Mult", result)

	result = a / b
	fmt.Println("Div", result)

}
