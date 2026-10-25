package fp

// MapKeys 返回 map 的所有 key（顺序不保证）
func MapKeys[K comparable, V any](m map[K]V) Slice[K] {
	keys := make(Slice[K], 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// MapValues 返回 map 的所有 value（顺序不保证）
func MapValues[K comparable, V any](m map[K]V) Slice[V] {
	values := make(Slice[V], 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// MapToSlice 把每个 (key, value) 通过 f 转成一个值，结果汇总成切片（顺序不保证）
func MapToSlice[K comparable, V any, R any](m map[K]V, f func(K, V) R) Slice[R] {
	result := make(Slice[R], 0, len(m))
	for k, v := range m {
		result = append(result, f(k, v))
	}
	return result
}

// MergeMaps 合并多个 map，后面 map 的 key 会覆盖前面的
func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// MapFilter 仅保留满足谓词的条目，返回新 map
func MapFilter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}
