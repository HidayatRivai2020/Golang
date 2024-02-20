package main

import "fmt"

func main() {
	fmt.Println(0 < 2 && 4 > 1) // => true
	fmt.Println(1 > 5 || 4 > 5) // => false
	fmt.Println(!(1 > 2))       // => true

}
