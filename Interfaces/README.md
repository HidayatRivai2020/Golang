# Interfaces
- collection of method signatures that an object which is most of the time a named type can implement
- defines the behavior of an object
- can achieve polymorphism
- not a generic type like in other programming language
- can make code clearer, shorter, and more readable
- can provide nice API between packages or client and servers
- format :
    - `<name>` : interface name
    - `<function>` : method name of interface
    - `<type>` : return type
```
type <name> interface {
    <function>() <type>
    ...
}
```
- No need a keyword to implement interface like other programming language
- if a type implements all the method in the interface, then that type is said to implement that interface

## Interface Nil Type
- creating an iterface type is creating an abstract type value
- the zero value of interface is nil
- the nil interface holds neither value nor concrete type
- calling a method on a nil iterface will runtime error

## Interface Dynamic type
- interface value can be thought as a pair of a concrete value and a dynamic type
- interface value holds a value of a specific underlying concrete type
- calling a method on an interface value or a function that takes in an interface value executes the method of the same name on its underlying type
- interface value hides its dynamic value

### Polymorphism : 
- interface variable can take poly or many dynamic forms at runtime
- dynamic type of an interface value may vary during executions

### Type Assertion
- A type assertion provides access to an interface's concrete value. 
- interface's dynamic values can not be access directly
- the methods can be access directly only for the methods that are defined inside the interface
- use type assertion to extract and return the dynamic value of the interface value
- format : `<var_name>.(<interface>)`
- return 2 values : underlying value and boolean value that report assertion is succeeded or not
- it is a good practice to check the assertion is succeed or not because sometime type assertion can fail

### Type Switch
- construct that permits serveral type assertion in series and runs the first case with maching type
- has the same syntax with type assertion, but the specific type is replaced with keyword `type`
- format : `<var_name>.(type)`

## Embedded Interface
- create a new interface by merging two or more interface
- Go is not supporting `inheritance`
- Interface can not `implement` or `extend` other interface in go
- Interface can include other interfaces defined in the current or in another package and import it
- when interface included in another interface, all the methods from the embedded interface added into that interface
- circular embedding is not allowed and error at compile time

## Empty Interface
- Interface with no method
- Empty Interface is a Core Concept in Go
- Any Go type satisfies the empty interface and that means it can represent any value.
- An empty interface may hold values of any type.
- empty interface can not be used directly in operations

