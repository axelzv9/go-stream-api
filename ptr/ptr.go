package ptr

func Ref[T any](v T) *T {
	return &v
}

func DefaultIfNil[T any](ptr *T, defaultValue T) T {
	if ptr == nil {
		return defaultValue
	}
	return *ptr
}

func DefaultIfZero[T comparable](v T, defaultValue T) T {
	var x T
	if x == v {
		return defaultValue
	}
	return v
}
