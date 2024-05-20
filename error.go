package gollection

import "errors"

var (
	ErrIsEmpty         = errors.New("collection is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
)
