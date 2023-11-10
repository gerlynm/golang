package main

import "fmt"

func main() {
	fmt.Println("welcome to pointer")

	// var ptr *int
	// fmt.Println(ptr)

	mynumber := 23

	var ptr = &mynumber

	fmt.Println("value of actual pointer is", ptr)
	fmt.Println("actual value inside of pointer", *ptr)

	*ptr = *ptr + 2
	fmt.Println("new value is", mynumber)
}
