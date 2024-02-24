# File
- the common way to work with file is using package `os`
- need to use valid path according to the OS
- more package that can be used to work with file
    - `io`
    - `bufio`

## Operations on Files
- Common Parameter
    - `<file>` : file name
- `Close()` : close the current file from `os.file`
- `Create(<file>)` : create a new file
    - create file if does not exists
    - truncate file if it exists
    - return `descriptor` which is pointer to `os.File` and error value
- `IsNotExists(<file>)` : check if the file not exists
- `Open(<file>)` : open the selected file in read only mode
- `OpenFile(<file>, <flag>, <perm>)` : open the selected with more option
    - `flag` : the mode of how file should be opened
    - `perm` ; the permission
    - opening attributes can be used individually or combined using `OR` between them
- `Rename(<oldpath>, <newpath>)` : rename or move the file
    - `<oldpath>` : the path of file that need to be changed
    - `<newpath>` : the new path of the file
- `Remove(<file>)` : Remove the file
- `Stat()` : getting file info
- `Truncate(<file>, <size>)` : changes the size of file
    - `<size>` : the new size for the file

### File Info
- `Name()` : base name of the file
- `Size()` : length in bytes for regular files; system-dependent for others
- `Mode()` : file mode bits
- `ModTime()` : modification time
- `IsDir()` : abbreviation for Mode().IsDir()
- `Sys()` : underlying data source (can return nil)

### File Flag
- `os.O_RDONLY` : Read only
- `os.O_WRONLY` : Write only
- `os.O_RDWR` : Read and write
- `os.O_APPEND` : Append to end of file
- `os.O_CREATE` : Create is none exist
- `os.O_TRUNC` : Truncate file when opening

## Writing into File
- Writing file Using `os.File.Write`
    - Open the file first. store it in a variable
    - Close the file using `defer file.Close()`
    - create the byte slices
    - use `<file>.Write(byteSlices)` to write
        - `<byteSlices>` : the text that will be added into file
- Writing file Using `os.WriteFile`
    - create the byte slices
    - use `os.Write(<file>, <byteSlices>, <perm>)` to write
        - `<byteSlices>` : the text that will be added into file
        - `<perm>` : the permission to write the file
- Writing file Using `os.WriteFile`
    - Open the file first. store it in a variable
    - Close the file using `defer file.Close()`
    - Create the `bufferedwriter` using `bufio.NewWriter(<file>)`
    - create the `byte slices`
    - Write the `byte slices` into buffered writter using `<bufferedwriter>.write(<bs>)`
        - `bufferedwriter` : the writer that has been created before
        - `bs` : the byte slices that has been created before
    - To Write the string into buffered writter using `<bufferedwriter>.writeString(<Text>)`
        - `bufferedwriter` : the writer that has been created before
        - `Text` : the new text that want to be added into a file
    - use `<bufferedWriter>.Flush()` to add the text into file
    - use `<bufferedWriter>.Reset()` to reset the buffered writer

## Reading from file
- Common Parameter :
    - `<file>` : file name
- Reading file Using a `io.ReadFull` : reading with the limit of `byteSlices`
    - Open the file first. store it in a variable
    - Close the file using `defer <file>.Close()`
    - create the byte slices
    - use `io.Readfull(<file>, <byteSlices>)` to read
        - `<byteSlices>` : the variable to store the data from the file
- Reading file Using `io.ReadAll` : read the entire file
    - Open the file first. store it in a variable
    - use `os.ReadAll(<file>)` to read every byte from the file
- Reading file Using `io.ReadFile` : read the entire file into memory
    - Open the file first. store it in a variable
    - use `os.ReadFile(<file>)` to read the entire file
- Reading file Using `bufio.Scanner` : read data line by line from a file using delimiter
    - Open the file first. store it in a variable
    - Close the file using `defer file.Close()`
    - Create the `<scanner>` using `bufio.NewScanner(<file>)`
    - Use `scanner.split(<type>)` if want to change how scanner split the word
        - `<type>` : split type for the scanner
    - Use `<scanner>.Scan()` to scan the next token
    - use `<scanner>.Text()` to get the value of the the current line
    - use `<scanner>.Byte()` to get the value of the the current line
    - use `<scanner>.Err()` to get the Error
    - use loop on `<scanner>.Scan()` to print every line

### Scanner Split Types
- `bufio.scanLines` : (Default) split by lines
- `bufio.scanRunes` : split by runes
- `bufio.scanWords` : split by words

## Reading from User Input
- create new scanner using `bufio.NewScanner(os.Stdin)` and store it in variable
    - `os.Stdin` : Operation System Standard Input, reading from command line
- use `<scanner>.Scan()` to read the Text from command line from theuntil the `new line`
- use `<scanner>.Bytes()` to read the Byetes from command line from theuntil the `new line`
- to read multiple lines and exit at specific value, use `for` loop and `break` using `if`