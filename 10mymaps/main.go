package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello to maps")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["RB"] = "ruby"

	fmt.Println("the value of map is:", languages)

	delete(languages, "JS")
	fmt.Println("final after the delete:", languages)

	for key, value := range languages {

		fmt.Printf("the value of %v is %v\n", key, value)
		fmt.Printf("the value of %v is %v\n", key, value)

	}
}
