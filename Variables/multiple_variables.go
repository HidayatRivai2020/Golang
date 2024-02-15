package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
	)
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 1998
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var i, j int
	i, j = 5, 8
	fmt.Println(i, j)
	j, i = i, j
	fmt.Println(i, j)

}
