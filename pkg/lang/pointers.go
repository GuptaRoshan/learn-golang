package lang

import "fmt"

// PointerExample demonstrates declaring and using pointers in Go.
func PointerExample() {
	// Declare a variable and a pointer
	num := 10

	// Pointer vraiable for storing address
	var ptr *int = &num

	// Print the value and address of num
	fmt.Println("Value of num:", num)    // Output: 10
	fmt.Println("Address of num:", &num) // Output: Address in memory
	fmt.Println("Value of ptr:", ptr)    // Output: Address of num
	fmt.Println("Value at ptr:", *ptr)   // Output: 10 (Dereferencing)
}

// PointerVsValueTypesExample demonstrates the difference between pointer and value types in Go.
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

// PointerToStructsExample demonstrates using pointers to structs in Go.
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

// PointerToFunctionsExample demonstrates using pointers to functions in Go.
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
