package slice

import (
	"sort"
)

// Stream represents a sequence of elements supporting sequential and parallel aggregate operations.
// Currently, only sequential operations are supported.
type Stream[T any] struct {
	source []T
}

// From creates a new Stream from a slice.
func From[T any](source []T) Stream[T] {
	return Stream[T]{source: source}
}

// Filter returns a stream consisting of the elements of this stream that match the given predicate.
func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	var result []T
	for _, v := range s.source {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return Stream[T]{source: result}
}

// Limit returns a stream consisting of the elements of this stream, truncated to be no longer than maxSize in length.
func (s Stream[T]) Limit(maxSize int64) Stream[T] {
	if int64(len(s.source)) <= maxSize {
		return s
	}
	return Stream[T]{source: s.source[:maxSize]}
}

// Skip returns a stream consisting of the remaining elements of this stream after discarding the first n elements of the stream.
func (s Stream[T]) Skip(n int64) Stream[T] {
	if n >= int64(len(s.source)) {
		return Stream[T]{source: []T{}}
	}
	return Stream[T]{source: s.source[n:]}
}

// Sorted returns a stream consisting of the elements of this stream, sorted according to the provided comparator.
// The comparator should return a negative number if a < b, zero if a == b, and a positive number if a > b.
func (s Stream[T]) Sorted(comparator func(a, b T) int) Stream[T] {
	result := make([]T, len(s.source))
	copy(result, s.source)
	sort.Slice(result, func(i, j int) bool {
		return comparator(result[i], result[j]) < 0
	})
	return Stream[T]{source: result}
}

// Distinct returns a stream consisting of the distinct elements (according to equality) of this stream.
// Note: T must be comparable for this operation.
func (s Stream[T]) Distinct() Stream[T] {
	seen := make(map[any]struct{})
	var result []T
	for _, v := range s.source {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return Stream[T]{source: result}
}

// DistinctBy returns a stream consisting of the distinct elements of this stream,
// where distinctness is determined by the key extracted by the keyExtractor function.
// Note: This is a top-level function because Go methods cannot have type parameters.
func DistinctBy[T any, K comparable](s Stream[T], keyExtractor func(T) K) Stream[T] {
	seen := make(map[K]struct{})
	var result []T
	for _, v := range s.source {
		key := keyExtractor(v)
		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return Stream[T]{source: result}
}

// Collect accumulates the elements of this stream into a slice.
func (s Stream[T]) Collect() []T {
	return s.source
}

// ForEach performs an action for each element of this stream.
func (s Stream[T]) ForEach(action func(T)) {
	for _, v := range s.source {
		action(v)
	}
}

// Reduce performs a reduction on the elements of this stream, using the provided identity value and an associative accumulation function, and returns the reduced value.
func (s Stream[T]) Reduce(identity T, accumulator func(T, T) T) T {
	result := identity
	for _, v := range s.source {
		result = accumulator(result, v)
	}
	return result
}

// Count returns the count of elements in this stream.
func (s Stream[T]) Count() int64 {
	return int64(len(s.source))
}

// AnyMatch returns whether any elements of this stream match the provided predicate.
func (s Stream[T]) AnyMatch(predicate func(T) bool) bool {
	for _, v := range s.source {
		if predicate(v) {
			return true
		}
	}
	return false
}

// AllMatch returns whether all elements of this stream match the provided predicate.
func (s Stream[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range s.source {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// FindFirst returns the first element of this stream and true, or a zero value and false if the stream is empty.
func (s Stream[T]) FindFirst() (T, bool) {
	if len(s.source) == 0 {
		var zero T
		return zero, false
	}
	return s.source[0], true
}
