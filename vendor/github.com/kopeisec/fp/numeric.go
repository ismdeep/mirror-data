package fp

import (
	"cmp"
	"sort"
)

// Number 数字类型约束，覆盖常见的整数与浮点
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum 求和
func Sum[T Number](s Slice[T]) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// Min 返回最小值；切片为空时第二个返回值为 false
func Min[T cmp.Ordered](s Slice[T]) (T, bool) {
	var zero T
	if len(s) == 0 {
		return zero, false
	}
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m, true
}

// Max 返回最大值；切片为空时第二个返回值为 false
func Max[T cmp.Ordered](s Slice[T]) (T, bool) {
	var zero T
	if len(s) == 0 {
		return zero, false
	}
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}

// Avg 返回算术平均值（float64）；切片为空时第二个返回值为 false
func Avg[T Number](s Slice[T]) (float64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	var total float64
	for _, v := range s {
		total += float64(v)
	}
	return total / float64(len(s)), true
}

// Sort 返回升序排序的新切片，原切片不变
func Sort[T cmp.Ordered](s Slice[T]) Slice[T] {
	result := make(Slice[T], len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// SortBy 按 keyFunc 返回的 key 升序排序，返回新切片，稳定排序
func SortBy[T any, K cmp.Ordered](s Slice[T], keyFunc func(T) K) Slice[T] {
	result := make(Slice[T], len(s))
	copy(result, s)
	sort.SliceStable(result, func(i, j int) bool {
		return keyFunc(result[i]) < keyFunc(result[j])
	})
	return result
}
