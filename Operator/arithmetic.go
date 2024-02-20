package main

import "fmt"

func main() {
	a, b := 10, 5.5

	fmt.Println(a + 5)   // => 15
	fmt.Println(3.1 - b) // => -2.4
	fmt.Println(a * a)   // => 100
	fmt.Println(a / a)   // => 1
	fmt.Println(11 / 5)  // => 2

	fmt.Println(a * int(b))     // => 50
	fmt.Println(float64(a) * b) // => 55

}
