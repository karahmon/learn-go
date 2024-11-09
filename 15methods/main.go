package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in Golang")

	monil := User{"Monil", "monil@kubers.club", true, 20}
	fmt.Printf("Monil Details are : %+v \n", monil) // to use get complete details like structure naming convention
	monil.GetStatus()
	monil.NewMail()
	fmt.Printf("Monil Details are : %+v \n", monil)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this User is : ", u.Email)
}
