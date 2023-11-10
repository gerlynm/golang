package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()

	input := currentTime.Format("01-02-2006")
	fmt.Println(input)

	currentMonth := time.Date(2006, time.August, 12, 12, 5, 2, 4, time.UTC)
	fmt.Println(currentMonth)
}
