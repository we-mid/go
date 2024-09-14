package util

func ZeroValue[T any]() T {
	// fix: compiler: cannot use nil as T value in return statement
	// return nil
	// fix: compiler: invalid composite literal type T
	// return T{}
	var zero T
	return zero
}
