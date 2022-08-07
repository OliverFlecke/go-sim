package utils

// Removes the element at index i from the slice.
// Order is not reserved.
func Remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
