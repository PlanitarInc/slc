package slc

// Includes detects if the given slice includes the given element.
// The type of the slice elements must be comparable.
func Includes[S ~[]E, E comparable](s S, e E) bool {
	for i := range s {
		if s[i] == e {
			return true
		}
	}

	return false
}

// IncludesFunc detects if the given slice includes an element satisfying the predicate.
func IncludesFunc[S ~[]E, E any](s S, predicateFn func(e E) bool) bool {
	for i := range s {
		if predicateFn(s[i]) {
			return true
		}
	}

	return false
}

// Index returns the index of the given element in the slice. The function
// returns -1 if the element is not found.
// The type of the slice elements must be comparable.
func Index[S ~[]E, E comparable](s S, e E) int {
	for i := range s {
		if s[i] == e {
			return i
		}
	}

	return -1
}

// IndexFunc returns the index of an element satisfying the predicate. The function
// returns -1 if such element is not found.
func IndexFunc[S ~[]E, E any](s S, predicateFn func(e E) bool) int {
	for i := range s {
		if predicateFn(s[i]) {
			return i
		}
	}

	return -1
}

// Every detects if all elements satisfy the given predicate.
func Every[S ~[]E, E any](s S, predicateFn func(e E) bool) bool {
	for i := range s {
		if !predicateFn(s[i]) {
			return false
		}
	}

	return true
}

// Some detects if there is at least one element satisfying the given
// predicate.
func Some[S ~[]E, E any](s S, predicateFn func(e E) bool) bool {
	for i := range s {
		if predicateFn(s[i]) {
			return true
		}
	}

	return false
}

// Find returns the first element in the given slice satisfying the predicate.
// A zero value is returned if no element is found.
func Find[S ~[]E, E any](s S, predicateFn func(e E) bool) E {
	for i := range s {
		if predicateFn(s[i]) {
			return s[i]
		}
	}

	var zero E
	return zero
}

// FindPtr returns a pointer to the first element in the given slice satisfying
// the predicate. Nil is returned if no element is found.
func FindPtr[S ~[]E, E any](s S, predicateFn func(e E) bool) *E {
	for i := range s {
		if predicateFn(s[i]) {
			return &s[i]
		}
	}

	return nil
}

// Map returns a copy of the given slice containing all elements transformed by
// the given function.
func Map[S1 ~[]E1, E1, E2 any](s S1, transformFn func(e E1) E2) []E2 {
	var res []E2
	for i := range s {
		res = append(res, transformFn(s[i]))
	}
	return res
}

// Reduce returns the result of applying the given function to each element of
// the given slice. The function is applied left-to-right.
func Reduce[S ~[]E, E, T any](s S, accumulateFn func(acc T, e E) T) T {
	var res T
	for i := range s {
		res = accumulateFn(res, s[i])
	}
	return res
}

// Filter returns a copy of the given slice containing all elements satisfying
// the given predicate.
//
// This function is the opposite of FilterOut().
func Filter[S ~[]E, E any](s S, predicateFn func(e E) bool) S {
	var res S
	for i := range s {
		if predicateFn(s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}

// FilterOut returns a copy of the given slice containing all elements **not**
// satisfying the given predicate.
//
// This function is the opposite of Filter().
func FilterOut[S ~[]E, E any](s S, predicateFn func(e E) bool) S {
	var res S
	for i := range s {
		if !predicateFn(s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}

// Uniq returns a copy of the given slice containing all unique elements.
// The type of the slice elements must be comparable.
func Uniq[S ~[]E, E comparable](s S) S {
	var res S
	for i := range s {
		if !Includes(res, s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}

// Uniq returns a copy of the given slice containing all unique elements.
// The type of the slice elements must be comparable.
func UniqFunc[S ~[]E, E any](s S, equalsFn func(n, m E) bool) S {
	var res S
	for i := range s {
		if !IncludesFunc(res, func(e E) bool { return equalsFn(s[i], e) }) {
			res = append(res, s[i])
		}
	}
	return res
}

// Overlap returns true if two slices have at least one common element.
// The type of the slice elements must be comparable.
func Overlap[S ~[]E, E comparable](s1, s2 S) bool {
	for i := range s1 {
		if Includes(s2, s1[i]) {
			return true
		}
	}

	return false
}

// OverlapFunc returns true if two slices have at least one common element.
// The elements are compared using the given function, equalsFn.
func OverlapFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](
	s1 S1,
	s2 S2,
	equalsFn func(e1 E1, e2 E2) bool,
) bool {
	for i := range s1 {
		for j := range s2 {
			if equalsFn(s1[i], s2[j]) {
				return true
			}
		}
	}

	return false
}

// Intersect returns the intersection of two slices.
// The type of the slice elements must be comparable.
func Intersect[S ~[]E, E comparable](s1, s2 S) S {
	var res S
	for i := range s1 {
		if !Includes(res, s1[i]) && Includes(s2, s1[i]) {
			res = append(res, s1[i])
		}
	}
	return res
}

// Intersect returns the intersection of two slices.
// The elements are compared using the given function, equalsFn.
func IntersectFunc[S1 ~[]E1, S2 ~[]E2, E1 comparable, E2 any](
	s1 S1,
	s2 S2,
	equalsFn func(e1 E1, e2 E2) bool,
) S1 {
	var res S1

OuterLoop:
	for i := range s1 {
		if Includes(res, s1[i]) {
			continue
		}

		for j := range s2 {
			if equalsFn(s1[i], s2[j]) {
				res = append(res, s1[i])
				continue OuterLoop
			}
		}
	}

	return res
}

// Diff returns the difference between the given slices: s1 - s2.
// The type of the slice elements must be comparable.
func Diff[S ~[]E, E comparable](s1, s2 S) S {
	var res S
	for i := range s1 {
		if !Includes(s2, s1[i]) {
			res = append(res, s1[i])
		}
	}
	return res
}

// DiffFunc returns the difference between the given slices: s1 - s2.
// The elements are compared using the given function, equalsFn.
func DiffFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](
	s1 S1,
	s2 S2,
	equalsFn func(e1 E1, e2 E2) bool,
) S1 {
	var res S1
	for i := range s1 {
		if !IncludesFunc(s2, func(e E2) bool { return equalsFn(s1[i], e) }) {
			res = append(res, s1[i])
		}
	}
	return res
}
