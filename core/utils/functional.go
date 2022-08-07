package utils

func Mapi[T any, Output any](as []T, mapper func(int, T) Output) []Output {
	result := make([]Output, len(as))

	for i, v := range as {
		result[i] = mapper(i, v)
	}

	return result
}

func Filteri[T any](as []T, predicate func(int, T) bool) []T {
	var result []T

	for i, v := range as {
		if predicate(i, v) {
			result = append(result, v)
		}
	}

	return result
}
