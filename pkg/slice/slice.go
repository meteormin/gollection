package slice

import (
	"math"
)

// Copy creates a copy of the input slice.
//
// It takes a slice s as input and returns a new slice that is a copy of s.
// The function has a type parameter T that specifies the type of elements in the slice.
// The return type is []T, which is a slice of type T.
func Copy[T interface{}](s []T) []T {
	copyS := make([]T, len(s))
	copy(copyS, s)
	return copyS
}

// Map applies a function to each element of a given slice and returns a new slice
// containing the results.
//
// Parameters:
//   - s: The slice to be mapped.
//   - fn: The function to be applied to each element of the slice. It takes two
//     arguments: the current element and its index.
//
// Returns:
//   - A new slice containing the results of applying the function to each element
//     of the original slice.
func Map[T interface{}, E interface{}](s []T, fn func(v T, i int) E) []E {
	var mapped []E

	for i, v := range s {
		mapped = append(mapped, fn(v, i))
	}

	return mapped
}

// Filter filters a slice of elements based on a given predicate function.
//
// The function takes a slice, `s`, of elements of any type, `T`, and a predicate function, `fn`.
// The predicate function takes an element of type `T` and its index, `i`, in the slice, and returns a boolean value.
// If the predicate function returns `true` for an element, it is included in the filtered slice.
// The filtered slice is then returned as the result.
//
// Parameters:
//   - s: a slice of elements of any type, `T`.
//   - fn: a predicate function that takes an element of type `T` and its index, `i`, in the slice, and returns a boolean value.
//
// Return:
//   - filtered: a slice of elements of type `T` that satisfy the predicate function.
func Filter[T interface{}](s []T, fn func(v T, i int) bool) []T {
	var filtered []T

	for i, v := range s {
		f := fn(v, i)
		if f {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Except filters a slice based on a given function.
//
// The function takes a slice `s` of type `T` and a function `fn` that
// takes a value `v` of type `T` and an index `i` of type `int` as
// arguments, and returns a boolean value. It iterates over the elements
// in the slice `s` and calls the function `fn` for each element. If the
// function `fn` returns `false` for an element, that element is appended
// to a new slice called `excepted`. Finally, the function returns the
// `excepted` slice.
//
// Parameters:
//   - s: The input slice of type `T`.
//   - fn: The function that takes a value of type `T` and an index of type
//     `int` and returns a boolean value.
//
// Return type:
// - []T: The filtered slice of type `T`.
func Except[T interface{}](s []T, fn func(v T, i int) bool) []T {
	var excepted []T

	for i, v := range s {
		f := fn(v, i)
		if !f {
			excepted = append(excepted, v)
		}
	}

	return excepted
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
func Chunk[T interface{}](s []T, chunkSize int, fn ...func(v []T, i int)) [][]T {
	chunkSlice := make([]T, 0)
	chunkedSlice := make([][]T, 0)
	chunkedSize := int(math.Ceil(float64(len(s) / chunkSize)))

	var callback func(v []T, i int)
	if len(fn) != 0 {
		callback = fn[0]
	}

	for i := 0; i < chunkedSize; i++ {
		if (i*chunkSize)+chunkSize <= (len(s) - 1) {
			chunkSlice = s[(i * chunkSize) : (i*chunkSize)+chunkSize]
		} else {
			chunkSlice = s[(i * chunkSize):]
		}

		callback(chunkSlice, i)

		chunkedSlice = append(chunkedSlice, chunkSlice)
	}

	return chunkedSlice
}

// For iterates over elements of type T in the slice s and applies the function fn to each element.
//
// Parameters:
//   - s: The slice of elements of type T.
//   - fn: The function to apply to each element of the slice s.
//
// Return:
//   - []T: The original slice s.
func For[T interface{}](s []T, fn func(v T, i int)) []T {
	for i, v := range s {
		fn(v, i)
	}

	return s
}

// Add appends a value to a slice.
//
// It takes a slice `s` and a value `v` as parameters.
// It returns a new slice with the value `v` appended to it.
func Add[T interface{}](s []T, v T) []T {
	return append(s, v)
}

// Remove removes an element from a slice at a specific index.
//
// Parameters:
//   - s: the slice from which to remove the element.
//   - index: the index of the element to remove.
//
// Return type:
//   - []T: the updated slice after removing the element.
func Remove[T interface{}](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

// Concat concatenates two slices of any type.
//
// It takes two slices of type T as input and returns a new slice of type T.
func Concat[T interface{}](s []T, s2 []T) []T {
	if len(s2) == 0 {
		return s
	}

	for _, v := range s2 {
		s = append(s, v)
	}

	return s
}

// Push pushes an element to the end of a slice.
//
// The function takes a slice and an element to be added to the slice.
// It returns a new slice with the element added to the end.
func Push[T interface{}](s []T, i T) []T {
	return Add(s, i)
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

// First returns the first element of a slice.
//
// It takes a slice of any type and returns the first element of that slice.
// The return type is the same as the type of the elements in the slice.
func First[T interface{}](s []T) T {
	return s[0]
}

// Last returns the last element of the given slice.
//
// The parameter s is a slice of type T.
// The function returns a value of type T.
func Last[T interface{}](s []T) T {
	return s[len(s)-1]
}

// Merge merges multiple slices into one.
//
// The function takes in a slice s1 of type T and one or more additional slices s2 of type T.
// It appends all the elements of s2 to s1 and returns the merged slice.
//
// The return type is []T, the merged slice.
func Merge[T interface{}](s1 []T, s2 ...[]T) []T {
	merge := s1

	For(s2, func(v []T, i int) {
		merge = append(merge, v...)
	})

	return merge
}

// Clear clears the given slice of type T.
//
// It takes a slice s as input and returns a new empty slice of type T.
func Clear[T interface{}](s []T) []T {
	s = make([]T, 0)

	return s
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
	reverse := make([]T, 0)
	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}

	return reverse
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
