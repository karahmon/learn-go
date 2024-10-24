package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and Case in Golang")

	diceNumeber := rand.Intn(6) + 1
	fmt.Println("Value of Dice is : ", diceNumeber)

	switch diceNumeber {
	case 1:
		fmt.Println("Dice Value is One and you can open")
		// fallthrough - if mentioned will go to next case and execute if valid
	case 2:
		fmt.Println("Dice Value is two and you can open")
	case 3:
		fmt.Println("Dice Value is three and you can open")
	case 4:
		fmt.Println("Dice Value is four and you can open")
	case 5:
		fmt.Println("Dice Value is five and you can open")
	case 6:
		fmt.Println("Dice Value is six and you can open")
	default:
		fmt.Println("Invalid Input")

	}

}
