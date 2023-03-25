package maps_test

import (
	"github.com/miniyus/gollection/pkg/maps"
	"log"
	"testing"
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
	mapped := maps.Map(m, func(value int, key string) int {
		return value + 1
	})

	log.Print(mapped)
}

func TestFor(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.For(m, func(value int, key string) {
		log.Print(key, value)
	})

	log.Print(mapped)
}

func TestFilter(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Filter(m, func(value int, key string) bool {
		log.Print(key, value)
		return key == "a"
	})

	log.Print(mapped)
}

func TestExcept(t *testing.T) {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	mapped := maps.Except(m, func(value int, key string) bool {
		log.Print(key, value)
		return key == "a"
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
