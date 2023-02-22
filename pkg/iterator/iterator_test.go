package iterator_test

import (
	"github.com/miniyus/gollection/pkg/iterator"
	"log"
	"testing"
	"time"
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

func TestNewAsyncIterator(t *testing.T) {
	iter := iterator.NewAsyncIterator(make([]string, 100))

	go func() {
		for {
			iter.Next(func(v string, i int) error {
				log.Print(v+".", i)
				time.Sleep(3 * time.Millisecond)
				return nil
			})
		}
	}()

	for i, v := range make([]string, 100) {
		log.Print(i, v+".")
		time.Sleep(3 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
}
