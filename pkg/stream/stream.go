package stream

import (
	"errors"

	"github.com/meteormin/gollection/pkg/slice"
)

var StreamIsEmptyError = errors.New("this stream is empty")

type Stream[T interface{}] interface {
	Get() ([]T, error)
	First() (*T, error)
	Last() (*T, error)
	Filter(fn func(item T) bool) Stream[T]
	Map(fn func(item T) interface{}) Stream[interface{}]
	FlatMap(fn func(item T) []interface{}) Stream[interface{}]
	Each(fn func(item T)) Stream[T]
	Sort(fn func(i, j int) bool) Stream[T]
}

type BaseStream[T interface{}] struct {
	items    []T
	callback func(items []T) []interface{}
	parent   Stream[interface{}]
}

// NewStream creates a new stream of type T using the given items.
// It takes a slice of items and returns a Stream of type T.
func NewStream[T interface{}](items []T) Stream[T] {
	return &BaseStream[T]{
		items:    items,
		callback: nil,
		parent:   nil,
	}
}

func chainStream[T interface{}, E interface{}](s Stream[T]) Stream[E] {
	return &BaseStream[E]{
		items:    nil,
		callback: nil,
		parent:   s.(Stream[interface{}]),
	}
}

// Each implements Stream.
func (b *BaseStream[T]) Each(fn func(item T)) Stream[T] {
	b.callback = func(items []T) []interface{} {
		slice.Each(items, func(item T, _ int) {
			fn(item)
		})

		return b.items
	}

	return chainStream[T, T](b)
}

// Filter implements Stream.
func (b *BaseStream[T]) Filter(fn func(item T) bool) Stream[T] {
	b.callback = func(items []T) []T {
		return slice.Filter(items, func(item T, _ int) bool {
			return fn(item)
		})
	}

	return chainStream[T, T](b)
}

// First implements Stream.
func (b *BaseStream[T]) First() (*T, error) {
	items, err := b.Get()
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, StreamIsEmptyError
	}

	return &items[0], nil
}

// FlatMap implements Stream.
func (b *BaseStream[T]) FlatMap(fn func(item T) []interface{}) Stream[interface{}] {
	b.callback = func(items []T) []T {
		return slice.FlatMap(items, func(item T, _ int) []interface{} {
			return fn(item)
		})
	}

	return chainStream[T, interface{}](b)
}

// Get implements Stream.
func (b *BaseStream[T]) Get() ([]T, error) {
	if b.parent == nil {
		return b.callback(b.items), nil
	} else {
		tmpItems, err := b.parent.Get()
		if err != nil {
			return make([]T, 0), err
		}

		bItems := make([]T, 0)
		for _, i := range tmpItems {
			bItems = append(bItems, i.(T))
		}

		return b.callback(bItems), nil
	}
}

// Last implements Stream.
func (b *BaseStream[T]) Last() (*T, error) {
	items, err := b.Get()
	if err != nil {
		return nil, err
	}

	length := len(items)
	if length == 0 {
		return nil, StreamIsEmptyError
	}

	return &items[length-1], nil
}

// Map implements Stream.
func (b *BaseStream[T]) Map(fn func(item T) interface{}) Stream[interface{}] {
	panic("unimplemented")
}

// Sort implements Stream.
func (b *BaseStream[T]) Sort(fn func(i int, j int) bool) Stream[T] {
	panic("unimplemented")
}
