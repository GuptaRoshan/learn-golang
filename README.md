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
- **Char or Runes**
  - char operations
- **Strings**
  - String operations
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

Pointers are a powerful feature in Go that allow you to directly manipulate memory addresses. A pointer holds the memory address of a value rather than the value itself. 

- **Declaring and Dereferencing Pointers**

  ```go

      var a int = 42
      // Pointer Declaration var p *int, assigning address &a of a
      var p *int = &a

      // Print the value of a
      fmt.Println("Value of a:", a) // Output: 42

      // Print the address of a
      fmt.Println("Address of a:", &a)

      // Print the value of the pointer p (which is the address of a)
      fmt.Println("Value of pointer p:", p)

      // Dereference the pointer to get the value stored at the address
      fmt.Println("Value at the address p points to:", *p) // Output: 42

      // Modify the value at the address using the pointer
      *p = 100
      fmt.Println("New value of a:", a) // Output: 100

  ```

- **Visual Explanation**

  1. Initial State : `a is an integer variable storing the value 42. p is a pointer to an integer.`

  ```go
      +----+     +------+
      | a  |     |  p   |
      +----+     +------+
      | 42 |     | nil  |
      +----+     +------+
  ```    

  2. Assign the Address of a to p : p = `&a assigns the address of a to p.`

  ```go
      +----+     +------+
      | a  |     |  p   |
      +----+     +------+
      | 42 |     | &a   |
      +----+     +------+
  ```    


  3. Dereference p to Access the Value of a : `*p gives you access to the value stored at the address p points to (which is a).`

  ```go
      +----+     +------+
      | a  |     |  p   |
      +----+     +------+
      | 42 |     | &a   |
      +----+     +------+
        ^           |
        |           |
        +-----------+
  ```

  4. Modify the Value at the Address : `*p = 100 changes the value at the address p points to, which modifies a.`

  ```go
      +----+     +------+
      | a  |     |  p   |
      +----+     +------+
      |100 |     | &a   |
      +----+     +------+
        ^           |
        |           |
        +-----------+
  ```      

- **Pass by copy and reference**

  ```go

    func PointerVsValueTypesExample() {
      // Value type
      var x int = 10
      y := x // y is a copy of x

      y = 20
      _ = y                // mark y unsused to avoid compile error
      fmt.Println("x:", x) // Output: 10 (unchanged)

      // Pointer type
      var ptr *int = &x
      *ptr = 30

      fmt.Println("x:", x) // Output: 30 (changed)
    }

  ```

-  **Pointers to Structs**

    ```go

      func PointerToStructsExample() {
        // Define a struct
        type Person struct {
          Name string
          Age  int
        }

        // Create an instance of Person using pointer
        var p *Person = &Person{
          Name: "Alice",
          Age:  30,
        }

        // Access fields using pointer dereferencing
        fmt.Println("Name:", p.Name) // Output: Alice
        fmt.Println("Age:", p.Age)   // Output: 30
      }

    ```

-  **Pointers to Function**

    ```go

      func PointerToFunctionsExample() {
        // Define a function type
        type MathFunc func(int, int) int

        // Function to add two numbers
        add := func(a, b int) int {
          return a + b
        }

        // Function to subtract two numbers
        subtract := func(a, b int) int {
          return a - b
        }

        // Pointer to functions
        var mathFunc MathFunc
        mathFunc = add

        // Call the function through the pointer
        result := mathFunc(10, 5)
        fmt.Println("Result of addition:", result) // Output: 15

        // Change pointer to subtract function
        mathFunc = subtract
        result = mathFunc(10, 5)
        fmt.Println("Result of subtraction:", result) // Output: 5
      }

    ```



## 5. Concurrency in Go

- **Goroutines**
  - Creating and Running Goroutines
     > A goroutine is a lightweight thread managed by the Go runtime. You create a goroutine by prefixing a function call with the go keyword.

    ```go
        package main
        
        import (
            "fmt"
            "time"
        )
        
        func sayHello() {
            fmt.Println("Hello")
        }
        
        func main() {
            go sayHello() // Launches sayHello as a goroutine
            time.Sleep(1 * time.Second) // Give the goroutine time to complete
        }
    ```
    
  - Goroutine Lifecycle
      >  A goroutine starts with the go keyword and ends when the function completes. They run concurrently with other goroutines.

    ```go
        package main
        
        import (
            "fmt"
            "time"
        )
        
        func main() {
            for i := 0; i < 3; i++ {
                go func(i int) {
                    fmt.Println(i)
                }(i)
            }
            time.Sleep(1 * time.Second)
        }
     ```

