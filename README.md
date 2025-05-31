# Go (Golang) Learning Plan

[Go Cheat Sheet](https://go.dev/ref/spec)

[Roadmap](https://roadmap.sh/golang)

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

### **Variables and Constants**

**Declaring Variables**

```go
var x int = 10
y := 20 // Short-hand declaration
```

**Variable Types and Type Inference**

```go
var a int = 5       // Explicit type
var b = 3.14        // Type inferred as float64
c := "hello"        // Type inferred as string
```

**Constants and `iota`**

```go
const PI = 3.14159

const (
	A = iota // 0
	B        // 1
	C        // 2
)
```

### **Data Types**

**Basic Types (`int`, `float`, `bool`, `string`)**

```go
var i int = 42
var f float64 = 3.14
var b bool = true
var s string = "GoLang"
```

**Type Conversions**

```go
var x int = 10
var y float64 = float64(x) // Convert int to float64
var z int = int(y)         // Convert float64 to int
```

### **Char or Runes**

**Rune Basics**

```go
var r rune = 'A'
fmt.Println(r)         // 65 (ASCII value)
fmt.Println(string(r)) // "A"

// Unicode rune example
var r2 rune = '你'
fmt.Println(r2)         // Unicode value (20320)
fmt.Println(string(r2)) // "你"
```

**Looping Over a String (Rune by Rune)**

```go
str := "Hello, 世界" // Contains both ASCII and Unicode characters

for i, r := range str {
    fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
}
```

**Convert a String to Rune Slice**

```go
str := "Go语言"
runes := []rune(str) // Convert string to rune slice
fmt.Println("Rune slice:", runes) // [71 111 35821 35328]

for i, r := range runes {
    fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
}
```

**Modify a Rune Slice**

```go
runes := []rune("Hello")

runes[0] = 'M' // Change 'H' to 'M'
fmt.Println(string(runes)) // "Mello"
```

**Count Number of Runes in a String**

```go
str := "Hello, 世界"
fmt.Println("Rune count:", utf8.RuneCountInString(str)) // Output: 9
```

**Check if a Character is a Letter, Digit, or Symbol**

```go
r := '你'
fmt.Println("Is Letter:", unicode.IsLetter(r))   // true
fmt.Println("Is Digit:", unicode.IsDigit(r))     // false
fmt.Println("Is Symbol:", unicode.IsSymbol(r))   // false
```

**Convert Uppercase to Lowercase (or vice versa)**

```go
r := 'A'
lower := unicode.ToLower(r)
upper := unicode.ToUpper('g')

fmt.Println(string(lower)) // "a"
fmt.Println(string(upper)) // "G"
```

**Remove Non-Letter Characters from a String**

```go
str := "G0!o@#d123"
var cleanRunes []rune

for _, r := range str {
    if unicode.IsLetter(r) {
        cleanRunes = append(cleanRunes, r)
    }
}

fmt.Println(string(cleanRunes)) // "Good"
```

### **Strings**

**String Operations**

```go
s1 := "Hello"
s2 := " World"
s3 := s1 + s2 // Concatenation
fmt.Println(s3)       // "Hello World"
fmt.Println(len(s3))  // String length
fmt.Println(s3[0])    // Access character (byte representation, ASCII of 'H')

// Accessing a character as a string
fmt.Println(string(s3[0])) // "H"

// String slicing
substr := s3[0:5] // Extracts "Hello"
fmt.Println(substr)

// Convert string to a slice of bytes
byteArr := []byte(s3)
fmt.Println(byteArr) // [72 101 108 108 111 32 87 111 114 108 100]

// Convert byte slice back to a string
newStr := string(byteArr)
fmt.Println(newStr) // "Hello World"

// Convert string to a slice of runes
runeArr := []rune(s3)
fmt.Println(runeArr) // Unicode points

// Convert rune slice back to a string
fmt.Println(string(runeArr)) // "Hello World"

// String comparison
fmt.Println(strings.Compare("abc", "xyz")) // -1 (less than)
fmt.Println(strings.Compare("abc", "abc")) // 0 (equal)
fmt.Println(strings.Compare("xyz", "abc")) // 1 (greater than)

// String Contains
fmt.Println(strings.Contains(s3, "World")) // true

// String Prefix and Suffix Check
fmt.Println(strings.HasPrefix(s3, "Hel"))  // true
fmt.Println(strings.HasSuffix(s3, "ld"))   // true

// String Replacement
fmt.Println(strings.Replace(s3, "World", "Golang", 1)) // "Hello Golang"

// String Split
splitStr := strings.Split(s3, " ")
fmt.Println(splitStr) // ["Hello", "World"]

// String Join
joinedStr := strings.Join(splitStr, "-")
fmt.Println(joinedStr) // "Hello-World"

// Uppercase and Lowercase Conversion
fmt.Println(strings.ToUpper(s3)) // "HELLO WORLD"
fmt.Println(strings.ToLower(s3)) // "hello world"

// Trimming Spaces
trimStr := "   GoLang   "
fmt.Println(strings.TrimSpace(trimStr)) // "GoLang"

// Finding Substring Index
fmt.Println(strings.Index(s3, "World")) // 6

// Repeat a String
fmt.Println(strings.Repeat("Go", 3)) // "GoGoGo"
```

### **Operators and Expressions**

**Arithmetic, Comparison, Logical, and Bitwise Operators**

```go
a, b := 10, 5

// Arithmetic
sum := a + b
diff := a - b
prod := a * b
quot := a / b
mod := a % b

// Comparison
fmt.Println(a > b)  // true

// Logical
fmt.Println(true && false) // false

// Bitwise
fmt.Println(a & b) // Bitwise AND
fmt.Println(a | b) // Bitwise OR
```

### **Control Structures**

**`if-else` Statements**

```go
x := 10
if x > 5 {
    fmt.Println("Greater than 5")
} else {
    fmt.Println("5 or less")
}
```

**`switch` Statements**

```go
switch day := 3; day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Other day")
}
```

**`for` Loops (Traditional, Range)**

```go
// Traditional loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Range loop
arr := []int{10, 20, 30}
for i, v := range arr {
    fmt.Println(i, v)
}
```

**`defer`, `break`, `continue`, and `goto`**

```go
// defer
defer fmt.Println("This runs last")

// break
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}

// continue
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}

// goto
x := 1
goto label
label:
fmt.Println("Jumped here")
```

### **Functions**

**Defining and Calling Functions**

```go
func add(a int, b int) int {
    return a + b
}

result := add(5, 10)
fmt.Println(result)
```

**Parameters and Return Values**

```go
func multiply(a, b int) (int, int) {
    return a * b, a + b
}

prod, sum := multiply(2, 3)
```

**Named Return Values**

```go
func divide(a, b float64) (quotient float64) {
    quotient = a / b
    return
}
```

**Variadic Functions**

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4))
```

**Anonymous Functions and Closures**

```go
func main() {
    // Anonymous function
    sum := func(a, b int) int {
        return a + b
    }
    fmt.Println(sum(3, 4))

    // Closure
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()

    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
}
```


## 3. Composite Types

### **1. Arrays**

**Declaring and Initializing Arrays**

```go
  // Declare an array with explicit size
  var arr1 [5]int // Default values: [0, 0, 0, 0, 0]
  fmt.Println(arr1)

  // Inline initialization
  arr2 := [3]int{1, 2, 3}
  fmt.Println(arr2) // Output: [1, 2, 3]

  // Auto-sizing based on values
  arr3 := [...]string{"Go", "Rust", "Python"}
  fmt.Println(arr3) // Output: [Go Rust Python]
