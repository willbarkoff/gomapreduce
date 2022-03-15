package gomapreduce

func Equal[T any](a, b []T, cmp func(a, b T) bool) bool {
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

func Map[A any, B any](f func(orig A, i int) B, arr []A) []B {
	result := []B{}
	for i, val := range arr {
		result = append(result, f(val, i))
	}
	return result
}
