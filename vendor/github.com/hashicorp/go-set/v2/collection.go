// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package set

// Collection is a minimal common interface that all sets implement.

// Fundamental set operations and familiar utility methods are part of this
// interface. Each of Set, HashSet, and TreeSet may also provide implementation
// specific methods not part of this interface.
type Collection[T any] interface {

	// Insert an element into the set.
	//
	// Returns true if the set is modified as a result.
	Insert(T) bool

	// InsertSlice will insert each element of a given slice into the set.
	//
	// Returns true if the set was modified as a result.
	InsertSlice([]T) bool

	// InsertSet will insert each element of a given Collection into the set.
	//
	// Returns true if the set was modified as a result.
	InsertSet(Collection[T]) bool

	// Remove will remove the given element from the set, if present.
	//
	// Returns true if the set was modified as a result of the operation.
	Remove(T) bool

	// RemoveSlice will remove each element of a slice from the set, if present.
	//
	// Returns true if the set was modified as a result of the operation.
	RemoveSlice([]T) bool

	// RemoveSet will remove each element of a Collection from the set.
	//
	// Returns true if the set was modified as a result of the operation.
	RemoveSet(Collection[T]) bool

	// RemoveFunc will remove each element from the set that satisfies the given predicate.
	//
	// Returns true if the set was modified as a result of the opearation.
	RemoveFunc(func(T) bool) bool

	// Contains returns whether an element is present in the set.
	Contains(T) bool

	// ContainsSlice returns whether the set contains the same set of elements as
	// the given slice. The elements of the slice may contain duplicates.
	ContainsSlice([]T) bool

	// Subset returns whether the given Collection is a subset of the set.
	Subset(Collection[T]) bool

	// ProperSubset returns whether the given Collection is a proper subset of the set.
	ProperSubset(Collection[T]) bool

	// Size returns the number of elements in the set.
	Size() int

	// Empty returns whether the set contains no elements.
	Empty() bool

	// Union returns a new set containing the unique elements of both this set
	// and a given Collection.
	//
	// https://en.wikipedia.org/wiki/Union_(set_theory)
	Union(Collection[T]) Collection[T]

	// Difference returns a new set that contains elements this set that are not
	// in a given Collection.
	//
	// https://en.wikipedia.org/wiki/Complement_(set_theory)
	Difference(Collection[T]) Collection[T]

	// Intersect returns a new set that contains only the elements present in
	// both this and a given Collection.
	//
	// https://en.wikipedia.org/wiki/Intersection_(set_theory)
	Intersect(Collection[T]) Collection[T]

	// Slice returns a slice of all elements in the set.
	//
	// Note: order of elements depends on the underlying implementation.
	Slice() []T

	// String creates a string representation of this set.
	//
	// Note: string representation depends on underlying implementation.
	String() string

	// StringFunc creates a string representation of this set, using the given
	// function to transform each element into a string.
	//
	// Note: string representation depends on underlying implementation.
	StringFunc(func(T) string) string

	// EqualSet returns whether this set and a given Collection contain the same
	// elements.
	EqualSet(Collection[T]) bool

	// EqualSlice returns whether this set and a given slice contain the same
	// elements, where the slice may contain duplicates.
	EqualSlice([]T) bool

	// EqualSliceSet returns whether this set and a given slice contain exactly
	// the same elements, where the slice must not contain duplicates.
	EqualSliceSet([]T) bool

	// ForEach will call the callback function for each element in the set.
	// If the callback returns false, the iteration will stop.
	//
	// Note: iteration order depends on the underlying implementation.
	ForEach(func(T) bool)
}

// InsertSliceFunc inserts all elements from items into col, applying the transform
// function to each element before insertion.
//
// Returns true if col was modified as a result of the operation.
func InsertSliceFunc[T, E any](col Collection[T], items []E, transform func(element E) T) bool {
	modified := false
	for _, item := range items {
		if col.Insert(transform(item)) {
			modified = true
		}
	}
	return modified
}

// InsertSetFunc inserts the elements of a into b, applying the transform function
// to each element before insertion.
//
// Returns true if b was modified as a result.
func InsertSetFunc[T, E any](a Collection[T], b Collection[E], transform func(T) E) bool {
	modified := false
	a.ForEach(func(item T) bool {
		if b.Insert(transform(item)) {
			modified = true
		}
		return true
	})
	return modified
}

// SliceFunc produces a slice of the elements in s, applying the transform
// function to each element first.
func SliceFunc[T, E any](s Collection[T], transform func(T) E) []E {
	slice := make([]E, 0, s.Size())
	s.ForEach(func(item T) bool {
		slice = append(slice, transform(item))
		return true
	})
	return slice
}

func insert[T any](destination, col Collection[T]) {
	col.ForEach(func(item T) bool {
		destination.Insert(item)
		return true
	})
}

func intersect[T any](destination, a, b Collection[T]) {
	var (
		big   Collection[T] = a
		small Collection[T] = b
	)
	if a.Size() < b.Size() {
		big, small = b, a
	}
	small.ForEach(func(item T) bool {
		if big.Contains(item) {
			destination.Insert(item)
		}
		return true
	})
}

func containsSlice[T any](col Collection[T], items []T) bool {
	for _, item := range items {
		if !col.Contains(item) {
			return false
		}
	}
	return true
}

func equalSet[T any](a, b Collection[T]) bool {
	// fast paths: sets are empty or different sizes
	sizeA, sizeB := a.Size(), b.Size()
	if sizeA == 0 && sizeB == 0 {
		return true
	}
	if sizeA != sizeB {
		return false
	}

	// look for any missing element
	different := false
	a.ForEach(func(item T) bool {
		if !b.Contains(item) {
			different = true
			return false // halt
		}
		return true // continue
	})
	return !different
}

func removeSet[T any](s, col Collection[T]) bool {
	modified := false
	col.ForEach(func(item T) bool {
		if s.Remove(item) {
			modified = true
		}
		return true
	})
	return modified
}

func removeFunc[T any](col Collection[T], predicate func(T) bool) bool {
	remove := make([]T, 0)
	col.ForEach(func(item T) bool {
		if predicate(item) {
			remove = append(remove, item)
		}
		return true
	})
	return col.RemoveSlice(remove)
}

func subset[T any](a, b Collection[T]) bool {
	if b.Size() > a.Size() {
		return false
	}
	missing := false
	b.ForEach(func(item T) bool {
		if !a.Contains(item) {
			missing = true
			return false // stop iteration
		}
		return true
	})
	return !missing
}
