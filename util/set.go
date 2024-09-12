package util

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	m := make(map[T]struct{})
	return Set[T]{m}
}

func InitSet[T comparable](ls []T) Set[T] {
	s := NewSet[T]()
	s.AddList(ls)
	return s
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}
func (s *Set[T]) AddList(ls []T) {
	for _, v := range ls {
		s.Add(v)
	}
}

func (s *Set[T]) Delete(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}
