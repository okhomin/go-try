# Golang Try Catch Finally 🔥

---

# 🚧 WARNING: DEMONSTRATION PURPOSES ONLY 🚧
Attention, visitors and developers! This repository is solely intended for demonstrating the remarkable features and capabilities of the Go programming language. It is NOT suitable for utilization in any production or real-world project.  
You can find detailed information on how to handle errors in Go in the article titled "Error Handling and Go." Access the article directly using the following link:  
[🔗 Error Handling and Go Documentation](https://go.dev/blog/error-handling-and-go)

---

Golang Try Catch Finally is a simple demonstration of how to implement the try-catch-finally functionality in Go.  
The key features of go-try:
* **try** - the try block, which can execute any code that may cause an error without the need of handling errors in idiomatic Go style.
* **catch** - the catch block, which can handle errors that occur in the try block.
* **finally** - the finally block, which is executed after the try block regardless of whether an error occurred or not.  

## Getting Started
With Go module support, simply add the following import
```go
import "github.com/okhomin/go-try"
```
to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.  
Otherwise, run the following Go command to install the `go-try` package:
```sh
go get github.com/okhomin/go-try
```

## Usage
The following code snippet demonstrates how to use the `go-try` package:
```go
package main

import (
	. "github.com/okhomin/go-try"
	"fmt"
	"io"
	"os"
)

func main() {
	var file *os.File
	var data []byte

	Try(func() {
		file = Try1(os.Open("test.txt"))
		data = Try1(io.ReadAll(file))
	}).Finally(func() {
		fmt.Println("This code block will be executed no matter what")
	}).Catch(CatchMap{
		os.ErrNotExist: func(err error) {
			fmt.Printf("File does not exist: %v\n", err)
		},
		AnyError: func(err error) {
			fmt.Println("This code block will be executed if any error occurs except os.ErrNotExist")
		},
	})

	fmt.Println(string(data))
}
```
Since the module is imported with the dot prefix, the `Try`, `Catch`, and `Finally` functions can be called directly without the need to specify the package name.  

## Try
The `Try` function is used to execute the code that may cause an error.  
```go
Try(func() {
    file = Try1(os.Open("test.txt"))
    data = Try1(io.ReadAll(file))
})
```

## Try0, Try1, Try2, ..., Try10
All the functions that can return an error must be wrapped with the `Try` function.  
In the example above, the `Try1` function is used to wrap the `os.Open` function, which returns a file and an error.  
```go
file = Try1(os.Open("test.txt"))
```
After using the `Try1` function, the `file` variable will contain the file if no error occurred. Otherwise, the `file` variable will be `nil` and the error will be handled in the `Catch` block.  
```go
data = Try1(io.ReadAll(file))
```
The `Try1` function can be used to wrap the `io.ReadAll` function, which returns a byte slice and an error.

## CatchMap
The `CatchMap` type is a map of error types and functions that handle errors of the corresponding types.  
```go
CatchMap{
    os.ErrNotExist: func(err error) {
        fmt.Printf("File does not exist: %v\n", err)
    },
    AnyError: func(err error) {
        fmt.Println("This code block will be executed if any error occurs except os.ErrNotExist")
    },
}
```
`AnyError` is a special error type that is used to handle any error that occurs in the `Try` block and is not handled by any other error type.

## Finally
The `Finally` function is used to execute the code that must be executed after the `Try` block regardless of whether an error occurred or not.  
```go
Finally(func() {
    fmt.Println("This code block will be executed no matter what")
})
```
The `Finally` block can be omitted if it is not needed.

## License

Distributed under the MIT License. See [`LICENSE`](./LICENSE) for more information.







