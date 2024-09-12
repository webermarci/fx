package fx

import (
	"bytes"
	"encoding/json"
)

type Option[T any] struct {
	value  T
	exists bool
}

func (option Option[T]) MarshalJSON() ([]byte, error) {
	if option.exists {
		return json.Marshal(option.value)
	}
	return []byte("null"), nil
}

func (option *Option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		option.exists = false
		return nil
	}

	if err := json.Unmarshal(data, &option.value); err != nil {
		return err
	}

	option.exists = true
	return nil
}

func (option Option[T]) IsSome() bool {
	return option.exists
}

func (option Option[T]) Get() T {
	return option.value
}

func (option Option[T]) Value() (T, bool) {
	return option.value, option.exists
}

func (option Option[T]) OrElse(value T) T {
	if option.exists {
		return option.value
	}
	return value
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value:  value,
		exists: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		exists: false,
	}
}
