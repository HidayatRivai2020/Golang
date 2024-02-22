package main

import "fmt"

func main() {
	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
