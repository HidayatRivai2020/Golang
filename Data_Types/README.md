# Data Types
- determines a set of value together with operations and method specific to thosse values
- `predeclared types` : introduced types with type declarations
- `composites types` : element types

## Predeclared Types (Built in Types)
- numeric types
- Bool Type
- String Type

### Numeric types
- number data types
- can use `_` to write large numbers for a better readability
- `_` will be ignored
- integer
    - int8 : (-128 to 127)
    - int16 : (-32768 to 32767)
    - int32 : (-2147483648 to 2147483647)
    - int64 : (-9223372036854775808 to 9223372036854775807)
    - int : alias for uint32 or uint64 based on platform
    - rune : alias for uint32
- unsigned integer
    - uint8 : (0 to 255)
    - uint16 : (0 to 65535)
    - uint32 : (0 to 4294967295)
    - uint64 : (0 to 18446744073709551615)
    - uint : alias for uint32 or uint64 based on platform
    - byte : alias for uint8
- float : zero before the decimal point separator can be omitted
    - float32 : IEEE-754 32-bit floating-point numbers
    - float64 : IEEE-754 64-bit floating-point numbers
    - complex64 : complex numbers with float32 real and imaginary parts
    - complex128 : complex numbers with float64 real and imaginary parts

### Bool Types
- predefined constant true and false

### String Types
- Unicode chars written enclosed by double quotes
- sequence of bytes (possibly empty)

## Composite Types
- Array
- Slice
- Map
- Struct
- Pointer
- Function
- Type
- Channel

### Array
- numbered sequence of elements of single type
- has fixed length

### Slice
- numbered sequence of elements of single type (element types)
- has dynamic length

### Map
- unordered group of elements of one type
- indexed by a set of unique keys of another type
- similar to dictionary in Python

### Struct
- sequence of named element called `fields`
- each of element has a name and type
- can be compared to `class` concept in `OOP`
- `type <name> struct { <var_name> <var type> ... }`

### Pointer
- variable that stores the memory addres of another variable
- default value : nil

### Function
- a type of a declared function

### Interface

### Channel
- provides a mechanism for concurrently execution functions to communicate by sending and receiving a specified element type

## Converting Types
- changes the type of an expression to the type specified by the conversion

### Numeric
- converting between numeric data types
- `<new_data_type>(<variable>)`

### Numbers into Strings
- converting numbers into string will return the character that has that int as ASCII code
- converting into ASCII character only works for uint
- the format to convert numbers into string : `string(<number>)`
- use `Sprintf` to convert numbers value into string value
- the format to convert numbers value into string : `Sprintf(<value>)`

### Convert String
- use `strconv` package to convert string into another type
- the format : `strconv.<ParseType>(<value>, size)`
