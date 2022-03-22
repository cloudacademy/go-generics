package main

import "fmt"

//attribution https://github.com/danielfurman

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, v := range values {
		m[v] = struct{}{}
	}
	return &Set[T]{
		values: m,
	}
}

func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.values[v] = struct{}{}
	}
}

func (s *Set[T]) Remove(values ...T) {
	for _, v := range values {
		delete(s.values, v)
	}
}

func (s *Set[T]) Contains(values ...T) bool {
	for _, v := range values {
		_, ok := s.values[v]
		if !ok {
			return false
		}
	}
	return true
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet(s.Values()...)
	for _, v := range other.Values() {
		if !result.Contains(v) {
			result.Add(v)
		}
	}
	return result
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	if s.Size() < other.Size() {
		return intersect(s, other)
	}
	return intersect(other, s)
}

func intersect[T comparable](smaller, bigger *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k, _ := range smaller.values {
		if bigger.Contains(k) {
			result.Add(k)
		}
	}
	return result
}

func (s *Set[T]) Values() []T {
	return s.toSlice()
}

func (s *Set[T]) Size() int {
	return len(s.values)
}

func (s *Set[T]) Clear() {
	s.values = map[T]struct{}{}
}

func (s *Set[T]) String() string {
	return fmt.Sprint(s.toSlice())
}

func (s *Set[T]) toSlice() []T {
	result := make([]T, 0, len(s.values))
	for k := range s.values {
		result = append(result, k)
	}
	return result
}

func main() {
	s1 := NewSet(3, 3, -7, 150)
	s2 := NewSet("cloud", "cloud", "academy", "generics", "golang")
	fmt.Println(s1.Size(), s2.Size())

	s1.Add(-200)
	s2.Add("blah")
	fmt.Println(s1.Size(), s2.Size())
	fmt.Println(s1.Contains(3), s2.Contains("blah"))

	s1.Remove(150)
	s2.Remove("cloud")
	fmt.Println(s1.Size(), s2.Size())

	fmt.Println(len(s1.Values()), len(s2.Values()))

	s3 := NewSet("golang", "generics", "constraints")
	fmt.Println(s2.Union(s3).Size())
	fmt.Println(s2.Intersect(s3))

	s1.Clear()
	s2.Clear()
	fmt.Println(s1.Size(), s2.Size())
}
