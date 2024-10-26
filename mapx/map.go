package mapx

// Keys 返回字典的键列表
func Keys[K comparable, V any](dict map[K]V) []K {
	keys := make([]K, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回字典的值列表
func Values[K comparable, V any](dict map[K]V) []V {
	values := make([]V, 0, len(dict))
	for _, v := range dict {
		values = append(values, v)
	}
	return values
}
