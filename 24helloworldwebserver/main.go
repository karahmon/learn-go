package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Homepage")
}
func About(w http.ResponseWriter, r *http.Request) {
	result := Addvalues(3, 5)
	fmt.Fprint(w, "This is about Page and 3 + 5 is ", result)
}

func Addvalues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err := Dividevalues(3, 0)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	fmt.Fprint(w, "This is Divide Page and 3/0 is ", result)
}
func Dividevalues(x, y float64) (float64, error) {
	if y <= 0 {
		return 0, fmt.Errorf("the Denominator cannot be 0")
	}
	return x / y, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting Application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
