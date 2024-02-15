# Variables
- name for a memory location where a value of a specific type is stored
- belongs and created at runtime
- mandatory
- variables type can not be changed after assigned
- All variables are initialized and have default value
- `_` is the blank identifier and mutes the compile-time error returned by unused variables

## Naming Convention
- Important for code readability and maintainability
- Start with a letter or underscore `_`
- Case Sensitive
- Keyword can not be used as a names
- Should use camelCase
- Use fewer letters in smaller scopes and the complete word in larger scopes
- Should use UPPERCASE for the acronym

## Type

### Go Zero Values
- numeric : 0
- bool : false
- string : ""
- pointer type : nil

## Declaration

### Single Variable
- using `var` keyword
    - `var <var_name>`
    - `var <var_name> = <var_value>`
    - `var <var_name> <var_type> = <var_value>`
- using `:=` symbols
    - can be used only to local variable
    - `<var_name> := <var_value>`

### Multiple variables
- using `var()` command
    - `var(<vars_name> <vars_type)`
- using `var` keyword
    - `var <var1_name>, <var2_name> <vars_type>`
- using `:=` symbols
    - declaring multiple variables at the same time
    - atleast 1 variable is a new variable
    - `<var1_name>, <var2_name> := <var1_value>, <var2_value>`