```

**Accessing and Modifying Array Elements**

```go
arr := [4]int{10, 20, 30, 40}
fmt.Println(arr[1]) // Access index 1 → Output: 20

arr[2] = 100 // Modify index 2
fmt.Println(arr) // Output: [10, 20, 100, 40]

// Loop through an array
for i, v := range arr {
    fmt.Println("Index:", i, "Value:", v)
}
```


### **2. Slices**

**Creating and Manipulating Slices**

```go
// Create a slice from an array
arr := [5]int{1, 2, 3, 4, 5}
slice1 := arr[1:4] // Creates [2, 3, 4]
fmt.Println(slice1)

// Create a slice using make()
slice2 := make([]int, 3, 5) // Length 3, Capacity 5
fmt.Println(slice2) // [0 0 0]
```

**Slice Internals and Capacity**

```go
s := []int{1, 2, 3}
fmt.Println(len(s), cap(s)) // Length: 3, Capacity: 3

s = append(s, 4, 5)
fmt.Println(len(s), cap(s)) // Length: 5, Capacity: 6 (may double)
```

**Append, Copy, and Slice Expressions**

```go
// Append elements
s := []int{1, 2, 3}
s = append(s, 4, 5)
fmt.Println(s) // Output: [1 2 3 4 5]

// Copy slices
s1 := []int{10, 20, 30}
s2 := make([]int, len(s1))
copy(s2, s1)
fmt.Println(s2) // Output: [10 20 30]

