package main

import (
	"fmt"
	"time"
)

type User struct {
	Firstname   string
	Lastname    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

//Receiver Function

func (m *User) printFirstName() string {
	return m.Firstname
}

func main() {
	user := User{
		Firstname: "Trevor",
		Lastname:  "Malik",
	}
	fmt.Println(user.Firstname)
	fmt.Println(user.Lastname)
	fmt.Println(user.BirthDate)
	fmt.Println(user.printFirstName())
}
