package maps

// Copy creates a copy of the input map.
//
// It takes a map m as input and returns a new map that is a copy of m.
// The function has two type parameters: k, which specifies the type of keys in the map,
// and v, which specifies the type of values in the map.
// The return type is map
func Copy[k comparable, v interface{}](m map[k]v) map[k]v {
	copyM := make(map[k]v, len(m))
	for key, value := range m {
		copyM[key] = value
	}

	return copyM
}

// Map applies a function to each key-value pair in the input map and returns a new map with the results.
//
// Parameters:
//   - m: The input map of type map[k]v.
//   - fn: The function to apply to each value in the input map. It takes a value of type v as input and returns a value of type e.
//
// Return type:
//   - map[k]e: A new map with the results of applying the function to each key-value pair in the input map.
func Map[k comparable, v interface{}, e interface{}](m map[k]v, fn func(value v) e) map[k]e {
	mapped := make(map[k]e)

	for key, value := range m {
		mapped[key] = fn(value)
	}

	return mapped
}

// Filter filters a map based on a given function.
//
// The function takes a map `m` of type `map[k]v` and a function `fn` that takes a value `v` of type `v` as
// arguments, and returns a boolean value. It iterates over the key-value pairs in the map `m` and calls the
// function `fn` for each pair. If the function `fn` returns `true` for a pair, that pair is included in the
// filtered map. The filtered map is then returned as the result.
//
// Parameters:
//   - m: The map to filter.
//   - fn: The function that takes a value of type `v` and returns a boolean value.
//
// Return type:
//   - map[k]v: The filtered map.
func Filter[k comparable, v interface{}](m map[k]v, fn func(value v) bool) map[k]v {
	filtered := make(map[k]v)

	for key, value := range m {
		f := fn(value)
		if f {
			filtered[key] = value
		}
	}

	return filtered
}

// Each applies the given function to each value in the map and returns the original map.
//
// Parameters:
// - m: The map to iterate over.
// - fn: The function to apply to each value in the map.
//
// Return type:
// - map[k]v: The original map.
func Each[k comparable, v interface{}](m map[k]v, fn func(value v)) map[k]v {
	for _, value := range m {
		fn(value)
	}

	return m
}

// Put adds or updates a key-value pair in the given map.
//
// Parameters:
// - m: The map to modify.
// - key: The key to add or update.
// - value: The value to associate with the key.
func Put[k comparable, v interface{}](m map[k]v, key k, value v) map[k]v {
	m[key] = value

	return m
}

// Put adds or updates a key-value pair in the given map.
//
// Parameters:
// - m: The map to modify.
// - key: The key to add or update.
// - value: The value to associate with the key.
func Delete[k comparable, v interface{}](m map[k]v, key k) map[k]v {
	delete(m, key)

	return m
}

// Each applies the given function to each value in the map and returns the original map.
//
// Parameters:
// - m: The map to iterate over.
// - fn: The function to apply to each value in the map.
//
// Return type:
// - map[k]v: The original map.
func Merge[k comparable, v interface{}](m1 map[k]v, m2 ...map[k]v) map[k]v {
	merge := Copy(m1)

	for _, m := range m2 {
		for key, value := range m {
			merge[key] = value
		}
	}

	return merge
}

// Clear clears the given map and returns an empty map of the same type.
//
// Parameters:
// - m: The map to clear.
//
// Return:
// - The cleared map.
func Clear[k comparable, v interface{}](m map[k]v) map[k]v {
	clear(m)
	return m
}
