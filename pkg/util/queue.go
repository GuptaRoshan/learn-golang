package util

// Queue represents a queue using a slice
type Queue[T any] struct {
	elements []T
}

// Enqueue adds an element to the end of the queue
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zeroValue T
		return zeroValue, false // Queue is empty
	}
	element := q.elements[0]
	q.elements = q.elements[1:] // Remove the first element
	return element, true
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.elements)
}
