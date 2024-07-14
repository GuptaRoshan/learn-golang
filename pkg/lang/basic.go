package lang

import (
	"fmt"
	"strings"
	"unicode"
)

// ----------------------------------------Variables and Types----------------------------------------//

// DeclareVariables demonstrates variable declaration and assignment.
func DeclareVariables() {
	var x int = 5 // Variable declaration and intialization
	fmt.Println("Value of x:", x)

	y := 3.14 // Short variable declaration with type inference
	fmt.Println("Value of y:", y)
}

// VariableTypes demonstrates variable types and type inference.
func VariableTypes() {
	var num int              // Declaration without initialization
	var pi = 3.14159         // Type inferred from initialization
	var isGo = true          // Type bool
	var greet = "Hello, Go!" // Type string

	fmt.Println("num:", num)
	fmt.Println("pi:", pi)
	fmt.Println("isGo:", isGo)
	fmt.Println("greet:", greet)
}

// ConstantsAndIota demonstrates constants and iota in Go.
func ConstantsAndIota() {
	const Pi = 3.14 // Constant declaration
	fmt.Println("Value of Pi:", Pi)

	const (
		A = iota // iota starts at 0
		B        // B = iota (implicitly incremented by 1)
		C        // C = iota (implicitly incremented by 1)
	)
	fmt.Println("A:", A, "B:", B, "C:", C)
}

// BasicTypes demonstrates basic data types in Go.
func BasicTypes() {
	var (
		num   int     = 42
		pi    float64 = 3.14159
		isGo  bool    = true
		greet string  = "Hello, Go!"
	)

	fmt.Println("num:", num)
	fmt.Println("pi:", pi)
	fmt.Println("isGo:", isGo)
	fmt.Println("greet:", greet)
}

// TypeConversions demonstrates type conversions in Go.
func TypeConversions() {
	var x int = 42
	var y float64 = float64(x)
	fmt.Println("x as float64:", y)

	var a float64 = 3.14
	var b int = int(a)
	fmt.Println("a as int:", b)
}

// ----------------------------------------Char or Rune----------------------------------------//
func CharOperations() {
	// String to rune slice
	str := "Hello, 世界"
	runes := []rune(str)
	fmt.Println("Original string:", str)

	// Iterate over runes
	fmt.Print("Runes: ")
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	// Modify a rune
	runes[0] = 'h'
	modifiedStr := string(runes)
	fmt.Println("Modified string:", modifiedStr)

	// Length of string in runes
	fmt.Println("Length of string in runes:", len(runes))

	// Check if a rune is a letter or digit
	r := '1'
	if unicode.IsLetter(r) {
		fmt.Println("r is a letter")
	} else if unicode.IsDigit(r) {
		fmt.Println("r is a digit")
	} else {
		fmt.Println("r is neither a letter nor a digit")
	}

	// Convert rune to string
	r2 := 'g'
	str2 := string(r2)
	fmt.Println("Rune to string:", str2)

	// Convert string to rune
	str3 := "h"
	r3 := []rune(str3)[0]
	fmt.Println("String to rune:", r3)

}

// ----------------------------------------Strings----------------------------------------//

func StringOperations() {
	// String Concatenation
	str1 := "Hello, "
	str2 := "World!"
	result := str1 + str2
	fmt.Println(result) // Output: Hello, World!

	// String Length
	length := len(result)
	fmt.Println(length) // Output: 13

	// Accessing Characters
	char := result[0]
	fmt.Printf("%c\n", char) // Output: H

	// String Comparison
	str3 := "Go"
	str4 := "Go"
	str5 := "Golang"
	fmt.Println(str3 == str4) // Output: true
	fmt.Println(str3 == str5) // Output: false

	// Substrings
	substr := result[7:12]
	fmt.Println(substr) // Output: World

	// Contains
	contains := strings.Contains(result, "World")
	fmt.Println(contains) // Output: true

	// Index
	index := strings.Index(result, "World")
	fmt.Println(index) // Output: 7

	// Replace
	newStr := strings.Replace(result, "World", "Go", 1)
	fmt.Println(newStr) // Output: Hello, Go!

	// Split
	parts := strings.Split(result, ", ")
	fmt.Println(parts) // Output: [Hello World!]

	// ToUpper and ToLower
	upper := strings.ToUpper(result)
	lower := strings.ToLower(result)
	fmt.Println(upper) // Output: HELLO, WORLD!
	fmt.Println(lower) // Output: hello, world!

	// Trim
	strWithSpaces := "  Hello, World!  "
	trimmed := strings.TrimSpace(strWithSpaces)
	fmt.Println(trimmed) // Output: Hello, World!

	// Join
	joined := strings.Join(parts, ", ")
	fmt.Println(joined) // Output: Hello, World!

	/*
	  String to Rune Conversion
	  In Go, a rune is an alias for int32 and represents a Unicode code point. While strings in Go are sequences of bytes (uint8),
	  runes are used to handle Unicode characters, allowing for more efficient and intuitive processing of multi-byte characters.
	*/
	runes := []rune(str1)
	for _, r := range runes {
		fmt.Printf("%c ", r) // Output: H e l l o ,
		fmt.Printf("%U ", r) // unicode
	}

	fmt.Println() // New line for better formatting
}

