package main

import "fmt"

func main() {

	//int type
	var i1 int8 = -128     //min value
	fmt.Printf("%T\n", i1) // => int8

	var i2 uint16 = 65535  //max value
	fmt.Printf("%T\n", i2) // => int16

	var i3 int64 = -324_567_345  // underscores are used to write large numbers for a better readability
	fmt.Printf("%T\n", i3)       // => int64
	fmt.Printf("i3 is %d\n", i3) // => i3 is -324567345 (underscores are ignored)

	//float64 type
	var f1, f2, f3 float64 = 1.1, -.2, 5. // trailing and leading zeros can be ignored
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	//rune type
	var r rune = 'f'
	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
	fmt.Printf("%c\n", r) // => f

	//bool type
	var b bool = true
	fmt.Printf("%T\n", b) // => bool

	//string type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s) // => string

}
