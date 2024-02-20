# Operator
- symbol of the promgramming language which is able to operate on operand values
- operator category : 
    - Arithmetic
    - Bitwise
    - Assignment
    - Increment & Decrement
    - Comparison
    - Logical
    - Operator for pointers an channels

## Operator Category

### Arithmetic
- applied to a numeric values
- used to perform common mathematical operation
- operation can not be performed into mismatched operands, so need to convert it first
- category :
    `+` : addition (can be used to concatenate string)
    `-` : subtraction
    `*` : multiplication
    `\` : division
    `%` : modulus (only working to integer numeric values)

### Bitwise
- operate on unsigned integers
    - `&` : and
    - `|` : or
    - `^` : Xor
    - `~` : Not
    - `<<` : shift left
    - `>>` : shift right

### Assignment
- assign values to variables
- category
    - `=` : simple assigment
    - `+=` : increment assigment
    - `-=` : decrement assigment
    - `*=` : multiplication assigment
    - `/=` : division assigment
    - `%=` : modulus assigment

### Increment & Decrement
- `++` : `increment`, add value to the operand by the untyped constant 1
- `--` : `decrement`, substract value to the operand by the untyped constant 1

### Comparison
- compare values between two operands
- yield a boolean value
- category :
    - `==` : `equals`,
    - `!=` : `not equals`
    - `<` : `less`
    - `<=` : `less equals`
    - `>` : `greater`
    - `>=` : `greater or equals`

### Logical
- apply to boolean values
- yield a result of the same type as the operands
- category : 
    `&&` : `and`, `true` if every values are `true`
    `||` : `or`, `true` if atleast one value is `true`
    `!` : `not`, the value is the opposite the current value

## Overflow
- the result of an arithmetic operation has more breaks than can be represented in the result type
- the hight order breaks that do not fit are silently discarded for efficiency
- Go does not check for overflow, the result will go to the cycle between minimum value or maximum value
- Go check for overflow only if the expression that generates the error is evaluated at compile time



