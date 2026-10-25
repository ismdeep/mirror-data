package fp

// Pair 表示两个元素的元组
type Pair[A any, B any] struct {
	First  A
	Second B
}

// FlatMap 将每个元素映射成切片后拼接成一个切片
func FlatMap[T any, U any](s Slice[T], f func(T) Slice[U]) Slice[U] {
	var result Slice[U]
	for _, v := range s {
		result = append(result, f(v)...)
	}
	return result
}

// Flatten 将切片的切片打平成单层切片
func Flatten[T any](s Slice[Slice[T]]) Slice[T] {
	var result Slice[T]
	for _, sub := range s {
		result = append(result, sub...)
	}
	return result
}

// GroupBy 按 keyFunc 返回的 key 对元素分组
func GroupBy[T any, K comparable](s Slice[T], keyFunc func(T) K) map[K]Slice[T] {
	m := make(map[K]Slice[T])
	for _, v := range s {
		k := keyFunc(v)
		m[k] = append(m[k], v)
	}
	return m
}

// Zip 按位置将两个切片对应元素配对，长度以较短者为准
func Zip[A any, B any](as Slice[A], bs Slice[B]) Slice[Pair[A, B]] {
	n := len(as)
	if len(bs) < n {
		n = len(bs)
	}
	result := make(Slice[Pair[A, B]], n)
	for i := 0; i < n; i++ {
		result[i] = Pair[A, B]{First: as[i], Second: bs[i]}
	}
	return result
}

// Unzip 将切片中的 Pair 拆分为两个切片
func Unzip[A any, B any](pairs Slice[Pair[A, B]]) (Slice[A], Slice[B]) {
	as := make(Slice[A], len(pairs))
	bs := make(Slice[B], len(pairs))
	for i, p := range pairs {
		as[i] = p.First
		bs[i] = p.Second
	}
	return as, bs
}

// ReduceTo 将切片归约成另一种类型的值
func ReduceTo[T any, U any](s Slice[T], accumulator func(U, T) U, initial U) U {
	result := initial
	for _, v := range s {
		result = accumulator(result, v)
	}
	return result
}
