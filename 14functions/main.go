package main

import "fmt"

func main() { // Function cannot be written inside a function
	fmt.Println("Welcome to Functions Class")
	greeter()
	greeterTwo()
	result := adder(3, 5)
	fmt.Println("Result is ", result)
	addValue := addallvalues(1, 2, 3, 4, 5)
	fmt.Println(addValue)
}

func addallvalues(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(x, y int) int {
	return x + y
}

func greeter() {
	fmt.Println("Namaste from Go Lang")
}

func greeterTwo() {
	fmt.Println("Another Method")
}
