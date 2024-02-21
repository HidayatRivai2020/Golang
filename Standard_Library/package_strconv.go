package main

import (
	"fmt"
	"strconv"
)

func main() {
	var result, err = strconv.ParseFloat("3.142", 64)
	if err == nil {
		fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
		fmt.Println("Cannot convert to float64!")
	}
}
