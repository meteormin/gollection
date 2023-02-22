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
	Next(fn func(v T, i int) error)
	Quit()
}

type StructAsyncIterator[T interface{}] struct {
	index  int
	values []T
	ch     chan T
	qch    chan bool
	errs   []error
}

func (a *StructAsyncIterator[T]) Next(fn func(v T, i int) error) {
	if len(a.values) <= a.index {
		a.qch <- true
	}

	next := a.values[a.index]
	a.ch <- next
	select {
	case n := <-a.ch:
		err := fn(n, a.index)
		if err != nil {
			a.errs[a.index] = err
		}
		a.index++
	case <-a.qch:
		return
	}
}

func (a *StructAsyncIterator[T]) Quit() {
	a.qch <- true
}

func NewAsyncIterator[T interface{}](values []T) AsyncIterator[T] {
	return &StructAsyncIterator[T]{
		index:  0,
		values: values,
		ch:     make(chan T, len(values)),
		qch:    make(chan bool),
		errs:   make([]error, len(values)),
	}
}
