package slice

import (
	"cmp"
	"math"
	"slices"
)

// Add appends a value to a slice.
//
// It takes a slice `s` and a value `v` as parameters.
// It returns a new slice with the value `v` appended to it.
func Add[T interface{}](s []T, v T) []T {
	return append(s, v)
}

// Concat concatenates two slices of any type.
//
// It takes two slices of type T as input and returns a new slice of type T.
func Concat[T interface{}](s []T, s2 []T) []T {
	if len(s2) == 0 {
		return s
	}

	return slices.Concat(s, s2)
}

// Compact removes any falsey values from the input slice.
//
// Parameters:
// - s: The input slice of type T.
//
// Returns:
// - []T: The compacted slice with any falsey values removed.
func Compact[T comparable](s []T) []T {
	return slices.Compact(s)
}

// CompactFunc removes any elements from the input slice that satisfy the provided function.
//
// Parameters:
// - s: the input slice of type T.
// - fn: the function to apply to each element of the slice. It takes two values of type T and returns a boolean.
//
// Returns:
// - []T: the compacted slice with elements removed based on the provided function.
func CompactFunc[T interface{}](s []T, fn func(T, T) bool) []T {
	return slices.CompactFunc(s, fn)
}

func Compare[T cmp.Ordered](s []T, s2 []T) int {
	return slices.Compare(s, s2)
}

func CompareFunc[T interface{}, E interface{}](s []T, s2 []E, fn func(T, E) int) int {
	return slices.CompareFunc(s, s2, fn)
}

// Contains checks if a given slice contains a specific value.
//
// Parameters:
// - s: The slice to be checked.
// - v: The value to search for in the slice.
//
// Returns:
// - bool: True if the value is found in the slice, false otherwise.
func Contains[T comparable](s []T, v T) bool {
	return slices.Contains(s, v)
}

// ContrainsFunc checks if a given slice contains an element for which the provided function returns true.
//
// Parameters:
// - s: the slice to be checked.
// - fn: the function to apply to each element of the slice. It takes a value of type T and returns a boolean.
func ContrainsFunc[T interface{}](s []T, fn func(v T) bool) bool {
	return slices.ContainsFunc(s, fn)
}

// Copy creates a copy of the input slice.
//
// It takes a slice s as input and returns a new slice that is a copy of s.
// The function has a type parameter T that specifies the type of elements in the slice.
// The return type is []T, which is a slice of type T.
func Copy[T interface{}](s []T) []T {
	return slices.Clone(s)
}

// Chunk splits a slice into smaller chunks of a specified size.
//
// Parameters:
// - s: the input slice to be chunked.
// - chunkSize: the size of each chunk.
// - fn: optional callback function to be called for each chunk.
//
// Returns:
// - chunkedSlice: a 2D slice containing the chunked sub-slices.
func Chunk[T interface{}](s []T, chunkSize int, fn ...func(v []T)) [][]T {
	var chunkSlice []T

	chunkedSlice := make([][]T, 0)
	chunkedSize := int(math.Ceil(float64(len(s)) / float64(chunkSize)))

	var callback func(v []T)
	if len(fn) != 0 {
		callback = fn[0]
	}

	for i := 0; i < chunkedSize; i++ {
		if (i*chunkSize)+chunkSize <= (len(s) - 1) {
			chunkSlice = s[(i * chunkSize) : (i*chunkSize)+chunkSize]
		} else {
			chunkSlice = s[(i * chunkSize):]
		}

		callback(chunkSlice)

		chunkedSlice = append(chunkedSlice, chunkSlice)
	}

	return chunkedSlice
}

// Clear clears the given slice of type T.
//
// It takes a slice s as input and returns a new empty slice of type T.
func Clear[T interface{}](s []T) []T {
	clear(s)
	return s
}

// Dequeue removes and returns the first element from the given slice and returns the modified slice.
//
// Parameters:
// - s: The input slice from which the first element needs to be dequeued.
//
// Returns:
// - out: The modified slice after dequeuing the first element.
// - deq: The first element that was dequeued from the input slice.
func Dequeue[T interface{}](s []T) (out []T, deq T) {
	deq = s[0]
	out = s[1:]

	return out, deq
}

// Each applies a function to each element of a slice.
//
// Parameters:
// - s: the slice to iterate over
// - fn: the function to apply to each element of the slice
// Returns the modified slice.
func Each[T interface{}](s []T, fn func(v T)) []T {
	for _, v := range s {
		fn(v)
	}

	return s
}

// Enqueue adds an element to the front of a given slice.
//
// Parameters:
// - s: the slice to which the element will be added.
// - i: the element to be added to the slice.
//
// Returns:
// - []T: the updated slice with the new element added to the front.
func Enqueue[T interface{}](s []T, i T) []T {
	return Concat([]T{i}, s)
}

