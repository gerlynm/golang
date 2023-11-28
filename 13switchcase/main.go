package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("hello to switch case")

	rand.Seed(time.Now().UnixMicro())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("this is case", diceNumber)

	case 2:
		fmt.Println("this is case", diceNumber)

	case 3:
		fmt.Println("this is case", diceNumber)
		fallthrough
	default:
		fmt.Println("Enter valid number")
	}

}
