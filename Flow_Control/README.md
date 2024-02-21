# Flow Control
- Decision Making
- Loop

## Decision Making
- `if else if else`
- `simple if`
- `switch`

### If Else If Else
- `if` 
    - mandatory
    - the first condition that will be checked
    - execute the block inside when the condition is true
- `else if` : 
    - optional
    - can be used multiple times after `if` and before `else`
    - the next condition that will be checked when `if` and `else if` above are false
    - execute the block inside when the condition is true
- `else` : 
    - optional
    - used in the last after `if` and `else if`
    - the block will be executed when none of `if` and `else if` are true

```
if some_condition_is_true {
    // execute this code
} else if some_other_condition_is_true {
    // execute this code
} else {
    // execute this code
}
```

### Simple If Statement
- a statement that is allowed in another statement like an if or a switch
- the variable declared are available only in the branches of that statement
- `if` 
    - mandatory
    - execute the statement first, then check the condition
    - execute the block inside when the condition is true
- `else if` : 
    - optional
    - can be used multiple times after `if` and before `else`
    - the next statement condition that will be executed when `if` and `else if` above are false
    - check the condition
    - execute the block inside when the condition is true
- `else` : 
    - optional
    - used in the last after `if` and `else if`
    - the block will be executed when none of `if` and `else if` are true

```
if statement, condition_is_true {
    // execute this code
} else if statement, some_othercondition_is_true {
    // execute this code
}
```

### Switch
- similar to `if else if else`
- Go converts `switch` statement into `if` statements behind the scenes automatically
- the purpose is to make very long `if` more readable
- can compare only the match value
- `switch` : keyword to use `switch` statement
- `<value>` : value that will be compared
- `case` : the option for the switch
- `<comparison_value>` : comparison value (mandatory)
- `<secondary_comparison_value>` : another comparison value (optional)
- `break` : break the switch
- `default` : execute this block of command where no value is the same with variable

```
switch <value> {
    case <comparison_value>:
        // execute this code
        break
    case <comparison_value2>, <secondary_comparison_value>:
        // execute this code
    default:
        // execute this code
} 
```

## Loop

### For
- used to execute a block of code repeatedly
- `for` : keyword to use for statement
- `iniialize` : variable to initialize the first value
- `condition` : the loop is repeat when the condition is true
- `iteration` : incremental and decremental that will be checked in condition
- infinite loop is the loop without iteration

```
for <initialize>; <condition>; <iteration> {
    execute this command
}

for <initialize>; <condition> {
    execute this command
    <iteration>
}

<initialize>;
for <condition> {
    execute this command
    <iteration>
}
```

#### Continue Statement
- conitinue to next iteration of the loop
- reject all the remanining statement
- used to skip the remaining code in the loop at some condition


#### Break Statement
- Stop the loop
- reject all the remanining statement
- used to terminate the inner most for or switch statement

#### Goto
- transfer control to the statement with the corresponding label within the same function
- illegal to add more commands after `goto`
- should be used only with caution because it makes the chord hard to read

#### Label
- used in `break`, `continue`, and `goto` statements
- illegal to define label that is never used
- labels are not block scoped
- labels do not conflict with identifiers that are not labels
- the scope of a label is the body of the function
- label declared abd excludes the body of any nested function
- most of the time labels are use to terminate outer enclosing loops

