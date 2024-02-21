package main

import (
	"fmt"
)

func main() {
	//the following piece of code creates a loop like a for statement does
	i := 0
loop: // label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//  goto todo //ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// todo:
	//  fmt.Println("something here")
}
