package gollection

import (
	"errors"
	"fmt"
	"github.com/miniyus/gollection/pkg/maps"
	"github.com/miniyus/gollection/pkg/slice"
	"sort"
)

// Collection interface
type Collection[T interface{}] interface {
	// Items returns a slice of type T.
	//
	// No parameters.
	// Returns a slice of type T.
	Items() []T

	// All returns all elements of type T.
	//
	// It does not take any parameters.
	// It returns a slice of type T.
	All() []T

	// Get retrieves the value associated with the given key.
	//
	// Parameters:
	// - key: an integer representing the key to search for.
	//
	// Returns:
	// - T: the value associated with the key, or the zero value of T if the key is not found.
	Get(key int) T

	// Copy returns a new Collection with a copy of all the elements.
	//
	// Return:
	// Collection[T] - a new Collection with a copy of all the elements.
	Copy() Collection[T]

	// Count returns the number of elements in the collection.
	//
	// It does not modify the collection.
	// The return type is int.
	Count() int

	// IsEmpty checks if the value is empty.
	//
	// It returns a boolean value indicating whether the value is empty or not.
	IsEmpty() bool

	// Add adds an item to the collection.
	//
	// item: The item to add.
	Add(item T)

	// Map applies a given function to each element of the collection and returns a new collection with the results.
	//
	// fn: The function to apply to each element of the collection. It takes an element of type T and its index as parameters and returns a value of type T.
	// Returns: A new collection of type Collection[T] with the results of applying the function to each element.
	Map(fn func(v T, i int) T) Collection[T]

	// Filter returns a new collection containing the elements that satisfy the given predicate function.
	//
	// fn - The predicate function that takes an element of type T and its index and returns a boolean value indicating whether the element should be included in the filtered collection.
	// Returns a new collection of type Collection[T] containing the elements that satisfy the predicate function.
	Filter(fn func(v T, i int) bool) Collection[T]

	// Except returns a new collection with all the elements that do not satisfy the provided function.
	//
	// fn: A function that takes an element and its index and returns a boolean value indicating whether the element should be excluded from the new collection.
	// Returns: A new collection containing all the elements that do not satisfy the provided function.
	Except(fn func(v T, i int) bool) Collection[T]

	// Chunk takes an integer chunkSize and a function fn as parameters. It divides a slice into smaller
	// chunks of size chunkSize and calls the function fn on each chunk along with its index. It returns
	// a slice of slices, where each inner slice represents a chunk of the original slice.
	//
	// Parameters:
	// - chunkSize: an integer representing the size of each chunk.
	// - fn: a function that takes a slice of T and an integer index as parameters.
	//
	// Returns:
	// - [][]T: a slice of slices, where each inner slice represents a chunk of the original slice.
	Chunk(chunkSize int, fn func(v []T, i int)) [][]T

	// For applies a function to each element of the collection.
	//
	// fn: The function to apply to each element.
	// v: The element of the collection.
	// i: The index of the element.
	For(fn func(v T, i int))

	// Remove removes an element at the specified index.
	//
	// index: The index of the element to be removed.
	// error: An error if the removal fails.
	Remove(index int) error

	// Concat concatenates the items of type T into a single string.
	//
	// The function takes a variadic parameter `items` of type T, which represents
	// the items that will be concatenated. The items can be of any type as long
	// as it is the same type as T.
	//
	// The function does not return any values.
	Concat(items ...T)

	// Push adds an item to the collection.
	//
	// item: The item to be added.
	Push(item T)

	// Pop returns the next element from the stack and removes it.
	//
	// It returns a pointer to the element and an error if the stack is empty.
	Pop() (*T, error)

	// Enqueue adds an item to the queue.
	//
	// item: the item to be added to the queue.
	Enqueue(item T)

	// Dequeue description of the Go function.
	//
	// None.
	// (*T, error).
	Dequeue() (*T, error)

	// First returns the first element of type T and an error, if any.
	//
	// It does not take any parameters.
	// It returns a pointer to a T and an error.
	First() (*T, error)

	// Last returns the last element of the T type slice and an error, if any.
	//
	// Returns:
	// - *T: The last element of the slice.
	// - error: An error, if any.
	Last() (*T, error)

	// Merge merges the elements of the given slice into the collection.
	//
	// merge - the slice to merge into the collection.
	// Returns a new collection containing the merged elements.
	Merge(merge []T) Collection[T]

	// Slice returns a new Collection[T] that is a slice of the current Collection[T].
	//
	// It takes two parameters, start and end, which specify the range of elements to include
	// in the slice. The start parameter is the index of the first element to include, and the
	// end parameter is the index of the first element to exclude.
	//
	// The function returns a new Collection[T] that contains the elements in the specified range.
	Slice(start, end int) Collection[T]

	// Reverse returns a new collection with the elements in reverse order.
	//
	// No parameters.
	// Returns a Collection[T].
	Reverse() Collection[T]

	// Sort sorts the elements of the collection using the provided less function.
	//
	// The less function should return true if the element at index i should be
	// positioned before the element at index j in the sorted collection.
	//
	// It returns a new sorted collection of the same type.
	Sort(func(i, j int) bool) Collection[T]
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

// All returns a copy of all items in the collection.
//
// It takes no parameters.
// It returns a slice of type T.
func (b *BaseCollection[T]) All() []T {
	return slice.Copy(b.items)
}

// Get retrieves the value associated with the given key from the collection.
//
// Parameters:
// - key: the key used to retrieve the value.
//
// Return type:
// - T: the value associated with the given key.
func (b *BaseCollection[T]) Get(key int) T {
	return b.items[key]
}

// Copy returns a new Collection containing all the elements of the BaseCollection.
//
// No parameters.
// Collection[T]
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
func (b *BaseCollection[T]) Concat(items ...T) {
	b.items = slice.Concat(b.items, items)
}

// Push adds an item to the collection.
//
// item: the item to be added to the collection.
func (b *BaseCollection[T]) Push(item T) {
	b.items = slice.Push(b.items, item)
}

// Pop removes and returns the last item from the collection.
//
// It returns a pointer to the popped item and an error if the collection is empty.
func (b *BaseCollection[T]) Pop() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	items, popItem := slice.Pop(b.items)
	b.items = items

	return &popItem, nil
}

