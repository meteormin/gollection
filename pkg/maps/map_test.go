package maps_test

import (
	"log"
	"testing"

	"github.com/meteormin/gollection/pkg/maps"
)

func TestCopy(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	c := maps.Copy(m)
	log.Print(c, m)
}

func TestMap(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Map(m, func(value int) int {
		return value + 1
	})

	log.Print(mapped)
}

func TestFilter(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Filter(m, func(value int) bool {
		log.Print(value)
		return value == 1
	})

	log.Print(mapped)
}

func TestPut(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Put(m, "c", 3)

	log.Print(mapped)
}

func TestDelete(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Delete(m, "a")

	log.Print(mapped)
}

func TestMerge(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	m2 := make(map[string]int)
	m2["c"] = 3
	m2["d"] = 4

	merge := maps.Merge(m, m2)

	log.Print(merge)
	log.Print(m)
}

func TestClear(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	clear := maps.Clear(m)
	log.Print(clear)
	log.Print(m)
}
