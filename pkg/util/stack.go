package util

import "errors"

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element of the stack. Returns an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top, nil
}

// Peek returns the top element of the stack without removing it. Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}
	top := s.elements[len(s.elements)-1]
	return top, nil
}

// IsEmpty checks if the stack has no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
