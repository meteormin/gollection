package iterator

import "errors"

type Iterator[T interface{}] interface {
	Next() (*T, error)
	HasNext() bool
	GetNext() (*T, error)
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
	if i.index < len(i.values) {
		return true
	}
	return false
}

func (i *StructIterator[T]) GetNext() (*T, error) {
	if i.HasNext() {
		return &i.values[i.index], nil
	}

	return nil, errors.New("has not next")
}

func NewIterator[T interface{}](values []T) Iterator[T] {
	return &StructIterator[T]{
		index:  0,
		values: values,
	}
}

type AsyncIterator[T interface{}] interface {
	Next() chan T
	HasNext() bool
	GetNext() chan T
	Quit()
	GetQuitChan() chan bool
}

type StructAsyncIterator[T interface{}] struct {
	index  int
	values []T
	ch     chan T
	qch    chan bool
}

func (a *StructAsyncIterator[T]) HasNext() bool {
	return len(a.values) > a.index
}

func (a *StructAsyncIterator[T]) Next() chan T {
	if !a.HasNext() {
		return nil
	}

	a.ch <- a.values[a.index]
	a.index++

	return a.ch
}

func (a *StructAsyncIterator[T]) GetNext() chan T {
	return a.ch
}

func (a *StructAsyncIterator[T]) Quit() {
	a.qch <- true
}

func (a *StructAsyncIterator[T]) GetQuitChan() chan bool {
	return a.qch
}

func NewAsyncIterator[T interface{}](values []T) AsyncIterator[T] {
	return &StructAsyncIterator[T]{
		index:  0,
		values: values,
		ch:     make(chan T, len(values)),
		qch:    make(chan bool),
	}
}
