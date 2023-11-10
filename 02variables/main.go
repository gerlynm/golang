package main

import "fmt"

var parentName string = "GOD"

const Blood string = "B+"

func main() {
	var name string = "jerrii"
	fmt.Println(name)
	fmt.Printf("variable type: %T \n", name)

	var pno int = 9445178314
	fmt.Println(pno)
	fmt.Printf("variable type: %T \n", pno)

	var isEducated bool = true
	fmt.Printf("variable type: %T \n", isEducated)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("variable type: %T \n", smallVal)

	var temp float32 = 101.7899364973246923
	fmt.Println(temp)
	fmt.Printf("variable type: %T \n", temp)

	var launchAngle float64 = 101.7899364973246923
	fmt.Println(launchAngle)
	fmt.Printf("variable type: %T \n", launchAngle)

	var futname string
	fmt.Println(futname)
	fmt.Printf("variable type: %T \n", futname)

	var futnum int
	fmt.Println(futnum)
	fmt.Printf("variable type: %T \n", futnum)

	var idcard = 87
	fmt.Println(idcard)
	fmt.Printf("variable type: %T \n", idcard)

	schoolName := "APJM"
	fmt.Println(schoolName)
	fmt.Printf("variable type: %T \n", schoolName)

	fmt.Println(parentName)

	fmt.Println(Blood)
}
