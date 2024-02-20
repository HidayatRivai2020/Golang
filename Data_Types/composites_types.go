package main

import "fmt"

func main() {
	//array type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers) // =>  [4]int

	//slice type
	var cities = []string{"London", "Bucharest", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // => []string

	//map type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
	}
	fmt.Printf("%T\n", balances) // => map[string]float64

	//struct type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you) // => main.Person

	//pointer type
	var x int = 2
	ptr := &x                                                 // pointer to int
	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr) // => ptr is of type *int with value 0xc000016168

	//function type
	fmt.Printf("%T\n", f) // => func()

}

func f() {
}
