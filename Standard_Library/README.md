# Standard Library
- `fmt` : implements formatted I/O with functions analogous to C's printf and scanf.
- `os` : communicates with the operating systems and allows you to get OS functionalities
- `strconv` : convert string into another data-type


## fmt

### Method
- `Print(<text>)` : print text 
    - `text` : the text that will be printed
- `Printf(<format>, <text>)` : print text with formatting 
    - `text` : the text that will be printed
    - `format` : the format that will be used for the text
- `Println(<text>)` :  print text and add a new line
    - `text` : the text that will be printed
- `Sprintf(<format>, <text>)` : return string with the format specifier
    - `text` : the text that will be printed

### Text Format :
- `%b` : biner
- `%c` : char
- `%08b` : 8bit
- `%d` : integer
- `%f` : float
- `%3f` : float with 3 decimal after comma
- `%p` : pointer
- `%q` : quoted string
- `%s` : string
- `%T` : variable type
- `%t` : boolean
- `%v` : any value
- `%#v` : representation of the value
- `%x` : hexadecimal
- `\n` : new line 
- `\"` : double quotes


## Os

### Properties
- `Args` : hold the command-line arguments in string, starting with the program name


## Strconv

### Method
- `ParseFloat(<text>, <bitSize>)` : convert string into float data-type 
    - `text` : the text that will be converted
    - `bitSize` : the size of the data-type 