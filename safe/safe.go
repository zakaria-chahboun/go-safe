package safe

// Value safely dereferences a pointer.
// If the pointer is nil, it returns the zero value of type T.
func Value[T any](ptr *T) T {
	var zero T
	return ValueOr(ptr, zero)
}

// ValueOr safely dereferences a pointer.
// If the pointer is nil, it returns the provided default value.
func ValueOr[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}
	return def
}

// Pointer safely returns the pointer itself,
// or a pointer to the zero value of type T.
func Pointer[T any](ptr *T) *T {
	var zero T
	return PointerOr(ptr, zero)
}

// PointerOr safely returns the pointer itself,
// or a pointer to the provided default value if nil.
func PointerOr[T any](ptr *T, def T) *T {
	if ptr != nil {
		return ptr
	}
	return &def
}
