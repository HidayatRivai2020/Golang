package main

import "fmt"

func main() {

	// the & (ampersand) operator
	name := "Andrei"
	fmt.Println(&name) // -> 0xc0000101e0

	//** DECLARING AND INITIALIZING POINTERS **//

	var x int = 2
	// the expression &x means the address of x and creates a pointer to an integer variable,
	// ptr is of type *int, which is pronounced "pointer to int".
	ptr := &x
	fmt.Printf("ptr is of type %T with value %v and address %p\n", ptr, ptr, &ptr) // -> p is of type *int with value 0xc000014140.

	// declaring a pointer without initializing it
	// its zero value is nil
	var ptr1 *float64
	_ = ptr1

	// creating a pointer using new() built-in function.
	p := new(int) // it creates a pointer called p that is a pointer to an int type

	x = 100
	p = &x // initializing p

	fmt.Printf("p is of type %T with value %v\n", p, p) // => p is of type *int with value 0xc000014140
	fmt.Printf("address of x is %p\n", &x)              // => address of x is 0xc000016120

	//** THE DEREFERENCING OPERATOR **//

	// * in front of a pointer is called the dereferencing operator
	*p = 90 //equivalent to x = 90 because *p is x
	// x and *p is the same thing.

	fmt.Println(*p)                  // => 90
	fmt.Println("*p == x:", *p == x) // => *p == x: true

	fmt.Println("Value of x:", *p) // => Value of x: 90 , equivalent to fmt.Println(x)

	*p = 10        // If I write *p = 10, this is equivalent to x = 10
	*p = *p / 2    //dividing x through the pointer
	fmt.Println(x) // -> 5

	//** COMPARING OPERATOR **//

	var p2 *int
	fmt.Printf("%#v\n", p2)

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // false because they pointing different variables

	p4 := &y
	fmt.Println(p4 == p2) // true because they pointing same variables

}
