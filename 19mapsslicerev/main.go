package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// Creating Map

	myMap := make(map[string]User)

	me := User{
		FirstName: "Monil",
		LastName:  "Karania",
	}
	myMap["One"] = me

	fmt.Println(myMap["One"].FirstName)

	//Slice
	var mySlice []string

	mySlice = append(mySlice, "Monil")
	fmt.Println(mySlice)

}
