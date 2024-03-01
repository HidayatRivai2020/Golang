# Functions

## Function
- small piece of code that is dedicated to a perform particular task based on some input values
- must be unique within the same oackage
- can return multiple values
- Go does not support function overloading
- use keyword `func` to create function
- when a function is called, the code in function body will be executed
- the code below return will not be executed

### Writing Function
- without parameter and return
    - `<name>` : function name
```
func <name> {
    // function body here
}
```
- with parameter and return
    - `<name>` : function name
    - `<param>` : parameter name
    - `<type>` : the type of parameter
```
func <name>(param1 type1, param2 type2) {
    // function body here
}
```
- with parameter
    - `<name>` : function name
    - `<param>` : parameter name
    - `<type>` : the type of parameter
```
func <name>(param1 type1, param2 type2) {
    // function body here
}
```
- with same type parameter
    - `<name>` : function name
    - `<param>` : parameter name
    - `<type>` : the type of parameter
```
func <name>(<param1>, <param2>, <param3> <type1>, <param4>, <param5> type2) <return_type> {
    // function body here
}
```
- with single return
    - `<name>` : function name
    - `<value>` : return value
    - `<return_type>` : return type
```
func <name> <return_type1>, <return_type2> {
    // function body here
    return <value>
}
```
- with multiple return
    - `<name>` : function name
    - `<value>` : return value
    - `<return_type>` : return type
```
func <name> {
    // function body here
    return <value1> , <value2>
}
```
- with named return
    - `<name>` : function name
    - `<value>` : return value
    - `<return_var>` : return variable
    - `<return_type>` : return type
    - sometimes called as naked return
    - should be used only in short function
```
func <name>(<return_var>, <return_type>) {
    // function body here
    return
}
```

### Calling Method
- `<name>()` ; method without parameter and return
- `<name>(<param1>, param2)` ; method with parameter
- `<var> := <name>(<param>)` ; method with return, store it in variable
- `<var1>, <var2> := <name>(<param>)` ; method with multiple return, store it in multiple variable

## Variadic Function
- functions that take a variable number of arguments.
- Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
- The function may be called with zero or more arguments for that parameter.
- If the function takes parameters of different types, only the last parameter of a function can be variadic.
- format : `<name>(...<type>)` 
- non-variadic parameters comes first when mixing between variadic and non-variadic parameters.
- best time to use variadic function
    - when the parameter is unknown.
    - when it unnecessary to make a temporary slice to use it as function parameter
    - easier to read because it use one parameter instead of multiple

## Defer Function
- defers or postpones the execution of a function until the surrounding function returns.
- Last In First Out, the last defer method executed before the first defer method

## Anonymous Function
- a function which does not contain any name and is declared inline using a function literal.
- can be used closures.

## Methods
- a function that takes a receiver as argument
- methods belongs to a type while function belongs to a package
- Methods parameter : 
    - `Receiver Function`
    - `Pointer Function`

### Receiver Function
- Go does not have classes, but methods can be defined on predefined types
- A type may have a method set associated with it which enhances the type with extra behavior
- Format : 
    - `<var>` = method receiver 
    - `<new type>` : a new predefined types
    - `<function>` : function name
    - `<params>` : (Optional) parameter 
```
func (<var> <new_type>) <function>(<params>) {
    // execute this command
}
```

### Pointer Function
- The values will not be changed inside the function when using receiver function (except for `slices` and `maps`) 
- need to use pointer as parameter if the values need to be changed inside the function
- A type may have a method set associated with it which enhances the type with extra behavior
- Format : 
    - `<var>` = method receiver 
    - `<new type>` : a new predefined types
    - `<function>` : function name
    - `<params>` : (Optional) parameter
```
func (<var> *<type>) <function>(<params>) {
    // (*<var>) to access the variable
    // execute this command
}
```