- **Channels**
    > Channels provide a way for goroutines to communicate and synchronize. They can be used to send and receive values between goroutines.
  - Creating and Using Channels
      > Channels are created using the make function.
    
    ```go
        package main
        
        import "fmt"
        
        func main() {
            ch := make(chan int)
        
            go func() {
                ch <- 42 // Send value to channel
            }()
        
            value := <-ch // Receive value from channel
            fmt.Println(value) // Output: 42
        }
    ```
    
  - Unbuffered Channels
       > Creation: `ch := make(chan int)`
    
       > Behavior: An unbuffered channel has zero capacity. This means that a send operation on an unbuffered channel blocks the sending goroutine until another goroutine is ready to receive the value. Similarly, a receive operation blocks until a value is sent.
    
      > Use cases: Unbuffered channels are useful for synchronizing goroutines. They enforce a strict one-to-one communication pattern, ensuring that the sender and receiver are coordinated at a specific point in their execution.
      
      ```go
        package main
        
        import "fmt"
        
        func main() {
            ch := make(chan int)
        
            go func() {
                fmt.Println("Sending value...")
                ch <- 42
                fmt.Println("Value sent!")
            }()
        
            fmt.Println("Receiving value...")
            value := <-ch
            fmt.Println("Value received:", value)
        }
      ```
    
  -  Buffered Channels

     > Creation: `ch := make(chan int, 2) (capacity of 2)`
  
     > Behavior: A buffered channel has a specified capacity to hold values. Send operations on a buffered channel only block when the channel is full. Receive operations block only when the channel is empty.
  
     > Use cases: Buffered channels are used when you want to create a pipeline-like pattern, where the sender can send multiple values without waiting for each one to be received immediately. This can improve performance by decoupling the sender and receiver.

      ```go
        package main
        
        import "fmt"
        
        func main() {
            ch := make(chan int, 2)
        
            ch <- 1
            ch <- 2
        
            fmt.Println(<-ch)
            fmt.Println(<-ch)
        }
      ```

  - Sending and Receiving Data  
      > Data is sent to and received from channels using the <- operator.

    ```go
      package main
      
      import "fmt"
      
      func main() {
          ch := make(chan string)
      
          go func() {
              ch <- "Hello from goroutine"
          }()
      
          message := <-ch
          fmt.Println(message) // Output: Hello from goroutine
      }
    ```
  
  - Channel Directions
     > In Go, channels can be directional, meaning they can be restricted to only sending or receiving values. This can help enforce correct usage patterns and make your code more readable and safer. There are three types of channel directions:
  
     > 1. Bidirectional channels: Can be used for both sending and receiving values.
     > 2. Send-only channels: Can only be used for sending values.
     > 3. Receive-only channels: Can only be used for receiving values.

    - Bidirectional channels
       > By default, channels are bidirectional, meaning they can be used to both send and receive values.
  
       ```go
          package main
    
          import "fmt"
          
          func main() {
              ch := make(chan int) // Create a bidirectional channel
          
              go func() {
                  ch <- 42 // Send value to the channel
              }()
          
              value := <-ch // Receive value from the channel
              fmt.Println("Received value:", value) // Output: 42
          }
       ```
   
    - Send-only channels
       > A send-only channel can only be used to send values. This is specified by using the chan<- syntax.
  
       ```go
          package main
          
          import "fmt"
          
          func send(ch chan<- int, value int) {
              ch <- value // Send value to the channel
          }
          
          func main() {
              ch := make(chan int)
          
              go send(ch, 42) // Send value using a send-only channel
          
              value := <-ch // Receive value from the channel
              fmt.Println("Received value:", value) // Output: 42
          }
       ```
 
  
    - Receive-only channels
      > A receive-only channel can only be used to receive values. This is specified by using the <-chan syntax.
  
       ```go
          package main
          
          import "fmt"
          
          func receive(ch <-chan int) {
              value := <-ch // Receive value from the channel
              fmt.Println("Received value:", value)
          }
          
          func main() {
              ch := make(chan int)
          
              go func() {
                  ch <- 42 // Send value to the channel
              }()
          
              receive(ch) // Receive value using a receive-only channel
          }
       ```
     
- **Select Statement**
  - Using `select` for Multiplexing
    > The select statement is used for multiplexing channels. It allows a goroutine to wait on multiple communication operations, proceeding with the first one that becomes ready.
    ```go
      package main
      
      import (
      	"fmt"
      	"time"
      )
      
      func main() {
      	ch1 := make(chan string)
      	ch2 := make(chan string)
      
      	go func() {
      		time.Sleep(1 * time.Second)
      		ch1 <- "from ch1"
      	}()
      
      	go func() {
      		time.Sleep(2 * time.Second)
      		ch2 <- "from ch2"
      	}()
      
      	for i := 0; i < 2; i++ {
      		select {
      		case msg1 := <-ch1:
      			fmt.Println(msg1)
      		case msg2 := <-ch2:
      			fmt.Println(msg2)
      		}
      	}
      }
    ```
