package stream

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
	callback interface{}
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
	b.callback = fn
	return chainStream[T, T](b)
}

// Filter implements Stream.
func (b *BaseStream[T]) Filter(fn func(item T) bool) Stream[T] {
	b.callback = fn
}

// First implements Stream.
func (b *BaseStream[T]) First() (*T, error) {
	panic("unimplemented")
}

// FlatMap implements Stream.
func (b *BaseStream[T]) FlatMap(fn func(item T) []interface{}) Stream[interface{}] {
	panic("unimplemented")
}

// Get implements Stream.
func (b *BaseStream[T]) Get() ([]T, error) {
	panic("unimplemented")
}

// Last implements Stream.
func (b *BaseStream[T]) Last() (*T, error) {
	panic("unimplemented")
}

// Map implements Stream.
func (b *BaseStream[T]) Map(fn func(item T) interface{}) Stream[interface{}] {
	panic("unimplemented")
}

// Sort implements Stream.
func (b *BaseStream[T]) Sort(fn func(i int, j int) bool) Stream[T] {
	panic("unimplemented")
}
