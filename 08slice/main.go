package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Print Slices")

	var fruitList = []string{"Apple", "Tomato"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Mango", "Banana") // Adds these values in slices
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3]) // to cut 1st value from array and stop at position 3
	fmt.Println(fruitList)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 235
	highscore[2] = 236
	highscore[3] = 237

	fmt.Println(highscore)

	sort.Ints(highscore)
	// fmt.Println(highscore)

	//How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