- **Synchronization (sync package)**
  > The sync package provides synchronization primitives, such as WaitGroups, Mutexes, RWMutexes, and Once.
  - WaitGroups
    > WaitGroups are used to wait for a collection of goroutines to finish executing.
    ```go
      package main
      
      import (
      	"fmt"
      	"sync"
      )
      
      func worker(id int, wg *sync.WaitGroup) {
      	defer wg.Done()
      	fmt.Printf("Worker %d starting\n", id)
      	// Simulate work
      	fmt.Printf("Worker %d done\n", id)
      }
      
      func main() {
      	var wg sync.WaitGroup
      
      	for i := 1; i <= 3; i++ {
      		wg.Add(1)
      		go worker(i, &wg)
      	}
      
      	wg.Wait()
      }
    ```
  - Mutexes (`sync.Mutex`)
    > Mutexes provide a way to ensure that only one goroutine accesses a critical section of code at a time.
    ```go
      package main
  
      import (
        "fmt"
        "sync"
      )
      
      var counter int
      var mu sync.Mutex
      
      func increment(wg *sync.WaitGroup) {
        defer wg.Done()
        mu.Lock()
        counter++
        mu.Unlock()
      }
      
      func main() {
        var wg sync.WaitGroup
      
        for i := 0; i < 1000; i++ {
          wg.Add(1)
          go increment(&wg)
        }
      
        wg.Wait()
        fmt.Println("Final counter value:", counter)
      }
    ```
  - RWMutex (`sync.RWMutex`)
    > RWMutex is a reader/writer mutual exclusion lock. It allows multiple readers or one writer at a time.
    ```go
      package main
      
      import (
      	"fmt"
      	"sync"
      )
      
      var counter int
      var rw sync.RWMutex
      
      func read(wg *sync.WaitGroup) {
      	defer wg.Done()
      	rw.RLock()
      	fmt.Println("Read counter value:", counter)
      	rw.RUnlock()
      }
      
      func write(wg *sync.WaitGroup) {
      	defer wg.Done()
      	rw.Lock()
      	counter++
      	fmt.Println("Incremented counter value:", counter)
      	rw.Unlock()
      }
      
      func main() {
      	var wg sync.WaitGroup
      
      	for i := 0; i < 5; i++ {
      		wg.Add(1)
      		go write(&wg)
      		wg.Add(1)
      		go read(&wg)
      	}
      
      	wg.Wait()
      }
    ```
  - Once (`sync.Once`)
    > Once ensures that a piece of code is executed only once, no matter how many times the Do method is called.
    ```go
       package main
  
      import (
        "fmt"
        "sync"
      )
      
      var once sync.Once
      
      func initialize() {
        fmt.Println("Initialized")
      }
      
      func worker(wg *sync.WaitGroup) {
        defer wg.Done()
        once.Do(initialize)
      }
      
      func main() {
        var wg sync.WaitGroup
      
        for i := 0; i < 10; i++ {
          wg.Add(1)
          go worker(&wg)
        }
      
        wg.Wait()
      }
    ```

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

### Go Project Structure

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

### Directory Structure Explanation:

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

### Creating and Using Packages

- **Package Declaration and Imports**

  ```go
    // Package declaration
    package main

  // Importing fmt package
    import (
      "fmt"
      "net/http"
    )
  ```

- **Exported and Unexported Identifiers**

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

### Modules and Dependency Management

- **go.mod and go.sum Files**

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

- **Initialize a Module**
  Create a new module in the current directory

  ```go
    go mod init example.com/myproject
  ```

- **Adding Dependencies**

  ```go
    go get github.com/some/dependency@v1.2.3
  ```

- **Updating Dependencies**

  ```go
    go get -u github.com/some/dependency
  ```

### Releasing and Versioning Modules

- **Tagging Versions**
  Tag module releases with versions.

  ```go
    git tag v1.0.0
    git push origin v1.0.0
  ```

- **Version Constraints**
  Specify version constraints in go.mod for controlled dependencies.

  ```go
    require github.com/some/dependency v1.2.0 // Use v1.2.0
  ```

### Vendoring Dependencies
In Go, vendoring refers to the practice of storing dependencies locally within your project, typically in a directory named `vendor/`. This approach ensures that your project uses specific versions of dependencies, which is crucial for reproducible builds and managing changes over time.

- **Enable Vendoring**

  ```go
      go mod vendor
  ```

- **Updating Vendored Dependencies**

  ```go
      go mod vendor -u
  ```

- **Cleaning Vendored Dependencies**

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
