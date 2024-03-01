# Packages

## Package
- project directory containing `.go` files with the same package statement at the begining
- contains many source files each ending in `.go` extension and belonging to a single directory
- package types
    - executable packages : generate executables files which can be run. the name of an exeutable package is predefined and is called main
    - non-executable pacages : libraries or dependencies that are used by other packages and can have any name, can not be executed, only imported 

### GOPATH and Code Organization
- code need to be organized in specific way
- go programmers typically keep all their go code in a single `workspace`
- `workspace` : directory in file system whose path is stored in the environment variable called `GOPATH`
- on windows : `%USERPROFILE%\go (for example: c:\users\name\go)`
- on linux : `~/go (for example: /home/name/go)`
- the values of environtment variable can be printed by `go env` command

### Creating a Package
- when importing a package into `main.go` file, it will search for the package relatively to `src` directory from go `workspace`
- create a new folder for new packages in `src` folder
- create go file of the packages
- the first line is the packages name
    - format : `package <name>`
- all files in directory must belong to the same package
- create the `function`
- import the package to `main.go`
    - format : `import <package_path>`
- use the function
    - format : `<package>.<function>()`

### Exporting Names
- `function` name should be started with uppercase
- only `function` name start with uppercase can be exported or can be accessed for other packages
- `function` name started with lowercase is private can be used only in its own directory

### Init Function
- does not take any argument or return anything
- called automatically when package initialized
- always called before the initialization
- can have multiple `init` in the same file
- most common use case of `init` function is to initialize variables that cannot be initialized in global context

## Modules
- collection of related go packages stored in a directory tree with a `go.mod` file at its root
- contains one or more packages
- `go.mod` defines the module path, which is also the import path and its dependency requirements
- Go command has direct support to work with modules, including recording and resolving dependencies on other modules.

### Using modules
- create project
- initialize modules : `go mod init <name>`
- import the package : `import <package_url>`
- use the function
- locally module stored at `GOPATH/pkg/mod/`

### Creating Modules on Github
- create new repository on github
- create the directory on local
- create the packages
- initialize modules
- publish to github