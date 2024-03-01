# Concurrency
- Go is the first major programming language that has concurrency built in
- `concurrency` is not `Parallelism`
- concurrency : 
    - means loading more `goroutines` at a time
    - if one `goroutines` block, another one is picked up and started
    - on single core CPU, we can run only concurent applications but they are not run in parallel
    - independently executing processes
    - dealing with multiple things at once
- `parallelism` :
    - means multiple `goroutines` executed at the same time
    - requires multiple CPUs
    - simultaneous executions of processes
    - requires multiple core CPUs

## Goroutines
- lightweight thread of executions
- key ingredients to achieve concurrency in go
- a function that is capable of running with concurrently with other function
- use keyword `go` and followed by a function invoation
- far smaller than `threads`
    - go : arround 2kb
    - thread : fixed size of 1-2Mb
- `OS Thread Stack` is fixed size but `goroutines stack size` shrink and grosw as needed
- schedulling a goroutine is much cheaper than schedulling a thread
- `Os threads` are scheduled by the OS kernel, but `goroutines` are scheduled by its own scheduler
- `goroutines` scheduler using technique called `m:n scheduling` because it multiplexes (or schedules) `m goroutines` on `n OS threads`
- `goroutines` have no identity

## WaitGroups
- solve the problem of `goroutines` where the app terminated before `goroutines` execute the code
- blocking or waiting until the `goroutins` within any `waitgroups` have been succesfully executed
- pattern :
    - Create a new variables of a `sync.WaitGroup` (wg)
    - Call `wg.Add(n)` where `n` is the number of goroutines to wait for
    - Execute `defer wg.Done()` in each goroutine to indicate to the WaitGroup that the goroutine has finished executing
    - Call `wg.Wait()` in main() where we want to block.

## Data Race
- executing many `goroutines` at the same time without special handling can introduce a dreaded error called `race condition` or `data race`
- most common and hardest to debug types of bugs in concurrent systems
- `race conditions` occurs because of unsynchronized access to shared memory
- the result will different based on which `goroutines` finished first

### Go Race Detector
- a new tool for finding `race conditions`
- currently available for linux, mac, and windows 64
- based on `C++` runtime library which has been used to detect many errors
- use `race` flag to go command line tool
- format : `go run -race <filename>.go`
- when race command line flag is set, the compiler instruments all memory accessed with code that records when and how memory was accessed 

### Mutexes
- one of the solutions for the `data race` issue
- abbr from mutual exclusion
- `explicit synchronization` : 
    - where variable access is protected through synchronization primitive such as mutex
    - programmer recognizes candidates for concurrent execution and the context in which they will be executed
- Step for mutexes
    - declare the mutex value from the `sync` package
    - `lock` method blocks the access to the variable until the `unlock` method is called
    - if one `goroutines` holds the `lock`, when the new `goroutines` trying to aquire the `lock` it will be blocked until the mutex is unlocked
    - method `unlock` can be called using `defer function` so it will be called when the `function` exit


### Channel
- another solutions for the `data race` issue
- provides a connection between 2 `goroutines`, allowing them to communicate
- the data sent through or receiving from a `channel` must be the same type
- the type specified when the `channel` declared
- declaration : `var <name> chan <type>`
- unbuffered channel : 
    - `<name> = make(chan <type>)`
    - the `sender` blocks on the `channel` until the `receiver` receives data from the `channel`
    - synchronous channel
- buffered channel : 
    - `<name> = make(chan <type>, <val>)`
    - the `sender` will block on the `channel` only when there is no empty slot in the `channel`
    - the `receiver` will block on the channel when it's empty.
    - first in first out
- the value of unitialized `channel` is nil
- a `channel` is like `pointer`, and has `memory address`

#### Channel Operation
- `close` : 
    - sets a flag ndicating that no more values will ever be sent on this `channel`
    - using built in function `close(<channel>)`
    - communicate with closed channel will raise `deadlock` error
- `receive`:
    - get the value from the `send` statement
    - receive the channel from closed channel will yield the values that have been sent until no more values are left
    - any receive operations thereafter yield the zero value of the channels element type
    - using symbol `<-`
- `send` : 
    - transmits a value from one `goroutine` through the `channel` to another `goroutine`
    - using symbol `<-`

#### Select Statement
- lets a `goroutines` wait on multiple communication
- blocks until one of its cases run then executes the case
- choose a random one if multiple are ready
- only used with `channel`
