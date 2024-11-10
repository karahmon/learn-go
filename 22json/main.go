package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	myJson := `
	[
		{
			"first_name" : "Monil",
			"last_name" : "Karania"

		},
		{
			"first_name" : "Kirat",
			"last_name" : "Singh"

		}
	]`

	// write json to a struct
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println("Error Unmarshalling json", err)
	}
	fmt.Println(unmarshalled)

	//write a struct to json

	var mySlice []Person
	var m1 Person
	m1.FirstName = "Harkirat"
	m1.LastName = "Singh"

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Har"
	m2.LastName = "Singh"

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		fmt.Println("Error Marshelling Struct", err)
	}
	fmt.Println(string(newJson))
}
