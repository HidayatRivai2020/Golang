package main

import (
	"fmt"
)

func main() {

	var x = 3   //int type
	var y = 3.2 //float type

	// x = x * y //compile error ->  mismatched types

	x = x * int(y) // converting float64 to int
	fmt.Println(x) // => 9

	y = float64(x) * y //converting int to float64
	fmt.Println(y)     // => 28.8

	x = int(float64(x) * y)
	fmt.Println(x) // => 259

	//In Go types with different names are different types.
	var a int = 5   // same size as int64 or int32 (platform specific)
	var b int64 = 2 // int and int64 are not the same type

	// a = b // error: cannot use b (type int64) as type int in assignment
	a = int(b) // converting int64 to int (explicit conversion required)

	fmt.Println(a)

}