// Enqueue adds an item to the collection.
//
// item: the item to be added.
func (b *BaseCollection[T]) Enqueue(item T) {
	b.items = slice.Enqueue(b.items, item)
}

// Dequeue removes and returns the first item from the collection.
//
// Returns a pointer to the dequeued item and an error if the collection is empty.
func (b *BaseCollection[T]) Dequeue() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	items, deqItem := slice.Dequeue(b.items)
	b.items = items

	return &deqItem, nil
}

// First returns the first element in the collection.
//
// It returns a pointer to the first element and an error if the collection is empty.
func (b *BaseCollection[T]) First() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	first := slice.First(b.All())
	return &first, nil
}

// Last returns the last element of the collection.
//
// It returns a pointer to the last element and an error if the collection is empty.
func (b *BaseCollection[T]) Last() (*T, error) {
	if b.IsEmpty() {
		return nil, errors.New("this collection is empty")
	}

	last := slice.Last(b.All())
	return &last, nil
}

// Merge merges the given slice into the collection and returns a new Collection.
//
// merge is a slice of type T that will be merged into the collection.
// The function returns a Collection of type T.
func (b *BaseCollection[T]) Merge(merge []T) Collection[T] {
	return NewCollection(slice.Merge(b.All(), merge))
}

// Slice returns a new Collection containing the elements from the start index to the end index (exclusive).
//
// Parameters:
// - start: the starting index of the slice.
// - end: the ending index of the slice.
//
// Return type(s):
// - Collection[T]: a new Collection containing the sliced elements.
func (b *BaseCollection[T]) Slice(start, end int) Collection[T] {
	return NewCollection(slice.Slice(b.All(), start, end))
}

// Reverse returns a new Collection with the elements in reverse order.
//
// No parameters.
// Returns a Collection.
func (b *BaseCollection[T]) Reverse() Collection[T] {
	return NewCollection(slice.Reverse(b.All()))
}

// Sort sorts the collection using the provided comparison function.
//
// The comparison function takes two indices as input (i, j) and returns true if the element at index i should be
// placed before the element at index j in the sorted collection.
//
// The function returns a new collection that is sorted according to the provided comparison function.
func (b *BaseCollection[T]) Sort(fun func(i, j int) bool) Collection[T] {
	items := b.All()
	sort.Slice(items, fun)
	return NewCollection(items)
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
