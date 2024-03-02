# Variables
- name for a memory location where a value of a specific type is stored
- belongs and created at runtime
- mandatory
- variables type can not be changed after assigned
- All variables are initialized and have default value
- `_` is the blank identifier and mutes the compile-time error returned by unused variables

## [Naming Convention](https://github.com/HidayatRivai2020/Golang/blob/main/Variables/naming_convention.go)
- Important for code readability and maintainability
- Start with a letter or underscore `_`
- Case Sensitive
- Keyword can not be used as a names
- Should use camelCase
- Use fewer letters in smaller scopes and the complete word in larger scopes
- Should use UPPERCASE for the acronym

## Type

### [Go Zero Values](https://github.com/HidayatRivai2020/Golang/blob/main/Variables/zero_values.go)
- numeric : 0
- bool : false
- string : ""
- pointer type : nil

## Declaration

### [Single Variable](https://github.com/HidayatRivai2020/Golang/blob/main/Variables/variables.go)
- using `var` keyword
    - `var <var_name>`
    - `var <var_name> = <var_value>`
    - `var <var_name> <var_type> = <var_value>`
- using `:=` symbols
    - can be used only to local variable
    - `<var_name> := <var_value>`

### [Multiple variables](https://github.com/HidayatRivai2020/Golang/blob/main/Variables/multiple_variables.go)
- using `var()` command
    - `var(<vars_name> <vars_type)`
- using `var` keyword
    - `var <var1_name>, <var2_name> <vars_type>`
- using `:=` symbols
    - declaring multiple variables at the same time
    - atleast 1 variable is a new variable
    - `<var1_name>, <var2_name> := <var1_value>, <var2_value>`

## [Scope](https://github.com/HidayatRivai2020/Golang/blob/main/Variables/scopes.go)
- visibility
- lifetime of a variable
- interval of time during which it exists as the program executes
- a name cannot be declaraed again in the same scope
- scope category in go
    - `File` Scope : only visible in the current file
    - `Package` Scope : only visible to the current package
    - `Block (local)` Scope : only visible to the current block below the declaration
- unused variable will error in `block scope`
