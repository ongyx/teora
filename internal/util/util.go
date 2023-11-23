package util

// Must unwraps a two-tuple of (T, error) into T.
// If the error is non-nil, a panic occurs.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
