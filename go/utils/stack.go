package utils

// Stack
type Stack[T any] struct {
	elems []T
	nr    int
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
	s.nr++
}

// Pop - get element if available as signaled by ok
func (s *Stack[K]) Pop() (elem K, ok bool) {
	if s.nr == 0 {
		return elem, false
	}
	elem = s.elems[s.nr-1]
	s.nr--
	s.elems = s.elems[:s.nr]
	return elem, true
}

func (s *Stack[K]) IsEmpty() bool {
	return s.nr == 0
}

func (s *Stack[K]) Depth() int {
	return s.nr
}

func (s *Stack[K]) Reverse() {
	for i := 0; i < s.nr/2; i++ {
		j := s.nr - 1 - i
		s.elems[i], s.elems[j] = s.elems[j], s.elems[i]
	}
}
