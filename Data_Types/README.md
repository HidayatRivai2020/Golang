# Data Types
- determines a set of value together with operations and method specific to thosse values
- `predeclared types` : introduced types with type declarations
- `composites types` : element types
- `defined types` : define new data types
- `aliases` : give new name to data types

## [Predeclared Types (Built in Types)](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/predeclared_types.go)
- numeric types
- Bool Type
- String Type

### [Numeric types](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/predeclared_types.go)
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

### [Bool Types](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/predeclared_types.go)
- predefined constant true and false

### [String Types]((https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/string.go))
- Unicode chars written enclosed by double quotes
- sequence of bytes (possibly empty)
- UTF-8 by default
- Use `backslash` to escape a character
- Use `backsticks` to create a raw string
- Use addition operator or `+` sign to concatenate strings
- character in string can be accessed by index started from zer : `<var>[index]`
- String elements are immutable and can not be changed
- series of byte values (slices of bytes)

### [Byte and Rune](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/byte_rune.go)
- used to distinguish characters from `integer` value
- Go does not have a `char` data type
- enclosed in single quotes
- represented using a unique code points
- unique code points : numeric value that represents a rune literal
- The character encoding scheme ASCII which is unicode subset, comprises 128 cide points
- any slice can be encoded in a `string` value
- `rune` : code point that reresent a single unicode character

## [Composite Types](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/composites_types.go)
- Array
- Slice
- Map
- Struct
- Pointer
- Function
- Type
- Channel

## [Converting Types](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/converting_types.go)
- changes the type of expression to the type specified by the conversion

### [Numeric](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/converting_number.go)
- converting between numeric data types
- `<new_data_type>(<variable>)`

### [Numbers into Strings](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/converting_string.go)
- converting numbers into string will return the character that has that int as ASCII code
- converting into ASCII character only works for uint
- the format to convert numbers into string : `string(<number>)`
- use `Sprintf` to convert numbers value into string value
- the format to convert numbers value into string : `Sprintf(<value>)`

### [Convert String]((https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/strconv_package.go))
- use `strconv` package to convert string into another type
- the format : `strconv.<ParseType>(<value>, size)`

## [Defined Types](https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/defined_type.go)
- Methods can be attached into newly defined types
- Type Safety : one type must be converted into another to perform operations
- readability : can represent specific new type outside the built-in type
- format `type <name> <data_type>

## [Aliases]((https://github.com/HidayatRivai2020/Golang/blob/main/Data_Types/aliases.go))
- Same type with a new name
- can be used together in operations without type conversions
- alias example :
    - byte and uint8
    - run and int32
- format `type <T1> <T2>`