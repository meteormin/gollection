package gollection

import (
	"github.com/meteormin/gollection/pkg/maps"
	"github.com/meteormin/gollection/pkg/slice"
)

// Collection interface
type Collection[T interface{}] interface {
	// Add adds an item to the collection
	Add(item T)

	// All returns all items
	All() []T

	// Chunk items in collection
	Chunk(chunkSize int, fn func(v []T)) [][]T

	// Count returns the number of items
	Count() int

	// Compact returns a new collection without empty items
	Compact(func(T, T) bool) Collection[T]

	// Compare items in collection
	Compare(items []T, fn func(v1, v2 T) int) int

	// Concat items in collection
	Concat(items ...T)

	// Copy returns a copy of the collection
	Copy() Collection[T]

	// Dequeue removes and returns the first item from the collection
	Dequeue() (*T, error)

	// Each applies a function to each item in the collection
	Each(fn func(v T))

	// Enqueue adds an item to the collection
	Enqueue(item T)

	// Equal checks if two collections are equal
	Equal(other Collection[T], fn func(v1, v2 T) bool) bool

	// Except returns a new collection with items that do not satisfy the given predicate function
	Except(fn func(T) bool) Collection[T]

	// First returns the first item from the collection
	First() (*T, error)

	// Filter returns a new collection with items that satisfy the given predicate function
	Filter(fn func(v T) bool) Collection[T]

	// Find returns the first item that satisfies the given predicate function
	Insert(index int, item T) Collection[T]

	// IsEmpty returns true if the collection is empty, false otherwise
	IsEmpty() bool

	// Items returns the items in the collection
	Items() []T

	// Last returns the last item from the collection
	IsSorted(less func(T, T) int) bool

	// Last returns the last item from the collection
	Last() (*T, error)

	// Map applies a function to each item in the collection and returns a new collection with the results
	Map(fn func(v T) T) Collection[T]

	// Max returns the maximum element in the collection based on the comparison function
	Max(less func(T, T) int) T

	// Merge merges two collections
	Merge(merge []T) Collection[T]

	// Min returns the minimum element in the collection based on the comparison function
	Min(less func(T, T) int) T

	// Pop removes and returns the last item from the collection
	Pop() (*T, error)

	// Push adds an item to the collection
	Push(item T)

	// Remove removes an item from the collection
	Remove(index int) error

	// Reverse returns a new collection with the items in reverse order
	Reverse() Collection[T]

	// Slice returns a new collection with the items in the specified range
	Slice(start, end int) Collection[T]

	// Sort sorts the collection using the provided comparison function
	Sort(less func(T, T) int) Collection[T]
}

type BaseCollection[T interface{}] struct {
	items []T
}

// Add item
func (b *BaseCollection[T]) Add(item T) {
	b.items = slice.Add(b.items, item)
}

// All items
func (b *BaseCollection[T]) All() []T {
	return b.items
}

// Chunk items in collection
func (b *BaseCollection[T]) Chunk(chunkSize int, fn func(v []T)) [][]T {
	return slice.Chunk(b.All(), chunkSize, fn)
}

func (b *BaseCollection[T]) Compact(fn func(T, T) bool) Collection[T] {
	items := slice.CompactFunc(b.All(), fn)
	return NewCollection(items)
}

func (b *BaseCollection[T]) Compare(items []T, fn func(v1, v2 T) int) int {
	return slice.CompareFunc(b.All(), items, fn)
}

// Concat items in collection
func (b *BaseCollection[T]) Concat(items ...T) {
	b.items = slice.Concat(b.All(), items)
}

// Copy returns a new Collection containing all the elements of the BaseCollection.
func (b *BaseCollection[T]) Copy() Collection[T] {
	return NewCollection(b.All())
}

// Count get items count
func (b *BaseCollection[T]) Count() int {
	return len(b.items)
}

// Dequeue removes and returns the first item from the collection.
func (b *BaseCollection[T]) Dequeue() (*T, error) {
	if b.IsEmpty() {
		return nil, ErrIsEmpty
	}

	items, deqItem := slice.Dequeue(b.items)
	b.items = items

	return &deqItem, nil
}

// Each iterates over each element in the BaseCollection and applies the provided function `fn` to each element.
func (b *BaseCollection[T]) Each(fn func(v T)) {
	slice.Each(b.All(), fn)
}

// Enqueue adds an item to the collection.
func (b *BaseCollection[T]) Enqueue(item T) {
	b.items = slice.Enqueue(b.items, item)
}

// Equal compares two collections and returns true if they are equal.
func (b *BaseCollection[T]) Equal(other Collection[T], fn func(T, T) bool) bool {
	return slice.EqualFunc(b.All(), other.All(), fn)
}

// Except removes elements from the collection that satisfy the given predicate function.
func (b *BaseCollection[T]) Except(fn func(v T) bool) Collection[T] {
	items := slice.Except(b.All(), fn)
	return NewCollection(items)
}

// Filter items in collection
func (b *BaseCollection[T]) Filter(fn func(v T) bool) Collection[T] {
	filtered := slice.Filter(b.All(), fn)
	return NewCollection(filtered)
}

// First returns the first element in the collection.
func (b *BaseCollection[T]) First() (*T, error) {
	if b.IsEmpty() {
		return nil, ErrIsEmpty
	}

	first := slice.First(b.All())
	return &first, nil
}

// Insert inserts a value at the specified index in the BaseCollection and returns the modified Collection.
func (b *BaseCollection[T]) Insert(i int, v T) Collection[T] {
	b.items = slice.Insert(b.items, i, v)
	return b
}

