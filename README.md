# Go
- An open-source programming language supported by Google
- Easy to learn and great for teams
- Good for 
  - webservices at scale
  - networking
  - massively concurrent application
  - operating systems
  - automations
  - command line tools
  - cryptography
- Built-in concurrency and a robust standard library
- Large ecosystem of partners, communities, and tools
- Static Programming Language
- Strong Programming Language
- Extension : `.go`

## How To Install Go in Microsoft Windows
- Link : [golang.org/dl/](https://golang.org/dl/)
- Select Featured downloads -> **.msi**
- Execute the installer
- Next until finish

## [Hello World](https://github.com/HidayatRivai2020/Golang/blob/main/main.go)
- `package main` : use package `main`
- `import fmt` : import package `fmt`
- `func main()` : create a function `main`
- `fmt.Print("Hello World")` : use `fmt` to print **"Hello World"**

## Go Terminal Command
- `go build` : Compiles the directory
- `go build <file_name>` : Compiles a file in the directory
- `go build -o <app_name>` : Compiles the directory with specific name
- `go env` : showing Go environment
- `go fmt` : change the code format into recommended format
- `go install` : build the application relative to `GOPATH/src`
- `go run <file_name>` : Compiles and runs the application, does not produce an executable
- `go work init` : initiate the go workspace in a folder
- `go work use <folder_name>` : add a folder into multiple module workspaces
- `GOARCH=<arch_name>` : change the architecture target of build
- `GOPATH=<path_name>` : change the relative path of Go
- `GOOS=<os_name>` : change the OS target of build


## List Of Contents
- [Playground](https://github.com/HidayatRivai2020/Golang/tree/main/Playground)
    - Interface
- [Variables](https://github.com/HidayatRivai2020/Golang/tree/main/Variables)
    - Naming Convention
    - Type
        - Go Zero Values
    - Declaration
        - Single Variable
        - Multiple Variables
    - Scope
- [Comment](https://github.com/HidayatRivai2020/Golang/tree/main/Comment)
- Constant
    - Declaration
        - Single Constant
        - Multiple Constant
    - Iota
- [Data Types](https://github.com/HidayatRivai2020/Golang/tree/main/Data_Types)
    - Predeclared Types
        - Numeric
        - Bool
        - String 
        - Byte and Rune 
    - Composite Types
        - Array
        - Slice
        - Map
        - Struct
        - Pointer
        - Function
        - Type
        - Channel
    - Converting Types
        - Numeric
        - Numbers and Strings
    - Defined Types
    - Aliases
- Operator
    - Operator Category
        - Arithmetic
        - Bitwise
        - Assignment
        - Increment & Decrement
        - Comparison
        - Logical
    Overflow
- Flow Control
    - Decision Making
        - If Else If Else
        - Simple If Statement
        - Switch
    - Loop
        - For 
            - Continue
            - Break
            - Goto
            - Label
- Data Structure
    - Array
        - Declaring Arrays
        - Slicing Arrays
    - Slices
        - Declaring Slices
        - Comparing Slices
        - Appending to Slices
        - Copying Slices
        - Slicing Slices
        - Slicing String
    - Slices Backing (Underlying) Array Concept
        - Slices Header
        - Appending Slice Concept
    - Maps
        - Maps Format
        - Comparing Maps
        - Map Header
- File
    - Operations on File    
        - File Info
        - File Flag
    - Writing Into File
    - Reading From File
        - Scanner Split Type
    - Reading From Input
- Structs
    - Struct
        - Creating Struct
        - Struct Field Value
    - Anonymous Struct
    - Anonymous Field
    - Embedded Struct
- Functions
    - Function
        - Writing Function
        - Calling Method
    - Variadic function
    - Defer function
    - Anonymous Function
    - Methods
        - Receiver Function
        - Pointer Function
- Pointer
    - Pointer Operation
    - Pointer Function
- Interface
    - Interface Nil Type
    - Interface Dynamic Type
        - Polymorphism
        - Type Assertion
        - Type Switch
    - Embedded Interface
    - Empty Interface
- Concurrency
    - Goroutines
    - WaitGroups
    - Data Race
        - Go Race Detector
        - Mutexes
        - Channel
            - Channel Operation
            - Select Statement
- Packages
    - Package
        - Gopath and Code Organization
        - Creating a Package
        - Exporting Names
        - Init Function
    - Modules
        - Using Modules
        - Creating Modules on Github

## Standard Library
- fmt
    - Method
    - Text Format
- log
    - Method
- os
    - Properties
- strconv
    - Method
- string
    - Method
- UTF-8
    - Method