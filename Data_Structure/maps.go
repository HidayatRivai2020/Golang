package main

import "fmt"

func main() {
	var employees map[string]string

	// the zero value of a map is nil
	fmt.Printf("%#v\n", employees) // -> map[string]string(nil).

	fmt.Printf("No. of elements: %d\n", len(employees)) // => No. of elements: 0

	// getting the corresponding value of a key
	fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"]) // => The value for key "Dan" is ""

	// declaring and initializing a map using a map literal
	people := map[string]float64{} // empty map

	// inserting key:value pairs in a map
	people["John"] = 30.5
	people["Marry"] = 22

	fmt.Printf("%v\n", people) // => map[John:30.5 Marry:22]

	// declaring and initializing a map using the make() function:
	map1 := make(map[int]int)
	fmt.Printf("map1: %#v\n", map1) // -> map1: map[int]int{}
	map1[4] = 8

	// declaring and initializing a map using a map literal
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
		//50: "ABC"  //illegal, all keys have the same type which is string and all values have the same type which is float64
		"CHF": 600, //this last comma (,) is mandatory when declaring the map on multiple lines
	}
	fmt.Println(balances) // => map[CHF:600 EUR:555.11 USD:233.11]

	//If we declare and initialize a map on a single line, it's not mandatory to use a comma after the last element
	m := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m

	// initializing a map with duplicate keys
	// n := map[int]int{1: 3, 4: 5, 6: 8, 1: 4} // error -> duplicate key 1 in map literal

	// if the key exists it updates its value and if the key doesn't exist it adds the key: value pair
	balances["USD"] = 500.5
	balances["GBP"] = 800.8
	fmt.Println(balances) // => map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	// "comma ok" idiom is used to distinguish between a missing key:value pair and an existing key with value zero
	v, ok := balances["RON"]

	//v is the key's corresponding value
	// ok is bool type value which is true if the key exists and false otherwise

	if ok {
		fmt.Println("The RON Balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	// iterating over a map
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	//starting with go 1.12 fmt.Printf() function prints out the map sorted by key.
	fmt.Printf("balances: %v\n", balances) // => balances: map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	// deleting a key:value pair from the map
	delete(balances, "USD")
}
