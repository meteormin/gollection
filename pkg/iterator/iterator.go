package iterator

import "errors"

type Iterator[T interface{}] interface {
	Next() (*T, error)
	HasNext() bool
	GetNext() (*T, error)
	GetIndex() int
}

type StructIterator[T interface{}] struct {
	index  int
	values []T
}

func (i *StructIterator[T]) Next() (*T, error) {
	if i.HasNext() {
		next := i.values[i.index]
		i.index++
		return &next, nil
	}
	return nil, errors.New("has not next")
}

func (i *StructIterator[T]) HasNext() bool {
	return i.index < len(i.values)
}

func (i *StructIterator[T]) GetNext() (*T, error) {
	if i.HasNext() {
		return &i.values[i.index], nil
	}

	return nil, errors.New("has not next")
}

func (i *StructIterator[T]) GetIndex() int {
	return i.index
}

func NewIterator[T interface{}](values []T) Iterator[T] {
	return &StructIterator[T]{
		index:  0,
		values: values,
	}
}
