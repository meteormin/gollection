package gollection

import (
	"errors"
	"github.com/miniyus/gollection/pkg/slice"
)

// Collection interface
type Collection[T interface{}] interface {
	Items() []T
	All() []T
	Count() int
	IsEmpty() bool
	Add(item T)
	Map(fn func(v T, i int) T) Collection[T]
	Filter(fn func(v T, i int) bool) Collection[T]
	Except(fn func(v T, i int) bool) Collection[T]
	Chunk(chunkSize int, fn func(v []T, i int)) [][]T
	For(fn func(v T, i int))
	Remove(index int) error
	Concat(items []T)
	Push(item T)
	Pop() (*T, error)
	First() (*T, error)
	Last() (*T, error)
	Merge(merge []T) Collection[T]
}

// BaseCollection base collection struct
// implements Collection interface
type BaseCollection[T interface{}] struct {
	items []T
}

// NewCollection create collection struct implements Collection interface
func NewCollection[T interface{}](items []T) Collection[T] {
	return &BaseCollection[T]{
		items: items,
	}
}

// Items get items
func (b *BaseCollection[T]) Items() []T {
	return b.items
}

// All get All items
func (b *BaseCollection[T]) All() []T {
	all := make([]T, b.Count())
	for _, v := range b.items {
		all = append(all, v)
	}

	return all
}

// Count get items count
func (b *BaseCollection[T]) Count() int {
	return len(b.items)
}

// IsEmpty check current collection is empty
func (b *BaseCollection[T]) IsEmpty() bool {
	return b.Count() == 0
}

// Add item
func (b *BaseCollection[T]) Add(item T) {
	b.items = slice.Add(b.items, item)
}

// Map items in collection
func (b *BaseCollection[T]) Map(fn func(v T, i int) T) Collection[T] {
	items := slice.Map(b.items, fn)
	return NewCollection(items)
}

// Filter items in collection
func (b *BaseCollection[T]) Filter(fn func(v T, i int) bool) Collection[T] {
	filtered := slice.Filter(b.items, fn)
	return NewCollection(filtered)
}

// Except items in collection
func (b *BaseCollection[T]) Except(fn func(v T, i int) bool) Collection[T] {
	excepts := slice.Except(b.items, fn)
	return NewCollection(excepts)
}

// Chunk items in collection
func (b *BaseCollection[T]) Chunk(chunkSize int, fn func(v []T, i int)) [][]T {
	return slice.Chunk(b.items, chunkSize, fn)
}

// For loop items in collection
func (b *BaseCollection[T]) For(fn func(v T, i int)) {
	slice.For(b.items, fn)
}

// Remove item in collection
func (b *BaseCollection[T]) Remove(index int) error {
	if b.IsEmpty() {
		return errors.New("this collection is empty")
	}
	b.items = slice.Remove(b.items, index)

	return nil
}

// Concat items in collection
func (b *BaseCollection[T]) Concat(items []T) {
	b.items = slice.Concat(b.items, items)
}

func (b *BaseCollection[T]) Push(item T) {
	b.Add(item)
}

func (b *BaseCollection[T]) Pop() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	popItem := b.items[len(b.items)-1]
	b.items = b.items[:len(b.items)-1]

	return &popItem, nil
}

func (b *BaseCollection[T]) First() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	return &b.items[0], nil
}

func (b *BaseCollection[T]) Last() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	return &b.items[b.Count()-1], nil
}

func (b *BaseCollection[T]) Merge(merge []T) Collection[T] {
	items := b.items
	items = append(items, merge...)
	return NewCollection(items)
}
