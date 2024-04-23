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

// Map applies a given function to each key-value pair in the map and returns a new map with the results.
//
// Parameters:
//   - m: The map to iterate over.
//   - fn: The function to apply to each key-value pair.
//
// Return type: The mapped map.
func Map[k comparable, v interface{}, e interface{}](m map[k]v, fn func(value v, key k) e) map[k]e {
	mapped := make(map[k]e)

	for key, value := range m {
		mapped[key] = fn(value, key)
	}

	return mapped
}

// / Filter filters a map based on a given function.
//
// The function takes a map `m` of type `map[k]v` and a function `fn` that takes a value `v` of type `v` and a key `k` of type `k` as arguments, and returns a boolean value. It iterates over the key-value pairs in the map `m` and calls the function `fn` for each pair. If the function `fn` returns `true` for a pair, that pair is included in the filtered map. The filtered map is then returned as the result.
//
// Parameters:
//   - m: The map to filter.
//   - fn: The function that takes a value of type `v` and a key of type `k` and returns a boolean value.
//
// Return type:
//   - map[k]v: The filtered map.
func Filter[k comparable, v interface{}](m map[k]v, fn func(value v, key k) bool) map[k]v {
	filtered := make(map[k]v)

	for key, value := range m {
		f := fn(value, key)
		if f {
			filtered[key] = value
		}
	}

	return filtered
}

// Except filters a map based on a given function.
//
// The function takes a map `m` of type `map[k]v` and a function `fn` that takes a value `v` of type `v` and a key `k` of type `k` as arguments, and returns a boolean value. It iterates over the key-value pairs in the map `m` and calls the function `fn` for each pair. If the function `fn` returns `false` for a pair, that pair is included in the filtered map. The filtered map is then returned as the result.
//
// Parameters:
//   - m: The map to filter.
//   - fn: The function that takes a value of type `v` and a key of type `k` and returns a boolean value.
//
// Return type:
//   - map[k]v: The filtered map.
func Except[k comparable, v interface{}](m map[k]v, fn func(value v, key k) bool) map[k]v {
	excepted := make(map[k]v)

	for key, value := range m {
		f := fn(value, key)
		if !f {
			excepted[key] = value
		}
	}

	return excepted
}

// For iterates over the key-value pairs in the given map and applies the provided function to each pair.
//
// Parameters:
//   - m: The map to iterate over.
//   - fn: The function to apply to each key-value pair.
//
// Return:
//   - The original map.
//
// Deprecated: Use the Each package instead.
func For[k comparable, v interface{}](m map[k]v, fn func(value v, key k)) map[k]v {
	for key, value := range m {
		fn(value, key)
	}

	return m
}

// Each iterates over the key-value pairs in the given map and applies the provided function to each pair.
//
// Parameters:
//   - m: The map to iterate over.
//   - fn: The function to apply to each key-value pair.
//
// Return:
//   - The original map.
func Each[k comparable, v interface{}](m map[k]v, fn func(value v, key k)) map[k]v {
	for key, value := range m {
		fn(value, key)
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

// Merge merges multiple maps into a single map.
//
// The function takes a map `m1` of type `map[k]v` and any number of additional maps `m2` of type `map[k]v`.
// It returns a new map of type `map[k]v` that contains all the key-value pairs from `m1` and `m2`.
func Merge[k comparable, v interface{}](m1 map[k]v, m2 ...map[k]v) map[k]v {
	merge := Copy(m1)

	for _, m := range m2 {
		For(m, func(value v, key k) {
			merge[key] = value
		})
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
	m = make(map[k]v)

	return m
}
