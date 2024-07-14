# Go (Golang) Learning Plan

## 1. Introduction to Go

- **Overview of Go**
  - History and Philosophy
  - Features and Use Cases
    > Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.It is syntactically similar to C, but also has memory safety, garbage collection, structural typing, and CSP-style concurrency.
- **Setting Up the Environment**
  - Installing Go
  - Setting up GOPATH and GOROOT
  - Go Workspace
- **Your First Go Program**
  - Hello World Program
  - Using `go run` and `go build`
  - Go Modules and Dependencies

## 2. Basic Syntax and Concepts

- **Variables and Constants**
  - Declaring Variables
  - Variable Types and Type Inference
  - Constants and `iota`
- **Data Types**
  - Basic Types (`int`, `float`, `bool`, `string`)
  - Type Conversions
- **Operators and Expressions**
  - Arithmetic, Comparison, Logical, and Bitwise Operators
- **Control Structures**
  - `if-else` Statements
  - `switch` Statements
  - `for` Loops (traditional, range)
  - `defer`, `break`, `continue`, and `goto`
- **Functions**
  - Defining and Calling Functions
  - Parameters and Return Values
  - Named Return Values
  - Variadic Functions
  - Anonymous Functions and Closures

## 3. Composite Types

- **Arrays**
  - Declaring and Initializing Arrays
  - Accessing and Modifying Array Elements
- **Slices**
  - Creating and Manipulating Slices
  - Slice Internals and Capacity
  - Append, Copy, and Slice Expressions
- **Maps**
  - Creating and Accessing Maps
  - Adding, Updating, and Deleting Map Elements
  - Iterating over Maps
- **Structs**
  - Defining and Using Structs
  - Nested Structs
  - Struct Methods and Receivers

## 4. Pointers

- **Understanding Pointers**
  - Declaring and Using Pointers
  - Pointer Dereferencing
- **Pointer vs. Value Types**
- **Pointers to Structs and Functions**
  - Passing Pointers to Functions

## 5. Concurrency in Go

- **Goroutines**
  - Creating and Running Goroutines
  - Goroutine Lifecycle
- **Channels**
  - Creating and Using Channels
  - Buffered vs. Unbuffered Channels
  - Sending and Receiving Data
  - Channel Directions
- **Select Statement**
  - Using `select` for Multiplexing
- **Synchronization (sync package)**
  - WaitGroups
  - Mutexes (`sync.Mutex`)
  - RWMutex (`sync.RWMutex`)
  - Once (`sync.Once`)

## 6. Error Handling

- **Error Types**
  - Built-in `error` Type
  - Custom Error Types
- **Error Handling Strategies**
  - Returning Errors from Functions
  - Error Wrapping and Unwrapping (`fmt.Errorf`, `errors` package)
- **Panic and Recover**
  - Using `panic` for Unexpected Errors
  - Using `recover` to Handle Panics

## 7. Advanced Topics

- **Interfaces**
  - Defining and Implementing Interfaces
  - Interface Composition
  - Empty Interface (`interface{}`)
  - Type Assertions and Type Switches
- **Reflection**
  - Using the `reflect` Package
  - Practical Uses of Reflection
- **Testing**
  - Writing Unit Tests with the `testing` Package
  - Table-Driven Tests
  - Test Coverage
  - Benchmarking and Profiling

## 8. Standard Library

- **IO and File Handling**
  - Reading and Writing Files
  - Working with Buffers (`bufio`)
- **Networking**
  - Creating HTTP Servers and Clients (`net/http`)
  - Handling Requests and Responses
  - WebSocket Communication
- **Working with JSON and XML**
  - JSON Encoding and Decoding (`encoding/json`)
  - XML Encoding and Decoding (`encoding/xml`)
- **Time and Date Manipulation**
  - Working with `time` Package
  - Parsing and Formatting Dates and Times
- **Context Package**
  - Using Contexts for Cancellation and Deadlines
  - Propagating Context in Goroutines

## 9. Package Management

#### Go Project Structure

