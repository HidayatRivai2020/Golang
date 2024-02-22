# Standard Library
- `fmt` : implements formatted I/O with functions analogous to C's printf and scanf.
- `os` : communicates with the operating systems and allows you to get OS functionalities
- `strconv` : convert string into another data-type
- `string` : implements simple functions to manipulate UTF-8 encoded strings
- `utf8` : support text encoded in UTF-8


## fmt

### Method
- Some generic parameter
    - `<str>` parameter : String parameter that will be used in method as input
    - `format` : the format that will be used for the text
- `Print(<str>)` : print text 
- `Printf(<format>, <str>)` : print text with formatting 
- `Println(<str>)` :  print text and add a new line
- `Sprintf(<format>, <str>)` : return string with the format specifier

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
- Some Generic Parameter
    - `<str>` parameter : String parameter that will be used in method as input
- `ParseFloat(<str>, <bitSize>)` : convert string into float data-type 
    - `bitSize` : the size of the data-type 

## String

## Method
- Some Generic Parameter
    - `<str>` parameter : String parameter that will be used in method as input
- `Contains(<str>, <substr>)` :  check if string contains substring
    - `<substr>` : the substring that will be checked in string
- `ContainsAny(<str>, <Unicode>)` :  check if string contains any unicode code point
    - `<Unicode>` : the unicode code point that will be checked in string
- `ContainsRune(<str>, <rune>)` :  check if string contains any unicode code point
- `Count(<str>, <substr>)` :  count the number of instance of substring in a string
    - `<substr>` : the substring that will be counted in string
    - if `<substr>` is `""` returns 1 + the number of runes in the string 
- `EqualFold(<str1>, <str2>)` : compare 2 strings without checking the case
- `Fields(<str>)` : split string by white-space and new-lines
- `Join(<slice_str>, <separator>)` : concatenates the elements of a slice of strings to create a single string.
    - `<slice_str>` : the slice of string that will be concatenates into a single string
    - `<separator>` : the separator that will be added between the slice element
- `Repeat(<str>, <count>)` : create a new string consisting of n copies
    - `<count>` : total copies of the string in repeat
- `Replace(<str>, <old>, <new>, <count>)` : replace some the substring of the old string with the new one
    - `<old>` :  old substring
    - `<new>` :  new substring
    - `<count>` : total substring that need to be replaced
    - if `<count>` is `-1` replaces all occurrences of old by new
- `ReplaceAll(<str>, <old>, <new>)` : replace all the substring of the old string with the new one
    - `<old>` :  old substring
    - `<new>` :  new substring
- `Split(<str>, <separator>)` : slices a string and returns a slice of the substrings between those separators.
    - `<separator>` : the substring that will be used to separate string
    - if `<separator>` is "" splits the string after each UTF-8 rune literal.
- `ToLower(<str>)` : change string to lowercase
- `ToUpper(<str>)` : change string to uppercase
- `Trim(<str>)` : remove all leading and trailing characters, use Trim()
- `TrimSpace(<str>)` : removes leading and trailing whitespaces and tabs


## UTF-8

### Method
- Some Generic Parameter
    - `<str>` parameter : String parameter that will be used in method as input
- `DecodeRuneInString(<str>)` : decode the rune in the string
- `RuneCountInString(<str>)` : returning the number of runes in the string