package main

import "fmt"

func main() {

	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	// finding 10 numbers divisible by 13
	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	fmt.Println("Finish")
}
