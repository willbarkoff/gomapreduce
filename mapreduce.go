package gomapreduce

// Equal takes two sices comparable with `==`, and returns true if they are equal when compared with ==.
func Equal[T comparable](a, b []T) bool {
	return EqualCmp(a, b, func(a, b T) bool { return a == b })
}

// EqualCmp takes two sices and a comparison function, and returns true if they are equal when each element is compared with the comparison function
func EqualCmp[T any](a, b []T, cmp func(a, b T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !cmp(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Map takes every value in a slice, applies `f` to it, and returns a new slice with the results.
func Map[A any, B any](f func(e A) B, arr []A) []B {
	return Mapi(func(e A, _ int) B { return f(e) }, arr)
}

// Map takes every value in a slice, applies `f` to it with the first argument as the value, and the second argument as the position in the slice, and returns a new slice with the results.
func Mapi[A any, B any](f func(e A, i int) B, arr []A) []B {
	result := []B{}
	for i, val := range arr {
		result = append(result, f(val, i))
	}
	return result
}
