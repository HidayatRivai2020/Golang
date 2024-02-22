package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 35}

	// both maps reference the same data structure in memory
	neighbors := friends

	// modifying friends AND neighbors
	friends["Dan"] = 30

	fmt.Println(neighbors) // -> map[Dan:30 Maria:35]

	// How to clone a map?
	// 1. initialize a new map
	colleagues := make(map[string]int)

	// colleagues = friends // -> ERROR, illegal with maps!

	// 2. use a for loop to copy each element into the new map
	for k, v := range friends {
		colleagues[k] = v
	}

	// colleagues and friends are different maps in memory!
}
