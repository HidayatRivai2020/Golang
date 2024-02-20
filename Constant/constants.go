package main

import "fmt"

func main() {
	const day int = 7
	const people = 3

	var workingDay = day / people

	fmt.Println(workingDay)

	const (
		a = 1
		b = 2
	)

	fmt.Println(a)
	fmt.Println(b)

	const (
		min1 = 1
		min2
		min3 = 3
		min4
	)

	fmt.Println(min1)
	fmt.Println(min2)
	fmt.Println(min3)
	fmt.Println(min4)
}
