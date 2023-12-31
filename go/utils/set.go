package utils

// Set - mathematical set with operations
// Empty struct requires zero bytes so is more efficient than bool
type Set[K comparable] map[K]struct{}

// CreateSet - create an empty set
func CreateSet[K comparable]() Set[K] {
	m := make(map[K]struct{})
	return Set[K](m)
}

// Contains - check if elem in set
func (s Set[K]) Contains(elem K) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[K]) Clone() Set[K] {
	m := make(map[K]struct{})
	for k := range s {
		m[k] = struct{}{}
	}
	return Set[K](m)
}

// Add - add elem to set
func (s Set[K]) Add(elem K) {
	s[elem] = struct{}{}
}

// Remove - remove elem from set (does not need to be in set)
func (s Set[K]) Remove(elem K) {
	delete(s, elem)
}

// Size returns size of set.
func (s Set[K]) Size() int {
	return len(s)
}

// Extend - extend set s with all elements in other (the result is union)
func (s Set[K]) Extend(other Set[K]) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Values returns values in Set
func (s Set[K]) Values() []K {
	v := make([]K, 0, len(s))
	for k := range s {
		v = append(v, k)
	}
	return v
}

// Subtract - remove all elements from s that are in other
func (s Set[K]) Subtract(other Set[K]) {
	for k := range other {
		_, ok := s[k]
		if ok {
			delete(s, k)
		}
	}
}

// Intersect - only keep elements in s which are also in other
func (s Set[K]) Intersect(other Set[K]) {
	deleteList := make([]K, 0, len(s))
	for k := range s {
		_, inOther := other[k]
		if !inOther {
			deleteList = append(deleteList, k)
		}
	}
	for _, k := range deleteList {
		delete(s, k)
	}
}

func (s Set[K]) GetOne() K {
	for k := range s {
		return k
	}
	panic("empty set")
}
