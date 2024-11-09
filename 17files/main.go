package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Golang")
	content := " Needs to go in file"
	file, err := os.Create("./a.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Lenght is : ", length)
	defer file.Close()
	readfile("./a.txt")
}

func readfile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