[Project Structure](https://medium.com/evendyne/getting-started-with-go-project-structure-ab8814ded9c3)

```
    myproject/
    ├── cmd/
    │   └── myapp/
    │       └── main.go
    ├── internal/
    │   └── pkgname/
    │       ├── module1/
    │       │   └── module1.go
    │       └── module2/
    │           └── module2.go
    ├── pkg/
    │   └── utils/
    │       └── utils.go
    ├── vendor/
    ├── go.mod
    ├── go.sum
    └── README.md
```

#### Directory Structure Explanation:

1. **`cmd/`**

   - Contains main applications for different services or executables.
   - Each subdirectory under `cmd/` represents a different executable.

2. **`internal/`**

   - Contains private application and library code.
   - Code here is accessible only to the current module (`go.mod`).

3. **`pkg/`**

    - Contains reusable packages and libraries used across different services.
    - Exported packages that can be imported by other modules.

4. **`vendor/`**

   - Contains dependencies (optional).
   - Used if vendoring is enabled (`go mod vendor`).

5. **`go.mod`**

   - Go module file that defines the module's path, dependencies, and version constraints.

6. **`go.sum`**
   - Contains expected cryptographic checksums of the content of specific module versions.
   - Ensures reproducible builds by verifying the integrity of dependencies.

#### Creating and Using Packages

**Package Declaration and Imports**

```go
  // Package declaration
  package main

 // Importing fmt package
  import (
    "fmt"
    "net/http"
   )
```

**Exported and Unexported Identifiers**

- Exported Identifiers: Start with a capital letter, accessible from outside the package.
- Unexported Identifiers: Start with a lowercase letter, accessible only within the same package.

```go
    package mypackage

    func ExportedFunction() { // Exported function
        fmt.Println("Exported function")
    }


    package mypackage

    func unexportedFunction() { // Unexported function
        fmt.Println("Unexported function")
    }
```

#### Modules and Dependency Management

**go.mod and go.sum Files**

`go.mod File`: Defines the module's path and dependencies.

```go
   module example.com/myproject

   go 1.16

   require (
       github.com/some/dependency v1.2.3
   )
```

`go.sum File`: Records cryptographic hashes of module contents for reproducible builds.

```go
   github.com/some/dependency v1.2.3 h1:abcdef...
```

**Initialize a Module**
Create a new module in the current directory

```go
   go mod init example.com/myproject
```

**Adding Dependencies**

```go
   go get github.com/some/dependency@v1.2.3
```

**Updating Dependencies**

```go
   go get -u github.com/some/dependency
```

#### Releasing and Versioning Modules

**Tagging Versions**
Tag module releases with versions.

```go
   git tag v1.0.0
   git push origin v1.0.0
```

**Version Constraints**
Specify version constraints in go.mod for controlled dependencies.

```go
   require github.com/some/dependency v1.2.0 // Use v1.2.0
```

#### Vendoring Dependencies
In Go, vendoring refers to the practice of storing dependencies locally within your project, typically in a directory named `vendor/`. This approach ensures that your project uses specific versions of dependencies, which is crucial for reproducible builds and managing changes over time.

**Enable Vendoring**

```go
    go mod vendor
```

**Updating Vendored Dependencies**

```go
    go mod vendor -u
```

**Cleaning Vendored Dependencies**

```go
    go mod tidy
```



## 10. Building and Deploying Go Applications

- **Compilation and Binary Distribution**
  - `go build`, `go install`, `go run`
  - Creating Executable Binaries
- **Cross-Compilation**
  - Compiling for Different Platforms
- **Deploying Go Applications**
  - Best Practices for Deployment
  - Continuous Integration and Deployment (CI/CD)
- **Dockerizing Go Applications**
  - Creating Dockerfiles for Go Applications
  - Building and Running Go Containers

## 11. Performance Optimization

- **Profiling and Benchmarking**
  - Using `pprof` for Profiling
  - Writing Benchmarks
- **Common Performance Pitfalls**
  - Avoiding Common Mistakes
- **Memory Management and Garbage Collection**
  - Understanding Go’s Garbage Collector
  - Memory Allocation Patterns

## 12. Commonly Used Frameworks and Tools

- **Web Frameworks**
  - Gin, Echo, Fiber
- **ORMs**
  - GORM, SQLBoiler
- **CLI Tools**
  - Cobra, Viper
- **Code Generation**
  - `go generate`, `protoc` for Protocol Buffers
