package fp

import "sync"

// Slice 是一个泛型切片类型
type Slice[T any] []T

// Filter 函数根据条件过滤切片中的元素
func (s Slice[T]) Filter(condition func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterWithKeyFunc 函数根据 KeyFunc 函数进行过滤
func (s Slice[T]) FilterWithKeyFunc(keyFunc func(T) string) Slice[T] {
	m := make(map[string]bool)
	return s.Filter(func(item T) bool {
		k := keyFunc(item)
		if _, ok := m[k]; ok {
			return false
		}
		m[k] = true
		return true
	})
}

// FilterInSlice 筛选出在 Slice 中的
func (s Slice[T]) FilterInSlice(slice []T, equal func(T, T) bool) Slice[T] {
	return s.Filter(func(item T) bool {
		return len(Wrap(slice).Filter(func(itemInSlice T) bool {
			return equal(itemInSlice, item)
		})) > 0
	})
}

// RemoveDuplicate 方法移除切片中的重复元素
func (s Slice[T]) RemoveDuplicate(equal func(a T, b T) bool) Slice[T] {
	result := Slice[T]{}
	for _, item := range s {
		duplicate := false
		for _, uniqueItem := range result {
			if equal(item, uniqueItem) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, item)
		}
	}
	return result
}

// Map 函数对切片中的每个元素应用变换函数
func (s Slice[T]) Map(transform func(T) T) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

// Reduce 函数将切片归约为一个值
func (s Slice[T]) Reduce(accumulator func(T, T) T, initial T) T {
	result := initial
	for _, v := range s {
		result = accumulator(result, v)
	}
	return result
}

// Transform 函数将切片中的每一个值转成新类型
func Transform[T any, B any](s Slice[T], transform func(T) B) Slice[B] {
	var result Slice[B]
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

// TransformAsync 函数将切片中的每一个值转成新类型
func TransformAsync[T any, B any](s Slice[T], transform func(T) B) Slice[B] {
	result := make(Slice[B], len(s))
	var wg sync.WaitGroup
	for i, v := range s {
		wg.Add(1)
		go func(i int, v T) {
			defer wg.Done()
			result[i] = transform(v)
		}(i, v)
	}
	wg.Wait()
	return result
}

func Wrap[T any](s Slice[T]) Slice[T] {
	return s
}

func (s Slice[T]) Monad(f func(T)) Slice[T] {
	for _, v := range s {
		f(v)
	}
	return s
}

func (s Slice[T]) MonadAsync(f func(T)) Slice[T] {
	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go func(v T) {
			defer wg.Done()
			f(v)
		}(v)
	}
	wg.Wait()
	return s
}

// Len 返回切片长度
func (s Slice[T]) Len() int {
	return len(s)
}

// IsEmpty 判断切片是否为空
func (s Slice[T]) IsEmpty() bool {
	return len(s) == 0
}

// First 返回第一个元素；切片为空时第二个返回值为 false
func (s Slice[T]) First() (T, bool) {
	var zero T
	if len(s) == 0 {
		return zero, false
	}
	return s[0], true
}

// Last 返回最后一个元素；切片为空时第二个返回值为 false
func (s Slice[T]) Last() (T, bool) {
	var zero T
	if len(s) == 0 {
		return zero, false
	}
	return s[len(s)-1], true
}

// Any 任一元素满足谓词即返回 true
func (s Slice[T]) Any(predicate func(T) bool) bool {
	for _, v := range s {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All 所有元素都满足谓词才返回 true（空切片返回 true）
func (s Slice[T]) All(predicate func(T) bool) bool {
	for _, v := range s {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Count 满足谓词的元素个数
func (s Slice[T]) Count(predicate func(T) bool) int {
	n := 0
	for _, v := range s {
		if predicate(v) {
			n++
		}
	}
	return n
}

// Find 返回第一个满足谓词的元素；未找到时第二个返回值为 false
func (s Slice[T]) Find(predicate func(T) bool) (T, bool) {
	var zero T
	for _, v := range s {
		if predicate(v) {
			return v, true
		}
	}
	return zero, false
}

// FindIndex 返回第一个满足谓词的元素的下标；未找到时返回 -1
func (s Slice[T]) FindIndex(predicate func(T) bool) int {
	for i, v := range s {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// Contains 判断切片中是否存在与 item 相等的元素
func (s Slice[T]) Contains(item T, equal func(T, T) bool) bool {
	for _, v := range s {
		if equal(v, item) {
			return true
		}
	}
	return false
}

// IndexOf 返回与 item 相等的元素的下标；未找到时返回 -1
func (s Slice[T]) IndexOf(item T, equal func(T, T) bool) int {
	for i, v := range s {
		if equal(v, item) {
			return i
		}
	}
	return -1
}

// Reverse 返回反转后的新切片，原切片不变
func (s Slice[T]) Reverse() Slice[T] {
	result := make(Slice[T], len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}

// Take 取前 n 个元素；n <= 0 返回空切片，n > len 返回全部
func (s Slice[T]) Take(n int) Slice[T] {
	if n <= 0 {
		return Slice[T]{}
	}
	if n >= len(s) {
		return append(Slice[T]{}, s...)
	}
	return append(Slice[T]{}, s[:n]...)
}

// Drop 丢弃前 n 个元素；n <= 0 返回全部，n > len 返回空切片
func (s Slice[T]) Drop(n int) Slice[T] {
	if n <= 0 {
		return append(Slice[T]{}, s...)
	}
	if n >= len(s) {
		return Slice[T]{}
	}
	return append(Slice[T]{}, s[n:]...)
}

// TakeWhile 从头开始连续保留满足谓词的元素，遇到第一个不满足的元素即停
func (s Slice[T]) TakeWhile(predicate func(T) bool) Slice[T] {
	result := Slice[T]{}
	for _, v := range s {
		if !predicate(v) {
			break
		}
		result = append(result, v)
	}
	return result
}

// DropWhile 从头开始连续丢弃满足谓词的元素，遇到第一个不满足后保留其余全部
func (s Slice[T]) DropWhile(predicate func(T) bool) Slice[T] {
	for i, v := range s {
		if !predicate(v) {
			return append(Slice[T]{}, s[i:]...)
		}
	}
	return Slice[T]{}
}

// Chunk 按 size 切分；size <= 0 时返回空切片
func (s Slice[T]) Chunk(size int) []Slice[T] {
	if size <= 0 || len(s) == 0 {
		return []Slice[T]{}
	}
	var result []Slice[T]
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunk := make(Slice[T], end-i)
		copy(chunk, s[i:end])
		result = append(result, chunk)
	}
	return result
}

// Concat 与另一切片拼接，返回新切片
func (s Slice[T]) Concat(other Slice[T]) Slice[T] {
	result := make(Slice[T], 0, len(s)+len(other))
	result = append(result, s...)
	result = append(result, other...)
	return result
}

// Partition 按谓词将切片切成两部分：第一个返回值满足谓词，第二个不满足
func (s Slice[T]) Partition(predicate func(T) bool) (Slice[T], Slice[T]) {
	matched := Slice[T]{}
	unmatched := Slice[T]{}
	for _, v := range s {
		if predicate(v) {
			matched = append(matched, v)
		} else {
			unmatched = append(unmatched, v)
		}
	}
	return matched, unmatched
}
