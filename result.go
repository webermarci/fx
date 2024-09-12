package fx

type Result[T any] struct {
	value T
	err   error
}

func (result Result[T]) IsOk() bool {
	return result.err == nil
}

func (result Result[T]) IsErr() bool {
	return result.err != nil
}

func (result Result[T]) Get() T {
	return result.value
}

func (result Result[T]) Error() error {
	return result.err
}

func (result Result[T]) Values() (T, error) {
	return result.value, result.err
}

func (result Result[T]) OrElse(value T) T {
	if result.err == nil {
		return result.value
	}
	return value
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		value: *new(T),
		err:   err,
	}
}
