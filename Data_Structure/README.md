# Data Structure
- arrays
- slices
- maps

## Arrays
- indexable type that stores a collection of elements of same type
- has fixed length
- every element must be of same type
- elements of the array stored in contiguous memory locations
- length and the elements type determines the type of array
- the length belongs to array type and determined at compile time

### [Declaring Arrays](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/arrays.go)
- single dimensional array : `[<length>]<type>{<value1>, <value2>, <...>}`
    - `length` : the length of array, the length is `...` if no specific value
    - `type` : the type of each element
    - `value` : the value of each element, the `,` is mandatory if the length is `...`
- get the value from single array : `<var>[index]` 
- multi dimensional array : `[<length1>][<length2>]<type>{<value1>, <value2>, <...>}`
    - `length` : the length of array
    - `type` : the type of each element
    - `value` : the value of each element
- get the value from multi array : `<var>[index1][index2]`
- array with the key : `[<length>]<type>{<key>: <value>, <...>}`
    - `length` : the length of array
    - `type` : the type of each element
    - `key` : the index of element, default is last index + 1
    - `value` : the value of each element

### [Slicing Array](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/array_key.go)
- array can be sliced
- keep the old `variable` and return a new `slice`
- format : `<var>[<start>:<stop>]`
    - `var` : `array` variable
    - `<start>` : slicing `array` starting from this index, if `<start>` is missing the index will be `0`
    -  `<stop>` : slicing `array` stopping before this index, if `<stop>` is missing the index will be `len(<var>)`

## Slices
- similar to array but has a dynamic length
- the length of a slice is not part of its type and belongs to runtime
- an un-initialize slices is equal to nil

### [Declaring Slices](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/slices.go)
- use keyword `make` to initialize slices with specific length : `make([]<type>, length)`
    - `type` : the type of each element
    - `length` : the length of slices when initialize
- single dimensional slices : `[]<type>{<value1>, <value2>, <...>}`
    - `type` : the type of each element
    - `value` : the value of each element
- get the value from single slices : `<var>[index]` 
- multi dimensional array : `[][]<type>{<value1>, <value2>, <...>}`
    - `type` : the type of each element
    - `value` : the value of each element
- get the value from multi array : `<var>[index1][index2]`
- array with the key : `[]<type>{<key>: <value>, <...>}`
    - `type` : the type of each element
    - `key` : the index of element, default is last index + 1
    - `value` : the value of each element

### [Comparing Slices](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/slice_comparing.go)
- slices can only be compared to nil
- to compare the slice, compare the length first then use loop and compare for each element

### [Appending to slices](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/appending_concept.go)
- `append` : method to append a new value to a slice
    - `append(<var>, <values>, <...>)`
        - `var` : slice variable
        - `<value>` : new value that will be added into a slice
        - can append multiple values at the same time
    - `append(<var1>, <var2>...)`
        - `var1` : slices variable
        - `var2` : another slices that will be added into a slice

### [Copying Slices](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/appending_concept.go)
- `copy` : copy the value from another slice into a slice with the length of the `destination`
    - `copy(<dst>, <src>)`
        - `dst` : destination variable
        - `src` : source variable

### [Slicing Slices](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/array_key.go)
- slice can be sliced
- keep the old `variable` and return a new `slice`
- format : `<var>[<start>:<stop>]`
    - `var` : `slice` variable
    - `<start>` : slicing `slice` starting from this index, if `<start>` is missing the index will be `0`
    -  `<stop>` : slicing `slice` stopping before this index, if `<stop>` is missing the index will be `len(<var>)`

### [Slicing String](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/slicing_string.go)
- efficient because it reuses the same backing arr
- returns bytes not runes
- need to convert the slice first

## [Slices Backing (Underlying) Array Concept](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/backing_array.go)
- when creating a slice, behind the scenes Go creates a hidden array called `backing array`
- the `backing array` stores the elements, not the slice
- Go implements a slice as a data structure called `slice header`
- nil slice does not have a `backing array`
- `slices expression` does not create a `backing array`
- `slice` from `slice expression` refers to the original `backing array`
- changing `slices` obtained from `slice expression` will change the `backing array` and all other sliced created from the same `backing array`
- use append to create a brand new different `slices`  
- the memory size used in `slice` is smaller than `array`

### Slice header
- runtime representation of a slice
- contains 3 fields : 
    - `pointer` : the address of the backing array
    - `len()` : the length of the slice
    - `cap()` : the capacity of the slice

### [Appending Slice Concept](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/appending_concept.go)
- new `backing array` is created if the `capacity` is full
- the `capacity` is increased by 2 times to avoid creating a new `backing array` every time
- accessing outside `length` will return an error
- new `slice` can be created outside the `length` but `capacity`

## Maps
- collection type that stores `key:value` pairs
- `add`, `get`, and `delete` operations take constant expected time
- `keys` and the `values` are statically typed and must have same type
- `keys` must be unique
- any comparable type can be used as `key`
- `float` is not recommended used as a `key` even if it possible
- `maps` can not be compared to another `maps`
- the structure is unordered

### [Maps format](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/maps.go)
- Declaring : 
    - `var <name> map[<key_type>]<value_type>`
    - if the `key` doesn't exist or the `map` is not initialized, returns the `0` for the `<value type>`
- Using make : `make(map[<key_type>]<value_type>)`
- Accessing : 
    - `<name>[<key_type>]` 
    - return `<value>, <isFound>`
- Changing Value : `<name>[<key_type>] = <new_value>`
- Delete Element : `Delete(<name>, <key>)`

### [Comparing Maps](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/maps_compare.go)
- `maps` can not be compared to another `maps`
- use `loop` and compare each element of the `map`
- use `Sprintf` to get the value as a string then compare

### [Map Header](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Structure/maps_header.go)
- Go creates a `pointer` to a `map header` in memory when declaring a `map`
- the `map` contains only the memory `address` of `map header`
- `key value` stored in memory at the `address` referenced by the `map header`
- when map copied into a new map, the internal data structure is not copied, just referenced
- to clone a `map` into different `address`, create a new `map` and copy each element using `loop`