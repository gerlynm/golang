package main

import "fmt"

func main() {
	fmt.Println("hello methods")

	values := User{"jeriii", 23, true}
	fmt.Println("print struct details", values)
	values.GetMeth()
}

type User struct {
	Name   string
	Age    int
	Online bool
}

func (u User) GetMeth() {

	fmt.Println("test struct with method", u.Age)

}
