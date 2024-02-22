# Constant
- represent fixed (unchanged) values
- used to avoid possible errors (variables that change when they should not)
- can replace a value only in one place instead of in many places
- all basic literals are in fact unnamed constants
- can store numbers, strings, or booleans
- belongs to compile-time so error can be detected earlier
- can not be initiated at runtime
- can not be initiated using variable
- can be initiated using built in function

## Declaration

### Single Constanst
- `constant <name> <type> = <value>`
- `constant <name> = <value>`

### Multiple Constanst 
- `constant <name1>, <name2>, <type> = <value1>, <value2>`
- `constant <name1>, <name2> = <value1>, <value2>`
- `constant (<name1> = <value> <...>)`

## Iota
- represents succesive untyped integer constant
- the value is the index of the respective `ConstSpec` in that constant declaration
- can be used to construct a set of related constants
- use `_` to skip the iota index