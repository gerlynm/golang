package main

import "fmt"

func main() {
	fmt.Println("hello to ifelse")

	count := 10
	var result string

	if count < 20 {

		result = "regular user"

	} else if count > 50 {

		result = "not a regular user"

	} else {

		result = "something else"
	}

	fmt.Println("the result statement is", result)
}
