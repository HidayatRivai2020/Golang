package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "X"}

	// fmt.Println(a == b) // error -> invalid operation: a == b (map can only be compared to nil)

	// getting a string representation of maps called a and b
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
