package maps

func Copy[k comparable, v interface{}](m map[k]v) map[k]v {
	copyM := make(map[k]v, len(m))
	for key, value := range m {
		copyM[key] = value
	}

	return copyM
}

func Map[k comparable, v interface{}, e interface{}](m map[k]v, fn func(value v, key k) e) map[k]e {
	mapped := make(map[k]e)

	for key, value := range m {
		mapped[key] = fn(value, key)
	}

	return mapped
}

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

func For[k comparable, v interface{}](m map[k]v, fn func(value v, key k)) map[k]v {
	for key, value := range m {
		fn(value, key)
	}

	return m
}

func Put[k comparable, v interface{}](m map[k]v, key k, value v) map[k]v {
	m[key] = value

	return m
}

func Delete[k comparable, v interface{}](m map[k]v, key k) map[k]v {
	delete(m, key)

	return m
}

func Merge[k comparable, v interface{}](m1 map[k]v, m2 ...map[k]v) map[k]v {
	merge := Copy(m1)

	for _, m := range m2 {
		For(m, func(value v, key k) {
			merge[key] = value
		})
	}

	return merge
}

func Clear[k comparable, v interface{}](m map[k]v) map[k]v {
	m = make(map[k]v)

	return m
}