// Slice expressions
s3 := s1[:2]  // First two elements
s4 := s1[1:]  // Elements from index 1 onward
fmt.Println(s3, s4)
```

### **3. Maps**

**Creating and Accessing Maps**

```go
// Create a map
m := make(map[string]int)
m["apple"] = 5
m["banana"] = 8

fmt.Println(m["apple"]) // Output: 5
```

**Adding, Updating, and Deleting Map Elements**

```go
m := map[string]int{"one": 1, "two": 2}

// Add or update value
m["three"] = 3
m["one"] = 100

// Delete a key
delete(m, "two")

// Check if key exists
value, exists := m["one"]
if exists {
    fmt.Println("Value:", value)
}
```

**Iterating Over Maps**

```go
m := map[string]int{"a": 10, "b": 20, "c": 30}

for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```

### **4. Structs**

**Defining and Using Structs**

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println(p.Name, p.Age)
}
```

**Nested Structs**

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name    string
    Age     int
    Address Address
}

func main() {
    emp := Employee{Name: "John", Age: 30, Address: Address{City: "NY", State: "USA"}}
    fmt.Println(emp.Name, emp.Address.City)
}
```

**Struct Methods and Receivers**

```go
type Car struct {
    Brand string
    Speed int
}

// Method with receiver
func (c Car) Info() {
    fmt.Println("Brand:", c.Brand, "Speed:", c.Speed)
}

// Method with pointer receiver (modifies struct)
func (c *Car) Accelerate(increase int) {
    c.Speed += increase
}

