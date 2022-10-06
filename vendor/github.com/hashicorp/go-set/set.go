// Package set provides a basic generic set implementation.
//
// https://en.wikipedia.org/wiki/Set_(mathematics)
package set

import (
	"fmt"
	"sort"
)

type nothing struct{}

var sentinel = nothing{}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// New creates a new Set with initial underlying capacity of size.
//
// A Set will automatically grow or shrink its capacity as items are added or
// removed.
func New[T comparable](size int) *Set[T] {
	return &Set[T]{
		items: make(map[T]nothing, max(0, size)),
	}
}

// From creates a new Set containing each item in items.
func From[T comparable](items []T) *Set[T] {
	s := New[T](len(items))
	s.InsertAll(items)
	return s
}

// Set is a simple, generic implementation of the set mathematical data structure.
// It is optimized for correctness and convenience, as a replacement for the use
// of map[interface{}]struct{}.
type Set[T comparable] struct {
	items map[T]nothing
}

// Insert item into s.
//
// Return true if s was modified (item was not already in s), false otherwise.
func (s *Set[T]) Insert(item T) bool {
	if _, exists := s.items[item]; exists {
		return false
	}
	s.items[item] = sentinel
	return true
}

// InsertAll will insert each item in items into s.
//
// Return true if s was modified (at least one item was not already in s), false otherwise.
func (s *Set[T]) InsertAll(items []T) bool {
	modified := false
	for _, item := range items {
		if s.Insert(item) {
			modified = true
		}
	}
	return modified
}

// InsertSet will insert each element of o into s.
//
// Return true if s was modified (at least one item of o was not already in s), false otherwise.
func (s *Set[T]) InsertSet(o *Set[T]) bool {
	modified := false
	for item := range o.items {
		if s.Insert(item) {
			modified = true
		}
	}
	return modified
}

// Remove will remove item from s.
//
// Return true if s was modified (item was present), false otherwise.
func (s *Set[T]) Remove(item T) bool {
	if _, exists := s.items[item]; !exists {
		return false
	}

	delete(s.items, item)
	return true
}

// RemoveAll will remove each item in items from s.
//
// Return true if s was modified (any item was present), false otherwise.
func (s *Set[T]) RemoveAll(items []T) bool {
	modified := false
	for _, item := range items {
		if s.Remove(item) {
			modified = true
		}
	}
	return modified
}

// RemoveSet will remove each element of o from s.
//
// Return true if s was modified (any item of o was present in s), false otherwise.
func (s *Set[T]) RemoveSet(o *Set[T]) bool {
	modified := false
	for item := range o.items {
		if s.Remove(item) {
			modified = true
		}
	}
	return modified
}

// Contains returns whether item is present in the set.
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// ContainsAll returns whether s contains every item in items.
func (s *Set[T]) ContainsAll(items []T) bool {
	if len(s.items) < len(items) {
		return false
	}

	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// Subset returns whether o is a subset of s.
func (s *Set[T]) Subset(o *Set[T]) bool {
	if len(s.items) < len(o.items) {
		return false
	}

	for item := range o.items {
		if !s.Contains(item) {
			return false
		}
	}

	return true
}

// Size returns the cardinality of s.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Union returns a set that contains all elements of s and o combined.
func (s *Set[T]) Union(o *Set[T]) *Set[T] {
	result := New[T](s.Size())
	for item := range s.items {
		result.items[item] = sentinel
	}
	for item := range o.items {
		result.items[item] = sentinel
	}
	return result
}

// Difference returns a set that contains elements of s that are not in o.
func (s *Set[T]) Difference(o *Set[T]) *Set[T] {
	result := New[T](max(0, s.Size()-o.Size()))
	for item := range s.items {
		if !o.Contains(item) {
			result.items[item] = sentinel
		}
	}
	return result
}

// Intersect returns a set that contains elements that are present in both s and o.
func (s *Set[T]) Intersect(o *Set[T]) *Set[T] {
	result := New[T](0)

	big, small := s, o
	if s.Size() < o.Size() {
		big, small = o, s
	}

	for item := range small.items {
		if big.Contains(item) {
			result.Insert(item)
		}
	}
	return result
}

// Copy creates a copy of s.
func (s *Set[T]) Copy() *Set[T] {
	result := New[T](s.Size())
	for item := range s.items {
		result.items[item] = sentinel
	}
	return result
}

// List creates a copy of s as a slice.
func (s *Set[T]) List() []T {
	result := make([]T, 0, s.Size())
	for item := range s.items {
		result = append(result, item)
	}
	return result
}

// String creates a string representation of s, using f to transform each element
// into a string. The result contains elements sorted by their string order.
func (s *Set[T]) String(f func(element T) string) string {
	l := make([]string, 0, s.Size())
	for item := range s.items {
		l = append(l, f(item))
	}
	sort.Strings(l)
	return fmt.Sprintf("%v", l)
}

// Equal returns whether s and o contain the same elements.
func (s *Set[T]) Equal(o *Set[T]) bool {
	if len(s.items) != len(o.items) {
		return false
	}

	for item := range s.items {
		if !o.Contains(item) {
			return false
		}
	}

	return true
}
