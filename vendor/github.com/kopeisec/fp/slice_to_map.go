package fp

func SliceToMap[D any, K comparable, V any](data []D, f func(D) (K, V)) map[K]V {
	m := make(map[K]V)
	for _, item := range data {
		k, v := f(item)
		m[k] = v
	}
	return m
}
