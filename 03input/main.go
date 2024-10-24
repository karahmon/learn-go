package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for our Pizza:")

	//Comma-ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of Rating is %T \n", input)

}
