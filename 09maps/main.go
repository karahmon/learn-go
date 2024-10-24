package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["GO"] = "golang"
	languages["PY"] = "python"

	fmt.Println(languages)
	fmt.Println("JS Stands for : ", languages["JS"])

	delete(languages, "JS")

	// Loops
	for key, value := range languages {
		fmt.Printf("For Key : %v \n ,value is: %v \n", key, value)
	}
}
