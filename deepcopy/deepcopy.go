package deepcopy

// Map 返回一个深拷贝的map
func Map[K comparable, V any](target map[K]V) map[K]V {
	return new(builder[map[K]V]).build(target)
}

// Slice 返回一个深拷贝的slice
func Slice[T any](target []T) []T {
	return new(builder[[]T]).build(target)
}

// Struct 返回一个深拷贝的结构体
func Struct[T any](target T) T {
	return new(builder[T]).build(target)
}
