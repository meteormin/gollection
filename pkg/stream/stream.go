package stream

type Stream[T interface{}] interface {
	Get() ([]T, error)
	First() (*T, error)
	Last() (*T, error)
}
