package slice

func First[T any](slice []T, def ...T) T {
	if len(slice) > 0 {
		return slice[0]
	}
	if len(def) > 0 {
		return def[0]
	}
	var zero T
	return zero
}

func Last[T any](slice []T, def ...T) T {
	if len(slice) > 0 {
		return slice[len(slice)-1]
	}
	if len(def) > 0 {
		return def[0]
	}
	var zero T
	return zero
}

func Unique[T comparable](slice []T) []T {
	ret := make([]T, 0)
	dict := make(map[T]struct{})
	for _, s := range slice {
		if _, ok := dict[s]; ok == false {
			dict[s] = struct{}{}
			ret = append(ret, s)
		}
	}
	return ret
}

func Map[T any, R any](slice []T, iter func(item T) R) []R {
	ret := make([]R, len(slice))
	for idx, item := range slice {
		ret[idx] = iter(item)
	}
	return ret
}

func Filter[T any](slice []T, iter func(index int, item T) bool) []T {
	ret := make([]T, 0)
	for idx, item := range slice {
		if iter(idx, item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func FilterZero[T comparable](slice []T) []T {
	ret := make([]T, 0)
	for _, item := range slice {
		var zero T
		if item != zero {
			ret = append(ret, item)
		}
	}
	return ret
}
