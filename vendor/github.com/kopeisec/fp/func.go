package fp

// Identity 返回参数本身，常用于占位
func Identity[T any](v T) T {
	return v
}

// Not 返回一个对原谓词取反的新谓词
func Not[T any](p func(T) bool) func(T) bool {
	return func(v T) bool {
		return !p(v)
	}
}

// Compose 数学风格的函数复合：Compose(f, g)(x) = f(g(x))
func Compose[A any, B any, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}

// Pipe 同类型链式应用：Pipe(f1, f2, f3)(x) = f3(f2(f1(x)))
func Pipe[T any](fns ...func(T) T) func(T) T {
	return func(v T) T {
		for _, fn := range fns {
			v = fn(v)
		}
		return v
	}
}
