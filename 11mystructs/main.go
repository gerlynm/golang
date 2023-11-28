package main

import "fmt"

func main() {
	fmt.Println("hello structs")

	jerrii := User{"jerrii", 23, true}
	fmt.Printf("the name %+v\n", jerrii)

}

type User struct {
	Name string
	Age  int
	Pass bool
}