func main() {
    car := Car{"Toyota", 100}
    car.Info()
    
    car.Accelerate(20)
    car.Info() // Updated speed
}
```


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
  > Go has a built-in error type for handling errors, and you can also define custom error types.
  - Built-in `error` Type
    > Go has a built-in error type for handling errors, and you can also define custom error types.'
    ```go
      package main
      
      import (
      	"errors"
      	"fmt"
      )
      
      func divide(a, b int) (int, error) {
      	if b == 0 {
      		return 0, errors.New("cannot divide by zero")
      	}
      	return a / b, nil
      }
      
      func main() {
      	result, err := divide(4, 0)
      	if err != nil {
      		fmt.Println("Error:", err)
      		return
      	}
      	fmt.Println("Result:", result)
      }
    ```
  - Custom Error Types
    > You can define custom error types by implementing the Error() method.
    ```go
      package main
      
      import (
      	"fmt"
      )
      
      type MyError struct {
      	Function string
      	Message  string
      }
      
      func (e *MyError) Error() string {
      	return fmt.Sprintf("%s: %s", e.Function, e.Message)
      }
      
      func riskyFunction() error {
      	return &MyError{Function: "riskyFunction", Message: "something went wrong"}
      }
      
      func main() {
      	err := riskyFunction()
      	if err != nil {
      		fmt.Println("Error:", err)
      	}
      }
    ```
- **Error Handling Strategies**
  > Error handling strategies in Go include checking for errors and handling them immediately, wrapping errors, and using panic and recover for unexpected situations.
  - Returning Errors from Functions
    > Functions in Go often return an error as the last return value.
    ```go
      package main
      
      import (
      	"fmt"
      	"os"
      )
      
      func readFile(filename string) ([]byte, error) {
      	data, err := os.ReadFile(filename)
      	if err != nil {
      		return nil, err
      	}
      	return data, nil
      }
      
      func main() {
      	data, err := readFile("example.txt")
      	if err != nil {
      		fmt.Println("Error:", err)
      		return
      	}
      	fmt.Println("File content:", string(data))
      }
    ```
  - Error Wrapping and Unwrapping (`fmt.Errorf`, `errors` package)
    > Wrapping errors with additional context can be done using fmt.Errorf and the errors package.
    ```go
      package main
      
      import (
      	"errors"
      	"fmt"
      )
      
      func readFile(filename string) ([]byte, error) {
      	// Simulate an error
      	return nil, errors.New("file not found")
      }
      
      func processFile(filename string) error {
      	data, err := readFile(filename)
      	if err != nil {
      		return fmt.Errorf("processing file %s: %w", filename, err)
      	}
      	fmt.Println("File content:", string(data))
      	return nil
      }
      
      func main() {
      	err := processFile("example.txt")
      	if err != nil {
      		fmt.Println("Error:", err)
      		unwrappedErr := errors.Unwrap(err)
      		fmt.Println("Unwrapped Error:", unwrappedErr)
      	}
      }
    ```
- **Panic and Recover**
  > panic is used for handling unexpected errors, and recover is used to regain control after a panic.
  - Using `panic` for Unexpected Errors
    > Use panic when an error should stop the normal execution of the program.
    ```go
      package main
      
      import (
      	"fmt"
      )
      
      func checkPositive(number int) {
      	if number < 0 {
      		panic("number is negative")
      	}
      	fmt.Println("Number is positive")
      }
      
      func main() {
      	checkPositive(10)
      	checkPositive(-1)
      }
    ```
  - Using `recover` to Handle Panics
    > recover allows you to regain control after a panic.
    ```go
      package main
      
      import (
      	"fmt"
      )
      
      func checkPositive(number int) {
      	defer func() {
      		if r := recover(); r != nil {
      			fmt.Println("Recovered from panic:", r)
      		}
      	}()
      	if number < 0 {
      		panic("number is negative")
      	}
      	fmt.Println("Number is positive")
      }
      
      func main() {
      	checkPositive(10)
      	checkPositive(-1)
      	fmt.Println("Program continues...")
      }
    ```

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

**cmd/**  

**Purpose:**  
This directory contains the main applications for different services or executables in a Go project. Each subdirectory inside `cmd/` represents a separate executable program.

**Example:**  
Imagine you are building a project with a REST API and a command-line tool.

```
myproject/
│── cmd/
│   ├── api/
│   │   ├── main.go   # Entry point for the API service
│   ├── cli/
│   │   ├── main.go   # Entry point for the CLI tool
```

- `cmd/api/main.go`: Starts the API service.  
- `cmd/cli/main.go`: Starts a command-line tool for interacting with the API.  


**internal/**  

**Purpose:**  
Contains private application and library code that **cannot** be imported by other modules. Only the current Go module can access this.

**Example:**  
Let's say your project has business logic that should only be used internally.

```
myproject/
│── internal/
│   ├── auth/
│   │   ├── jwt.go   # Contains JWT authentication logic
│   ├── database/
│   │   ├── db.go    # Handles database connections
```

- `internal/auth/jwt.go`: Authentication logic used within the project.  
- `internal/database/db.go`: Database handling, not accessible to other Go modules.  

**Why use `internal/`?**  
It prevents accidental use of internal packages by external projects.


**pkg/** 

**Purpose:**  
Contains reusable Go packages that **can be imported** by other modules or services.

**Example:**  
Let's say you have a logging package that should be shared across multiple services.

```
myproject/
│── pkg/
│   ├── logger/
│   │   ├── logger.go  # Contains a reusable logging utility
```

- `pkg/logger/logger.go`: A reusable logging package that can be used in `cmd/api/` and `cmd/cli/`.  
- Other Go projects can import `myproject/pkg/logger`.  


 **vendor/**  

**Purpose:**  
Contains third-party dependencies if vendoring is enabled (`go mod vendor`). It ensures all dependencies are locally available.

**Example:**  
If your project relies on external libraries (e.g., `github.com/gin-gonic/gin` for an API framework), running:

```sh
go mod vendor
```

Creates:

```
myproject/
│── vendor/
│   ├── github.com/
│   │   ├── gin-gonic/
│   │   │   ├── gin/    # Local copy of the Gin framework
```

- Helps in environments where external dependencies shouldn't be fetched from the internet.  


**go.mod**  

**Purpose:**  
Defines the module name, dependencies, and version constraints.

**Example (`go.mod` file):**
```go
module myproject

go 1.20

require (
    github.com/gin-gonic/gin v1.8.1
)
```
- Declares `myproject` as the module.  
- Specifies `go 1.20` as the Go version.  
- Lists dependencies (`gin` framework at version `v1.8.1`).  


**go.sum**  

**Purpose:**  
Contains cryptographic checksums to verify dependencies, ensuring reproducible builds.

**Example (`go.sum` file):**
```
github.com/gin-gonic/gin v1.8.1 h1:abc123...
github.com/gin-gonic/gin v1.8.1/go.mod h1:def456...
```
- Ensures that `gin v1.8.1` is always the exact same version when installed.  


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

  - **Exported Identifiers**: Start with a capital letter, accessible from outside the package.
  - **Unexported Identifiers**: Start with a lowercase letter, accessible only within the same package.

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

- **Initialize a Module** - Create a new module in the current directory

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

- **Tagging Versions** - Tag module releases with versions.

  ```go
  git tag v1.0.0
  git push origin v1.0.0
  ```

- **Version Constraints** - Specify version constraints in go.mod for controlled dependencies.

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
