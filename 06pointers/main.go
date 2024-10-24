package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	var ptr *int // This is how we declare pointer, where hard value is to be taken
	fmt.Println(ptr)

	myNumber := 23
	var ref = &myNumber // This is how we get memory allocation reference
	fmt.Println(ref)
}
