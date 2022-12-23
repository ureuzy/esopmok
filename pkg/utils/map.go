package utils

func Map[T, E any](t []T, f func(t T) E) []E {
	result := make([]E, len(t))
	for i, n := range t {
		result[i] = f(n)
	}
	return result
}

func MapP[T, E any](t *[]T, f func(t T) E) *[]E {
	result := make([]E, len(*t))
	for i, n := range *t {
		result[i] = f(n)
	}
	return &result
}