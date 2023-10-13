package collections

import (
	"fmt"
	"strings"
)

// Collection that stores a set of homogenous elements.
type Set[T CollectionElement] struct {
	items map[T]struct{}
}

// Factory method to create an empty set with predefined capacity.
func NewSet[T CollectionElement]() *Set[T] {
	s := &Set[T]{}
	s.items = make(map[T]struct{})
	return s
}

// Factory method to create a set from an array.
func ToSet[T CollectionElement](array []T) *Set[T] {
	s := &Set[T]{}
	s.items = make(map[T]struct{}, len(array))
	for _, item := range array {
		s.items[item] = struct{}{}
	}
	return s
}

// Add element to set.
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Checks whether an element is present in the set.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Updates the current set to be it's union with another set.
func (s *Set[T]) Extend(s2 *Set[T]) {
	for key := range s2.items {
		s.items[key] = struct{}{}
	}
}

// Removes all elements from the set.
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Removes occurence of the given element from the set.
// Returns an error if the element is not present in the set.
func (s *Set[T]) Remove(item T) error {
	if _, ok := s.items[item]; !ok {
		return ErrItemNotFound
	}
	delete(s.items, item)
	return nil
}

// Returns an slice containing elements of the set.
func (s *Set[T]) ToArray() []T {
	result := make([]T, 0, s.Size())
	for key := range s.items {
		result = append(result, key)
	}
	return result
}

// Returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s.items)
}

// Returns a string description of the set.
func (s Set[T]) String() string {
	resultStrings := make([]string, 0, s.Size())
	for e := range s.items {
		resultStrings = append(resultStrings, fmt.Sprint(e))
	}

	return "{" + strings.Join(resultStrings, ",") + "}"
}

func (_ Set[T]) Type() CollectionType {
	return TypeSet
}
