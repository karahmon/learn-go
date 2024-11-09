package main

import "fmt"

func main() {
	defer fmt.Println("Hello World")
	fmt.Println("Defers in Golang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
