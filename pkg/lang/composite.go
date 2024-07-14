package lang

import "fmt"

//----------------------------------------Arrays----------------------------------------//

// ArraysExample demonstrates declaring and initializing arrays in Go.
func ArraysExample() {
	// Declaring and initializing arrays
	var arr [5]int            // Array of 5 integers, initialized with zero values
	arr2 := [3]int{1, 2, 3}   // Array of 3 integers with initialization
	arr3 := [...]int{4, 5, 6} // Array with length inferred from elements

	fmt.Println("Array 1:", arr)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)
}

// ArraysAccessModify demonstrates accessing and modifying elements in arrays.
func ArraysAccessModify() {
	arr := [5]int{10, 20, 30, 40, 50}

	// Accessing element at index 0
	fmt.Println("Element at index 0:", arr[0])

	// Modifying element at index 1
	arr[1] = 25
	fmt.Println("Modified array:", arr)
}

//----------------------------------------Slices----------------------------------------//

// SlicesExample demonstrates creating and manipulating slices in Go.
func SlicesExample() {
	// Creating slices from arrays
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // Slice from index 1 to 3 (exclusive)

	// Creating slices directly
	slice2 := []int{6, 7, 8}

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
}

// SliceInternals demonstrates the internals of slices in Go.
func SliceInternals() {
	slice := []int{1, 2, 3, 4, 5}

	// Length of the slice
	fmt.Println("Length of slice:", len(slice))

	// Capacity of the underlying array
	fmt.Println("Capacity of slice:", cap(slice))
}

// SliceOperations demonstrates append, copy, and slice expressions in Go.
func SliceOperations() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	// Appending to a slice
	slice1 = append(slice1, 4, 5)

	// Copying slices
	copy(slice2, slice1) // Copies elements from slice1 to slice2

	fmt.Println("Appended slice 1:", slice1)
	fmt.Println("Copied slice 2:", slice2)
}

//----------------------------------------Maps----------------------------------------//

// MapsExample demonstrates creating and accessing maps in Go.
func MapsExample() {
	// Creating and initializing a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}

	fmt.Println("Ages map:", ages)
}

// MapsOperations demonstrates adding, updating, and deleting map elements in Go.
func MapsOperations() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}

	// Adding a new key-value pair
	ages["Eve"] = 20

	// Updating an existing value
	ages["Bob"] = 31

	// Deleting a key-value pair
	delete(ages, "Alice")

	fmt.Println("Updated ages map:", ages)
}

// IterateMaps demonstrates iterating over maps in Go.
func IterateMaps() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   20,
	}

	// Iterating over map
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}
}

//----------------------------------------Structs----------------------------------------//

// StructsExample demonstrates defining and using structs in Go.
func StructsExample() {
	// Defining a struct
	type Person struct {
		Name string
		Age  int
	}

	// Initializing a struct
	var p1 Person
	p1.Name = "Alice"
	p1.Age = 30

	// Using a struct
	fmt.Println(p1.Name, "is", p1.Age, "years old")
}

// NestedStructsExample demonstrates defining and using nested structs in Go.
func NestedStructsExample() {
	// Defining nested structs
	type Address struct {
		City  string
		State string
	}

	type Person struct {
		Name    string
		Age     int
		Address Address
	}

	// Initializing nested structs
	p := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}

	// Accessing nested struct fields
	fmt.Println(p.Name, "lives in", p.Address.City, "age is", p.Age)
}

// Defining a struct
type Circle struct {
	Radius float64
}

// Area is attached to a Circle struct called receiver function
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// StructMethodsExample demonstrates defining and using struct methods in Go.
func StructMethodsExample() {
	// Initializing a struct
	circle := Circle{Radius: 5}

	// Calling a struct method
	fmt.Println("Area of circle with radius", circle.Radius, "is", circle.Area())
	fmt.Println("Area of circle with perimeter", circle.Perimeter(), "is", circle.Area())
}

// StructMethodsExample demonstrates defining and using struct methods in Go.
func StructMethodsWithAnonymousExample() {

	type Circle struct {
		Radius float64
	}

	// Initializing a struct
	circle := Circle{Radius: 5}

	// Using an anonymous function to calculate the area
	area := func(c Circle) float64 {
		return 3.14 * c.Radius * c.Radius
	}(circle) // Immediately call with circle as argument

	// Printing the calculated area
	fmt.Println("Area of circle with radius", circle.Radius, "is", area)
}
