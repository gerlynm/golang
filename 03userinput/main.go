package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	rating := "please enter the rating"
	fmt.Println("quote is", rating)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)

}
