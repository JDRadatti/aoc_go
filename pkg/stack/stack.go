package stack

type Stack[T any] []T

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Empty() {
        return *new(T), false
    }
    last := (*s)[len(*s) - 1]
    *s = (*s)[:len(*s) - 1]
    return last, true
}

func (s Stack[T]) Empty() bool {
	return len(s) == 0
}

