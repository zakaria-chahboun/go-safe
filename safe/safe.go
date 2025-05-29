package safe

// Value safely dereferences a pointer.
// If the pointer is nil, it returns the provided default (if any),
// or the zero value of type T.
func Value[T any](ptr *T, defaults ...T) T {
	if ptr != nil {
		return *ptr
	}
	if len(defaults) > 0 {
		return defaults[0]
	}
	var zero T
	return zero
}

// Pointer safely returns the pointer itself,
// or a pointer to the provided default (if any),
// or a pointer to the zero value of type T.
func Pointer[T any](ptr *T, defaults ...T) *T {
	if ptr != nil {
		return ptr // ← reuse same pointer!
	}
	val := Value(ptr, defaults...)
	return &val // new pointer only if original was nil
}
