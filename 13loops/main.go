package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops!")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for d := range days {
	// 	fmt.Println(days[d])
	// }

	// for index, days := range days {
	// 	fmt.Printf("index is %v \n, and value is %v \n", index, days)
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
