package main

import "fmt"

func main() {
	fmt.Println("hello functions")
	greeter()
	result := adder(5, 7)
	prores := proAdder(3, 2, 1, 4, 5, 3)
	fmt.Println("the addition is", result)
	fmt.Println("the PRoaddition is", prores)

}

func adder(val01 int, val02 int) int {

	return val01 + val02

}

func greeter() {
	fmt.Println("this is greeter function")
}

func proAdder(value ...int) int {

	count := 0
	for _, value := range value {
		count += value
	}

	return count

}
