package iterator_test

import (
	"log"
	"testing"

	"github.com/meteormin/gollection/pkg/iterator"
)

func TestNewIterator(t *testing.T) {
	iter := iterator.NewIterator([]int{1, 2, 3})

	for iter.HasNext() {
		next, err := iter.Next()
		if err != nil {
			t.Error(err)
		}
		log.Print(*next)
	}
}
