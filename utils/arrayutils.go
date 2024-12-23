package utils

func Map[T, V any](slc []T, fn func(T) V) []V {
	res := make([]V, len(slc))
	for i, el := range slc {
		res[i] = fn(el)
	}
	return res
}

func All[T any](slc []T, fn func(T) bool) bool {
	isAll := true
	for _, el := range slc {
		isAll = isAll && fn(el)
		if !isAll {
			return isAll
		}
	}
	return isAll
}

func Any[T any](slc []T, fn func(T) bool) bool {
	isAny := false
	for _, el := range slc {
		isAny = isAny || fn(el)
		if isAny {
			return isAny
		}
	}
	return isAny
}

func Find[T any](slc []T, fn func(T) bool) (int, T) {
	for idx, el := range slc {
		if fn(el) {
			return idx, el
		}
	}
	var def T
	return -1, def
}

func Filter[T any](slc []T, fn func(T) bool) []T {
	var filtered []T
	for _, el := range slc {
		if fn(el) {
			filtered = append(filtered, el)
		}
	}
	return filtered
}