// IsEmpty check current collection is empty
func (b *BaseCollection[T]) IsEmpty() bool {
	return b.Count() == 0
}

// IsSorted checks if the collection is sorted.
func (b *BaseCollection[T]) IsSorted(fn func(T, T) int) bool {
	return slice.IsSortedFunc(b.All(), fn)
}

// Items get items
func (b *BaseCollection[T]) Items() []T {
	return b.items
}

// Last returns the last element of the collection.
func (b *BaseCollection[T]) Last() (*T, error) {
	if b.IsEmpty() {
		return nil, ErrIsEmpty
	}

	last := slice.Last(b.All())
	return &last, nil
}

// Map items in collection
func (b *BaseCollection[T]) Map(fn func(v T) T) Collection[T] {
	items := slice.Map(b.All(), fn)
	return NewCollection(items)
}

// Max returns the maximum element in the collection based on the comparison function.
func (b *BaseCollection[T]) Max(fn func(T, T) int) T {
	return slice.MaxFunc(b.All(), fn)
}

// Merge merges the given slice into the collection and returns a new Collection.
func (b *BaseCollection[T]) Merge(merge []T) Collection[T] {
	return NewCollection(slice.Merge(b.All(), merge))
}

// Min returns the minimum element in the collection based on the comparison function.
func (b *BaseCollection[T]) Min(fn func(T, T) int) T {
	return slice.MinFunc(b.All(), fn)
}

// Pop removes and returns the last item from the collection.
func (b *BaseCollection[T]) Pop() (*T, error) {
	if b.IsEmpty() {
		return nil, ErrIsEmpty
	}

	items, popItem := slice.Pop(b.items)
	b.items = items

	return &popItem, nil
}

// Push adds an item to the collection.
func (b *BaseCollection[T]) Push(item T) {
	b.items = slice.Push(b.items, item)
}

func (b *BaseCollection[T]) Remove(i int) error {
	if b.IsEmpty() {
		return ErrIsEmpty
	}

	b.items = slice.Remove(b.items, i)

	return nil
}

// Reverse returns a new Collection with the elements in reverse order.
func (b *BaseCollection[T]) Reverse() Collection[T] {
	return NewCollection(slice.Reverse(b.All()))
}

// Slice returns a new Collection containing the elements from the start index to the end index (exclusive).
func (b *BaseCollection[T]) Slice(start, end int) Collection[T] {
	return NewCollection(slice.Slice(b.All(), start, end))
}

// Sort sorts the collection using the provided comparison function.
func (b *BaseCollection[T]) Sort(fun func(T, T) int) Collection[T] {
	items := b.All()
	slice.SortFunc(items, fun)
	return NewCollection(items)
}

func NewCollection[T any](items []T) Collection[T] {
	return &BaseCollection[T]{
		items: items,
	}
}

// CollectionMap interface
type CollectionMap[K comparable, V interface{}] interface {
	All() map[K]V
	Copy() CollectionMap[K, V]
	Count() int
	Each(fn func(value V))
	Filter(fn func(value V) bool) CollectionMap[K, V]
	Get(key K) V
	IsEmpty() bool
	Items() map[K]V
	Map(fn func(value V) V) CollectionMap[K, V]
	Merge(merge map[K]V) CollectionMap[K, V]
	Put(key K, item V)
	Remove(key K) error
}

type BaseCollectionMap[k comparable, v interface{}] struct {
	items map[k]v
}

func (b *BaseCollectionMap[k, v]) All() map[k]v {
	return b.items
}

func (b *BaseCollectionMap[k, v]) Count() int {
	return len(b.items)
}

func (b *BaseCollectionMap[k, v]) Copy() CollectionMap[k, v] {
	return NewCollectionMap(b.All())
}

func (b *BaseCollectionMap[k, v]) Each(fn func(value v)) {
	for _, value := range b.items {
		fn(value)
	}
}

func (b *BaseCollectionMap[k, v]) Filter(fn func(value v) bool) CollectionMap[k, v] {
	filtered := make(map[k]v)
	for key, value := range b.items {
		if fn(value) {
			filtered[key] = value
		}
	}
	return NewCollectionMap(filtered)
}

func (b *BaseCollectionMap[k, v]) Get(key k) v {
	return b.items[key]
}

func (b *BaseCollectionMap[k, v]) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *BaseCollectionMap[k, v]) Items() map[k]v {
	return b.items
}

func (b *BaseCollectionMap[k, v]) Map(fn func(value v) v) CollectionMap[k, v] {
	mapped := make(map[k]v)
	for key, value := range b.items {
		mapped[key] = fn(value)
	}
	return NewCollectionMap(mapped)
}

func (b *BaseCollectionMap[k, v]) Merge(merge map[k]v) CollectionMap[k, v] {
	newItems := maps.Merge(b.items, merge)
	return NewCollectionMap(newItems)
}

func (b *BaseCollectionMap[k, v]) Put(key k, item v) {
	b.items[key] = item
}

func (b *BaseCollectionMap[k, v]) Remove(key k) error {
	if b.IsEmpty() {
		return ErrIsEmpty
	}

	delete(b.items, key)

	return nil
}

// NewCollectionMap creates a new CollectionMap with the provided key-value pairs.
//
// Parameters:
// - items: The initial key-value pairs to populate the CollectionMap.
// Returns a CollectionMap with the specified key-value pairs.
func NewCollectionMap[k comparable, v interface{}](items map[k]v) CollectionMap[k, v] {
	return &BaseCollectionMap[k, v]{
		items: items,
	}
}
