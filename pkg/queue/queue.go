package queue

// FIFO Queue
type Queue[T any] []T

func (s *Queue[T]) Enqueue(elem T) {
	*s = append(*s, elem)
}

func (s *Queue[T]) Dequeue() (T, bool) {
	if s.Empty() {
        return *new(T), false
    }
    first := (*s)[0]
    *s = (*s)[1:]
    return first, true
}

func (s Queue[T]) Empty() bool {
	return len(s) == 0
}
