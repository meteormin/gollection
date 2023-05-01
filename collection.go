package gollection

import (
	"errors"
	"fmt"
	"github.com/miniyus/gollection/pkg/maps"
	"github.com/miniyus/gollection/pkg/slice"
)

// Collection interface
type Collection[T interface{}] interface {
	Items() []T
	All() []T
	Get(key int) T
	Copy() Collection[T]
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
	Slice(start int, end int) Collection[T]
	Reverse() Collection[T]
}

// BaseCollection base collection struct
// implements Collection interface
type BaseCollection[T interface{}] struct {
	items []T
}

// NewCollection create collection struct implements Collection interface
func NewCollection[T interface{}](items []T) Collection[T] {
	copyItems := make([]T, len(items))
	copy(copyItems, items)
	return &BaseCollection[T]{
		items: copyItems,
	}
}

// Items get items
func (b *BaseCollection[T]) Items() []T {
	return b.items
}

// All get All items
func (b *BaseCollection[T]) All() []T {
	return slice.Copy(b.items)
}

func (b *BaseCollection[T]) Get(key int) T {
	return b.items[key]
}

func (b *BaseCollection[T]) Copy() Collection[T] {
	return NewCollection(b.All())
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
	items := slice.Map(b.All(), fn)
	return NewCollection(items)
}

// Filter items in collection
func (b *BaseCollection[T]) Filter(fn func(v T, i int) bool) Collection[T] {
	filtered := slice.Filter(b.All(), fn)
	return NewCollection(filtered)
}

// Except items in collection
func (b *BaseCollection[T]) Except(fn func(v T, i int) bool) Collection[T] {
	excepts := slice.Except(b.All(), fn)
	return NewCollection(excepts)
}

// Chunk items in collection
func (b *BaseCollection[T]) Chunk(chunkSize int, fn func(v []T, i int)) [][]T {
	return slice.Chunk(b.All(), chunkSize, fn)
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
	b.items = slice.Push(b.items, item)
}

func (b *BaseCollection[T]) Pop() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	items, popItem := slice.Pop(b.items)
	b.items = items

	return &popItem, nil
}

func (b *BaseCollection[T]) First() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	first := slice.First(b.All())
	return &first, nil
}

func (b *BaseCollection[T]) Last() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	last := slice.Last(b.All())
	return &last, nil
}

func (b *BaseCollection[T]) Merge(merge []T) Collection[T] {
	return NewCollection(slice.Merge(b.All(), merge))
}

func (b *BaseCollection[T]) Slice(start int, end int) Collection[T] {
	return NewCollection(slice.Slice(b.All(), start, end))
}

func (b *BaseCollection[T]) Reverse() Collection[T] {
	return NewCollection(slice.Reverse(b.All()))
}

// CollectionMap interface
type CollectionMap[k comparable, v interface{}] interface {
	Items() map[k]v
	All() map[k]v
	Get(key k) v
	Copy() CollectionMap[k, v]
	Count() int
	IsEmpty() bool
	Put(key k, item v)
	Map(fn func(v v, k k) v) CollectionMap[k, v]
	Filter(fn func(v v, k k) bool) CollectionMap[k, v]
	Except(fn func(v v, k k) bool) CollectionMap[k, v]
	For(fn func(v v, k k))
	Remove(key k) error
	Merge(merge map[k]v) CollectionMap[k, v]
}

type BaseCollectionMap[k comparable, v interface{}] struct {
	items map[k]v
}

func NewCollectionMap[k comparable, v interface{}](items map[k]v) CollectionMap[k, v] {
	return BaseCollectionMap[k, v]{
		items: items,
	}
}

func (b BaseCollectionMap[k, v]) Items() map[k]v {
	return b.items
}

func (b BaseCollectionMap[k, v]) All() map[k]v {
	return maps.Copy(b.items)
}

func (b BaseCollectionMap[k, v]) Get(key k) v {
	return b.items[key]
}

func (b BaseCollectionMap[k, v]) Copy() CollectionMap[k, v] {
	return NewCollectionMap(b.All())
}

func (b BaseCollectionMap[k, v]) Count() int {
	return len(b.items)
}

func (b BaseCollectionMap[k, v]) IsEmpty() bool {
	return b.Count() == 0
}

func (b BaseCollectionMap[k, v]) Put(key k, item v) {
	b.items[key] = item
}

func (b BaseCollectionMap[k, v]) Map(fn func(v v, k k) v) CollectionMap[k, v] {
	mapped := maps.Map(b.All(), fn)
	return NewCollectionMap(mapped)
}

func (b BaseCollectionMap[k, v]) Filter(fn func(v v, k k) bool) CollectionMap[k, v] {
	filtered := maps.Filter(b.All(), fn)
	return NewCollectionMap(filtered)
}

func (b BaseCollectionMap[k, v]) Except(fn func(v v, k k) bool) CollectionMap[k, v] {
	excepted := maps.Except(b.All(), fn)
	return NewCollectionMap(excepted)
}

func (b BaseCollectionMap[k, v]) For(fn func(v v, k k)) {
	maps.For(b.items, fn)
}

func (b BaseCollectionMap[k, v]) Remove(key k) error {
	if _, ok := b.items[key]; ok {
		delete(b.items, key)
		return nil
	}

	return errors.New(fmt.Sprintf("this map has not key: %v", key))
}

func (b BaseCollectionMap[k, v]) Merge(merge map[k]v) CollectionMap[k, v] {
	return NewCollectionMap(maps.Merge(b.All(), merge))
}
