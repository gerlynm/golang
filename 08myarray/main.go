package main

import "fmt"

func main() {

	fmt.Println("welcome to array")

	var nm [3]string

	nm[0] = "jerrii"
	nm[1] = "potter"
	nm[2] = "hermoine"

	fmt.Println(nm)
	fmt.Println(len(nm))

	var fruits = [5]string{"graphes", "orange", "lichi"}
	fmt.Println(fruits)

}
