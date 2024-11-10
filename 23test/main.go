package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(3.4, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cant Divide by 0")
	}
	result = x / y
	return result, nil

}
