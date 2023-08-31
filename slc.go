package slc

// Includes detects if the given slice includes the given element. The type of
// the slice elements must be comparable.
func Includes[T comparable](slice []T, elem T) bool {
	for i := range slice {
		if slice[i] == elem {
			return true
		}
	}

	return false
}

// Index returns the index of the given element in the slice. The function
// returns -1 if the element is not found.
// The type of the slice elements must be comparable.
func Index[T comparable](slice []T, elem T) int {
	for i := range slice {
		if slice[i] == elem {
			return i
		}
	}

	return -1
}

// Every detects if all elements satisfy the given predicate.
func Every[T any](slice []T, predicateFn func(elem T) bool) bool {
	for i := range slice {
		if !predicateFn(slice[i]) {
			return false
		}
	}

	return true
}

// Some detects if there is at least one element satisfying the given
// predicate.
func Some[T any](slice []T, predicateFn func(elem T) bool) bool {
	for i := range slice {
		if predicateFn(slice[i]) {
			return true
		}
	}

	return false
}

// Find returns the first element in the given slice satisfying the predicate.
// A zero value is returned if no element is found.
func Find[T any](slice []T, predicateFn func(elem T) bool) T {
	for i := range slice {
		if predicateFn(slice[i]) {
			return slice[i]
		}
	}

	var zero T
	return zero
}

// FindPtr returns a pointer to the first element in the given slice satisfying
// the predicate. Nil is returned if no element is found.
func FindPtr[T any](slice []T, predicateFn func(elem T) bool) *T {
	for i := range slice {
		if predicateFn(slice[i]) {
			return &slice[i]
		}
	}

	return nil
}

// Filter returns a copy of the given slice containing all elements satisfying
// the given predicate.
//
// This function is the opposite of FilterOut().
func Filter[T any](slice []T, predicateFn func(elem T) bool) []T {
	var res []T
	for i := range slice {
		if predicateFn(slice[i]) {
			res = append(res, slice[i])
		}
	}
	return res
}

// FilterOut returns a copy of the given slice containing all elements **not**
// satisfying the given predicate.
//
// This function is the opposite of Filter().
func FilterOut[T any](slice []T, predicateFn func(elem T) bool) []T {
	var res []T
	for i := range slice {
		if !predicateFn(slice[i]) {
			res = append(res, slice[i])
		}
	}
	return res
}

// Diff returns the difference between the given slices: a - b. The type of the
// slice elements must be comparable.
func Diff[T comparable](a, b []T) []T {
	var res []T
	for i := range a {
		if !Includes(b, a[i]) {
			res = append(res, a[i])
		}
	}
	return res
}
