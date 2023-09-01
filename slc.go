package slc

// Includes detects if the given slice includes the given element.
// The type of the slice elements must be comparable.
func Includes[T comparable](slice []T, elem T) bool {
	for i := range slice {
		if slice[i] == elem {
			return true
		}
	}

	return false
}

// IncludesFunc detects if the given slice includes an element satisfying the predicate.
func IncludesFunc[T any](slice []T, predicateFn func(elem T) bool) bool {
	for i := range slice {
		if predicateFn(slice[i]) {
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

// IndexFunc returns the index of an element satisfying the predicate. The function
// returns -1 if such element is not found.
func IndexFunc[T any](slice []T, predicateFn func(elem T) bool) int {
	for i := range slice {
		if predicateFn(slice[i]) {
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

// Map returns a copy of the given slice containing all elements transformed by
// the given function.
func Map[T, S any](slice []T, transformFn func(elem T) S) []S {
	var res []S
	for i := range slice {
		res = append(res, transformFn(slice[i]))
	}
	return res
}

// Reduce returns the result of applying the given function to each element of
// the given slice. The function is applied left-to-right.
func Reduce[T, S any](slice []T, accumulateFn func(acc S, elem T) S) S {
	var res S
	for i := range slice {
		res = accumulateFn(res, slice[i])
	}
	return res
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

// Uniq returns a copy of the given slice containing all unique elements.
// The type of the slice elements must be comparable.
func Uniq[T comparable](slice []T) []T {
	var res []T
	for i := range slice {
		if !Includes(res, slice[i]) {
			res = append(res, slice[i])
		}
	}
	return res
}

// Uniq returns a copy of the given slice containing all unique elements.
// The type of the slice elements must be comparable.
func UniqFunc[T any](slice []T, equalsFn func(n, m T) bool) []T {
	var res []T
	for i := range slice {
		if !IncludesFunc(res, func(elem T) bool { return equalsFn(slice[i], elem) }) {
			res = append(res, slice[i])
		}
	}
	return res
}

// Overlap returns true if two slices have at least one common element.
// The type of the slice elements must be comparable.
func Overlap[T comparable](a, b []T) bool {
	for i := range a {
		if Includes(b, a[i]) {
			return true
		}
	}

	return false
}

// OverlapFunc returns true if two slices have at least one common element.
// The elements are compared using the given function, equalsFn.
func OverlapFunc[T, S any](a []T, b []S, equalsFn func(e1 T, e2 S) bool) bool {
	for i := range a {
		for j := range b {
			if equalsFn(a[i], b[j]) {
				return true
			}
		}
	}

	return false
}

// Intersect returns the intersection of two slices.
// The type of the slice elements must be comparable.
func Intersect[T comparable](a, b []T) []T {
	var res []T
	for i := range a {
		if !Includes(res, a[i]) && Includes(b, a[i]) {
			res = append(res, a[i])
		}
	}
	return res
}

// Intersect returns the intersection of two slices.
// The elements are compared using the given function, equalsFn.
func IntersectFunc[T comparable, S any](a []T, b []S, equalsFn func(e1 T, e2 S) bool) []T {
	var res []T

OuterLoop:
	for i := range a {
		if Includes(res, a[i]) {
			continue
		}

		for j := range b {
			if equalsFn(a[i], b[j]) {
				res = append(res, a[i])
				continue OuterLoop
			}
		}
	}

	return res
}

// Diff returns the difference between the given slices: a - b.
// The type of the slice elements must be comparable.
func Diff[T comparable](a, b []T) []T {
	var res []T
	for i := range a {
		if !Includes(b, a[i]) {
			res = append(res, a[i])
		}
	}
	return res
}

// DiffFunc returns the difference between the given slices: a - b.
// The elements are compared using the given function, equalsFn.
func DiffFunc[T, S any](a []T, b []S, equalsFn func(e1 T, e2 S) bool) []T {
	var res []T
	for i := range a {
		if !IncludesFunc(b, func(elem S) bool { return equalsFn(a[i], elem) }) {
			res = append(res, a[i])
		}
	}
	return res
}
