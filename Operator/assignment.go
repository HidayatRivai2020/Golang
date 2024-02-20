package main

import "fmt"

func main() {
	a := 10

	a = 11
	fmt.Println(a)
	a += 2
	fmt.Println(a)
	a -= 3
	fmt.Println(a)
	a *= 2
	fmt.Println(a)
	a /= 3
	fmt.Println(a)
	a %= 5
	fmt.Println(a)

}