// ----------------------------------------Operators----------------------------------------//

// OperatorsExpressions demonstrates various operators in Go.
func OperatorsExpressions() {
	// Arithmetic operators
	fmt.Println("5 + 3 =", 5+3)
	fmt.Println("5 - 3 =", 5-3)
	fmt.Println("5 * 3 =", 5*3)
	fmt.Println("5 / 3 =", 5/3)
	fmt.Println("5 % 3 =", 5%3)

	// Comparison operators
	fmt.Println("5 > 3 is", 5 > 3)
	fmt.Println("5 < 3 is", 5 < 3)
	fmt.Println("5 == 3 is", 5 == 3)
	fmt.Println("5 != 3 is", 5 != 3)

	// Logical operators
	fmt.Println("true && false is", true && false)
	fmt.Println("true || false is", true || false)
	fmt.Println("!true is", !true)

	// Bitwise operators
	fmt.Println("5 & 3 =", 5&3)
	fmt.Println("5 | 3 =", 5|3)
	fmt.Println("5 ^ 3 =", 5^3)
	fmt.Println("5 << 1 =", 5<<1)
	fmt.Println("5 >> 1 =", 5>>1)
}

// ----------------------------------------Loops----------------------------------------//

// SwitchStatements demonstrates switch statements in Go.
func SwitchStatements() {
	x := 3

	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}
}

// ForLoops demonstrates for loops in Go.
func ForLoops() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Traditional loop:", i)
	}

	// Range-based for loop
	fruits := []string{"apple", "banana", "cherry"}
	for idx, fruit := range fruits {
		fmt.Printf("Range loop: Index: %d, Fruit: %s\n", idx, fruit)
	}
}

// while and do while loop simulation
func SimulatingWhileAndDoWhileLoop() {
	// Simulating while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// Simulating do-while loop
	j := 1
	for {
		fmt.Println(i)
		j++
		if j > 5 {
			break
		}
	}
}

// DeferBreakContinue demonstrates defer, break, continue, and goto in Go.
func DeferBreakContinue() {
	// Defer statement
	defer fmt.Println("Deferred until the end")
	fmt.Println("Normal execution")

	// Break and continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // Exit loop
		}
		fmt.Println("Break loop:", i)
	}

	for j := 0; j < 5; j++ {
		if j == 2 {
			continue // Skip iteration
		}
		fmt.Println("Continue loop:", j)
	}

	// Goto statement (avoid using goto in practice)
	num := 0
loop:
	for num < 5 {
		fmt.Println("Goto loop:", num)
		num++
		goto loop
	}
}

// ----------------------------------------Functions----------------------------------------//

// DefineCallFunctions demonstrates defining and calling functions in Go.
func DefineCallFunctions() {
	// Function definition
	greet := func() {
		fmt.Println("Hello, World!")
	}

	// Function call
	greet()
}

// ParametersReturnValues demonstrates functions with parameters and return values in Go.
func ParametersReturnValues() {
	// Function with parameters and return value
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println("Result of addition:", result)
}

// NamedReturnValues demonstrates functions with named return values in Go.
func NamedReturnValues() {
	// Function with named return values
	divide := func(dividend, divisor float64) (quotient float64) {
		quotient = dividend / divisor
		return // Return statement without explicit return value
	}

	result := divide(10.0, 2.0)
	fmt.Println("Result of division:", result)
}

// VariadicFunctions demonstrates variadic functions in Go.
func VariadicFunctions() {
	// Variadic function that sums integers
	sum := func(numbers ...int) int {
		total := 0
		for _, num := range numbers {
			total += num
		}
		return total
	}

	result := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers:", result)
}

// AnonymousFunctionsClosures demonstrates anonymous functions and closures in Go.
func AnonymousFunctionsClosures() {
	// Anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	// Closure function
	x := 10
	increment := func() int {
		x++
		return x
	}

	fmt.Println("Closure increment:", increment()) // Output: 11
}
