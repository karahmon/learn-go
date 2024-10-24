package main

import "fmt"

func main() {
	fmt.Println("Learning If-Else")

	loginCount := 1000

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 && loginCount < 100 {
		result = "Pro User"
	} else {
		result = "Something else"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Numer is Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
