package main

import (
	"fmt"
	"learn-golang/pkg/util"
)

func main() {
	// Example with integer stack
	intStack := &util.Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	fmt.Println("Integer stack size:", intStack.Size()) // Output: 2

	poppedInt, _ := intStack.Pop()
	fmt.Println("Popped from integer stack:", poppedInt) // Output: 20

	// Example with string stack
	stringStack := &util.Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")
	fmt.Println("String stack size:", stringStack.Size()) // Output: 2

	poppedString, _ := stringStack.Pop()
	fmt.Println("Popped from string stack:", poppedString) // Output: "world"
}
