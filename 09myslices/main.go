package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("hello jerrii")
	var bucketList = []string{"eating", "camping", "wandering"}
	fmt.Printf("type of bucketlist is %T\n", bucketList)
	bucketList = append(bucketList, "travelling to chennai")
	fmt.Println(bucketList)

	//bucketList = append(bucketList, "ji")
	fmt.Println(bucketList)

	highScores := make([]int, 4)

	highScores[0] = 23
	highScores[1] = 45
	highScores[2] = 76
	highScores[3] = 44

	fmt.Println(highScores)

	highScores = append(highScores, 45, 212)
	sort.Ints(highScores)

	fmt.Println(highScores)

	courses := []string{"golang", "python", "rust"}
	fmt.Println("listed courses are", courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
