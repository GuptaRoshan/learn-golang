package util

import "errors"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	top := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]

	return top, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	top := s.elements[len(s.elements)-1]

	return top, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