// Equal checks if two slices are equal.
//
// The function takes two slices, `s1` and `s2`, of elements of any type `T` that are comparable.
func Equal[T comparable](s1, s2 []T) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc checks if two slices are equal using a custom comparison function.
//
// The function takes two slices, `s1` and `s2`, of elements of any type, `T`, and a comparison function, `fn`.
func EqualFunc[T interface{}](s1, s2 []T, fn func(v1, v2 T) bool) bool {
	return slices.EqualFunc(s1, s2, fn)
}

func Except[T interface{}](s []T, fn func(T) bool) []T {
	_ = slices.DeleteFunc(s, fn)
	return s
}

// Filter filters a slice of elements based on a given predicate function.
//
// The function takes a slice, `s`, of elements of any type, `T`, and a predicate function, `fn`.
// The predicate function takes an element of type `T`
// If the predicate function returns `true` for an element, it is included in the filtered slice.
// The filtered slice is then returned as the result.
//
// Parameters:
//   - s: a slice of elements of any type, `T`.
//   - fn: a predicate function that takes an element of type `T`
//
// Return:
//   - filtered: a slice of elements of type `T` that satisfy the predicate function.
func Filter[T interface{}](s []T, fn func(v T) bool) []T {
	filtered := make([]T, 0)

	for _, v := range s {
		f := fn(v)
		if f {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// First returns the first element of a slice.
//
// It takes a slice of any type and returns the first element of that slice.
// The return type is the same as the type of the elements in the slice.
func First[T interface{}](s []T) T {
	return s[0]
}

// FlatMap applies a function to each element of a given 2D slice and returns a new slice
// containing the flattened results.
//
// Parameters:
//   - s: The 2D slice to be mapped.
func FlatMap[T interface{}, E interface{}](s [][]T, fn func(v []T) []E) []E {
	var mapped []E

	for _, v := range s {
		mapped = append(mapped, fn(v)...)
	}

	return mapped
}

// Index returns the index of the first occurrence of the value i in the slice s,
// or -1 if i is not present in s.
//
// Parameters:
// - s: the slice to search for the value i.
// - i: the value to search for in the slice s.
//
// Returns:
//   - int: the index of the first occurrence of the value i in the slice s,
//     or -1 if i is not present in
func Index[T comparable](s []T, i T) int {
	return slices.Index(s, i)
}

// IndexFunc returns the index of the first element in the slice `s` that satisfies the condition defined by the function `fn`.
//
// Parameters:
// - s: the slice to search for the element.
// - fn: the function that defines the condition to be satisfied by the element.
//
// Returns:
// - int: the index of the first element that satisfies the condition, or -1 if no element satisfies the condition.
func IndexFunc[T interface{}](s []T, fn func(T) bool) int {
	return slices.IndexFunc(s, fn)
}

// Insert inserts a value `v` into a slice `s` at index `i`.
//
// Parameters:
// - s: the slice to insert the value into.
// - i: the index at which to insert the value.
// - v: the value to insert into the slice.
//
// Returns:
// - []T: the modified slice with the value inserted at the specified index.
func Insert[T interface{}](s []T, i int, v T) []T {
	return slices.Insert(s, i, v)
}

// IsSorted checks if the given slice is sorted in ascending order.
//
// Parameters:
// - s: the slice to be checked for sorting.
//
// Returns:
// - bool: true if the slice is sorted, false otherwise.
func IsSorted[T cmp.Ordered](s []T) bool {
	return slices.IsSorted(s)
}

// IsSortedFunc checks if the given slice is sorted in ascending order based on the comparison function.
//
// Parameters:
// - s: the slice to be checked for sorting.
// - fn: the comparison function that takes two elements of type T and returns an integer. The function should return a negative value if the first element is less than the second, a positive value if the first element is greater than the second, and zero if the elements are equal.
//
// Returns:
// - bool: true if the slice is sorted, false otherwise.
func IsSortedFunc[T interface{}](s []T, fn func(T, T) int) bool {
	return slices.IsSortedFunc(s, fn)
}

// Last returns the last element of the given slice.
//
// The parameter s is a slice of type T.
// The function returns a value of type T.
func Last[T interface{}](s []T) T {
	return s[len(s)-1]
}

// Map applies a function to each element of a given slice and returns a new slice
// containing the results.
//
// Parameters:
//   - s: The slice to be mapped.
//   - fn: The function to be applied to each element of the slice. It takes two
//     arguments: the current element.
//
// Returns:
//   - A new slice containing the results of applying the function to each element
//     of the original slice.
func Map[T interface{}, E interface{}](s []T, fn func(v T) E) []E {
	var mapped []E

	for _, v := range s {
		mapped = append(mapped, fn(v))
	}

	return mapped
}

// Max returns the maximum element from a slice `s` based on the natural ordering of its elements.
//
// Parameters:
// - s: The slice of elements of type `T`.
//
// Returns:
// -
func Max[T cmp.Ordered](s []T) T {
	return slices.Max(s)
}

// MaxFunc returns the maximum element from a slice `s` based on the comparison function `fn`.
//
// Parameters:
// - s: The slice of elements of type `T`.
// - fn: The comparison function that takes two elements of type `T` and returns an integer.
//
// Returns:
// - The maximum element from the slice `s` based on the comparison function `fn`.
func MaxFunc[T interface{}](s []T, fn func(T, T) int) T {
	return slices.MaxFunc(s, fn)
}

// Merge merges multiple slices into one.
//
// The function takes in a slice s1 of type T and one or more additional slices s2 of type T.
// It appends all the elements of s2 to s1 and returns the merged slice.
//
// The return type is []T, the merged slice.
func Merge[T interface{}](s1 []T, s2 ...[]T) []T {
	merge := s1

	Each(s2, func(v []T) {
		merge = append(merge, v...)
	})

	return merge
}

// Min returns the minimum element from a slice `s` based on the natural ordering of its elements.
//
// Parameters:
// - s: The slice of elements of type `T`.
//
// Returns:
// - The minimum element of type `T`.
func Min[T cmp.Ordered](s []T) T {
	return slices.Min(s)
}

// MinFunc returns the minimum element in the given slice `s` based on the comparison function `fn`.
//
// Parameters:
// - s: The slice of elements of type `T`.
// - fn: The comparison function that takes two elements of type `T` and returns an integer. The function should return a negative value if the first element is less than the second, a positive value if the first element is greater than the second, and zero if the elements are equal.
//
// Returns:
// - The minimum element of type `T`.
func MinFunc[T interface{}](s []T, fn func(T, T) int) T {
	return slices.MinFunc(s, fn)
}

// Pop removes and returns the last element of the given slice and the remaining slice without the last element.
//
// Parameters:
// - s: the slice from which to remove the last element.
//
// Returns:
// - out: the slice without the last element.
// - pop: the last element of the slice.
func Pop[T interface{}](s []T) (out []T, pop T) {
	pop = s[len(s)-1]
	out = s[:len(s)-1]

	return out, pop
}

// Push pushes an element to the end of a slice.
//
// The function takes a slice and an element to be added to the slice.
// It returns a new slice with the element added to the end.
func Push[T interface{}](s []T, i T) []T {
	return Add(s, i)
}

// Remove removes an element from a slice at a specific index.
//
// Parameters:
//   - s: the slice from which to remove the element.
//   - index: the index of the element to remove.
//
// Return type:
//   - []T: the updated slice after removing the element.
func Remove[T interface{}](s []T, i int) []T {
	return slices.Delete(s, i, i+1)
}

// Replace replaces an element at a specific index in a slice with a new value.
//
// Parameters:
// - s: the slice to modify.
// - i: the index of the element to replace.
// - v: the new value to replace the element with.
//
// Return type:
// - []T: the modified slice with the element replaced.
func Replace[T interface{}](s []T, i int, v T) []T {
	return slices.Replace(s, i, i+1, v)
}

// Reverse reverses the elements of a slice.
//
// The function takes a slice `s` of any type `T` and returns a new slice `reverse`
// containing the elements of `s` in reverse order.
//
// Parameters:
// - s: the input slice to be reversed.
//
// Returns:
// - reverse: a new slice containing the elements of `s` in reverse order.
func Reverse[T interface{}](s []T) []T {
	slices.Reverse(s)
	return s
}

// Slice returns a new slice containing the elements of s starting from index start
// up to, but not including, index end.
//
// Parameters:
// - s: The original slice.
// - start: The starting index.
// - end: The ending index.
//
// Returns:
// - The new slice containing the specified elements.
func Slice[T interface{}](s []T, start int, end int) []T {
	return s[start:end]
}

// Sort sorts the elements of the given slice in ascending order based on the natural ordering of the elements.
//
// The function takes a slice `s` of elements of type `T` that implements the `cmp.Ordered` interface.
// The elements of the slice must be comparable.
func Sort[T cmp.Ordered](s []T) []T {
	slices.Sort(s)
	return s
}

// SortFunc sorts the elements of the given slice in ascending order based on the comparison function `fn`.
//
// Parameters:
// - s: The slice to be sorted.
// - fn: The comparison function that takes two elements of type `T` and returns an integer. The function should return a negative value if the first element is less than the second, a positive value if the first element is greater than the second, and zero if the elements are equal.
//
// Returns:
// - The sorted slice.
func SortFunc[T interface{}](s []T, fn func(T, T) int) []T {
	slices.SortFunc(s, fn)
	return s
}

// SortStableFunc sorts the elements of the given slice in a stable manner based on the comparison function.
//
// The function takes a slice `s` of elements of type `T` and a comparison function `fn`.
// The comparison function should take two elements of type `T` and return an integer.
// The function should return a negative value if the first element is less than the second,
// a positive value if the first element is greater than the second, and zero if the elements are equal.
//
// Returns:
// - The sorted slice `s`.
func SortStableFunc[T interface{}](s []T, fn func(T, T) int) []T {
	slices.SortStableFunc(s, fn)
	return s
}
