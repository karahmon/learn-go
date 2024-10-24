package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in Golang")
	// no inheritance in golang; No Super or Parent

	hitesh := User{"Monil", "monil@kubers.club", true, 20}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh Details are : %+v \n", hitesh) // to use get complete details like structure naming convention

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
