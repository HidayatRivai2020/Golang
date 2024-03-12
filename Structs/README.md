# Structs
- sequence of named elements called `field`
- each `field` has `name` and `type`
- like a `class` in `OOP` because Go does not have a `class object architecture`
- a schema containing a blueprint of data a structure will hold
- the blueprint is fixed at compile time
- name or the type of the field cannot be changed at runtime
- fields from a struct can not be added o removed at runtime

## [Struct](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/structs.go)

### [Creating Struct](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/structs.go)
- format : 
```
type <name> struct {
    <field1> <type1>
    <field2>, <field3> <type2>
    ...
}
```
- `<name>` : struct name
- `<field1>` : the field name
- `<type1>` : the type of the field
- field with the same type can be written in the same line
- creating a new variable from struct using the field order: `<new_variable> := <name>{<var1>, <var2>, <var3>}`
- creating a new variable from struct without order: `<new_variable> := <name>{<field3>: <var3>, <field1>: <var1>, <field2>: <var2>}`
- creating a new variable from struct only with specific field: `<new_variable> := <name>{<field2>: <var2>}`

### [Struct Field Value](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/structs.go)
- retrieving value : `<struct_variable>.<field_name>`
- updating value : `<struct_variable>.<field_name> = <new_value>`
- comparing struct : <struct_variable1> == <struct_variable2> 
- Copying struct : <struct_variable1> := <struct_variable2>
    - each struct use different memory address 

## [Anonymous Struct](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/anonymous_struct.go)
- struct with no explicitly defined struct type alias.
- format : 
```
<new_variable> := struct {
    <field1>, <field2> <type1>
    <field3> <type3>
    ...
}{
    <field1>: <value1>,
    <field2>: <value2>,
    <field3>: <value3>,
    ...
}
```
- `<new_variable>` : variable to store the struct
- `<field>` : field name
- `<type>` : data type of field
- `value` : the value of field

## [Anonymous Field](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/anonymouse_struct.go)
- field in struct where the data types becomes the fields name
- format : 
```
type <name> struct {
    <type1>
    <type2>
    ...
}
```
- `<name>` : structs name
- `<type>` : data type of field

## [Embedded Struct](https://github.com/HidayatRivai2020/Golang/blob/main/Structs/embedded_struct.go)
- struct that acts like a field inside another struct
- format : 
```
type <name> struct {
    <field> <other_struct>
    ...
}
```
- `<name>` : structs name
- `<field>` : field name
- `<other_struct>` : other struct that will be used as a fields type