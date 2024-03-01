# Pointer
- The computer memory (RAM) can be thought as a sequence of boxes or cells, placed one after another in a line
- Each cell is labeled with a unique number which increments sequentially
- the number is the address of the cell or its `memory location`
- Each cell holds a single value
- Everything the CPU does is fetching and storing values ito memory cells
- `variable` : alphanumeric nickname or label for a memory location
- `pointer` : variable that stores the memory address of another variable
    - the address of a variable or nil if it has not been initialized yet
    - have the power to mutate or change the data they are pointing to
    - Go has no pointer arithmetic unlike C

## Pointer operation
- `ampersand operator` : 
    - address of operator
    - using symboal `&` before the variable
    - return the memory address of the variable
- dereferencing operators
    - using symbol `*` before the variable
    - return the value of the memory address
- comparing pointer
    - a pointer can have the address of another pointer
    - the value can be dereferenced by adding more `*`
    - a pointer can be compared with another pointer
    - different variables have different pointer even the values are the same

## Pointer function
- use `*` before the `parameter type`
- use `*` before the `return type`
- Go is a pass by value language (no exception here)
    - the value of the variable is copied
    - function will create a new pointer
    - the variable works locally
- use pointer as parameter to change the value of variable inside the `function` 
- Go allow function to return a pointer to a local variable
- `slice` and `maps` values can be changed inside a `function` without using pointer as parameter
- `slices` and `maps` are not meant to be used with pointers.
- `array` and `struct` values can be changed inside a function with pointer as parameter only.