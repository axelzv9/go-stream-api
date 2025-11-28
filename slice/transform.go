package slice

// Map returns a stream consisting of the results of applying the given function to the elements of this stream.
// Note: This is a top-level function because Go methods cannot have type parameters.
func Map[T, R any](s Stream[T], mapper func(T) R) Stream[R] {
	result := make([]R, len(s.source))
	for i, v := range s.source {
		result[i] = mapper(v)
	}
	return Stream[R]{source: result}
}

// FlatMap returns a stream consisting of the results of replacing each element of this stream with the contents of a mapped stream produced by applying the provided mapping function to each element.
// Note: This is a top-level function because Go methods cannot have type parameters.
func FlatMap[T, R any](s Stream[T], mapper func(T) []R) Stream[R] {
	var result []R
	for _, v := range s.source {
		result = append(result, mapper(v)...)
	}
	return Stream[R]{source: result}
}
