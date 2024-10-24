package main

import (
	"fmt"
)

const LoginToken int = 123456543 //Public

func main() {
	var username string = "Monil"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username) // To find a type of variable

	var isPresent bool = true
	fmt.Println(isPresent)
	fmt.Printf("Variable is of type : %T \n", isPresent) // To find a type of variable

	var category uint8 = 23
	fmt.Println(category)
	fmt.Printf("Variable is of type : %T \n", category) // To find a type of variable

	var percentage float64 = 23.234567234
	fmt.Println(percentage)
	fmt.Printf("Variable is of type : %T \n", percentage) // To find a type of variable

	//default values and some aliases
	//Int Default Values is Zero
	var rollno int
	fmt.Println(rollno)
	fmt.Printf("Variable is of type : %T \n", rollno)
	//String Default Values is ""
	var description string
	fmt.Println(description)
	fmt.Printf("Variable is of type : %T \n", description)

	//implicit type
	var website = "https://kubers.club"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	//no var style - no need to add var and type - not allowed in global only under function
	numberOfUser := 3
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type : %T \n", numberOfUser)

	//Global Token Check
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
