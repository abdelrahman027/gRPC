package main

import "fmt"

//TRY ADDING COMMENT
func Greeting(n string) {
	fmt.Println("Hello", n)
}

func CycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}
func main() {
	names := []string{"Ahmed", "Body", "Ali"}
	CycleNames(names, Greeting)
}
