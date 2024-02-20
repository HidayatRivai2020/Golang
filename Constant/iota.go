package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c4 = iota
		c5
		c6
	)

	fmt.Println(c4, c5, c6)

	const (
		c7 = iota * 2
		c8
		c9
	)

	fmt.Println(c7, c8, c9)

	const (
		c10 = (iota * 2) + 5
		_
		c11
		c12
	)
	fmt.Println(c10, c11, c12)

}
