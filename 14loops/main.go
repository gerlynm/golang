package main

import "fmt"

func main() {
	fmt.Println("hello to loops")

	//myarr := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// for i := 0; i < len(myarr); i++ {
	// 	fmt.Println("the value are", myarr[i])
	// }

	// for i := range myarr {

	// 	fmt.Println("the values are", myarr[i])

	// }

	// for index, value := range myarr {
	// 	fmt.Printf("the index  %v  values is %v\n", index, value)
	// }

	count := 2

	for count < 10 {

		if count == 5 {
			// count++
			// continue
			goto tiger
		}
		fmt.Println("the value is", count)
		count++
	}

tiger:
	fmt.Println("hai hello just miss")
}
