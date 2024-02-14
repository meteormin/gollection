package gollection_test

import (
	"github.com/meteormin/gollection"
	"log"
	"testing"
)

var testData = []int{
	1, 2, 3,
}

func TestBaseCollection_Items(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	items := collection.Items()

	for i, n := range items {
		log.Print(i, n)
		if n != testData[i] {
			t.Errorf("not match! %d:%d", i, n)
		}
	}
}

func TestBaseCollection_Get(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	if testData[1] != collection.Get(1) {
		t.Errorf("test failed: testData(%v) != collectData(%v)", testData[1], collection.Get(1))
	}

	log.Print(collection.Get(1))
}

func TestBaseCollection_All(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	all := collection.All()

	for i, v := range all {
		log.Print(i, v)
		if v != testData[i] {
			t.Errorf("not match! %d:%d", i, v)
		}
	}
}

func TestBaseCollection_Count(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	if len(testData) != collection.Count() {
		t.Errorf("diff count... test: %d, collection: %d", len(testData), collection.Count())
	}
}

func TestBaseCollection_Add(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	collection.Add(4)
	i := collection.Items()

	log.Print(i[len(i)-1])

	if 4 != i[len(i)-1] {
		t.Error("result must be 4")
	}

}

func TestBaseCollection_Chunk(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	chunked := collection.Chunk(1, func(n []int, i int) {
		log.Print(n, i)
	})

	if len(chunked) != len(testData) {
		t.Error("failed chunk!")
	}
}

func TestBaseCollection_Concat(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	collection.Concat([]int{4, 5, 6}...)

	resultData := []int{1, 2, 3, 4, 5, 6}

	for i, n := range collection.Items() {
		log.Print(i, n)
		if n != resultData[i] {
			t.Errorf("not match!! %d:%d", i, n)
		}
	}
}

func TestBaseCollection_Except(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	result := collection.Except(func(v int, i int) bool {
		return v == 1
	})

	for _, n := range result.Items() {
		if n == 1 {
			t.Error("FAIL!")
		}
	}
}

func TestBaseCollection_Filter(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	resultCollection := collection.Filter(func(v int, i int) bool {
		return v == 1
	})

	for _, n := range resultCollection.Items() {
		if n != 1 {
			t.Error("FAIL!")
		}
	}
}

func TestBaseCollection_For(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	collection.For(func(v int, i int) {
		if v != testData[i] {
			t.Error("FAIL!")
		}
	})
}

func TestBaseCollection_Map(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	result := collection.Map(func(v int, i int) int {
		return i + 1
	})

	for i, n := range result.Items() {
		log.Print(i, n)
		if n != testData[i] {
			t.Error("Fail")
		}
	}
}

func TestBaseCollection_Remove(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	err := collection.Remove(0)
	if err != nil {
		t.Error(err)
	}

	collection.For(func(v int, i int) {
		log.Print(i, v)
		if v == 1 {
			t.Error("not removed")
		}
	})
}

func TestBaseCollection_First(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	first, err := collection.First()
	if err != nil {
		t.Error(err)
	}
	log.Print(*first)
}

func TestBaseCollection_IsEmpty(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	if collection.IsEmpty() {
		t.Error("test data is not empty!")
	}

}

func TestBaseCollection_Last(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	last, err := collection.Last()
	if err != nil {
		t.Error(err)
	}
	log.Print(*last)
}

func TestBaseCollection_Merge(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	merge := []int{4, 5, 6}
	mergeCollection := collection.Merge(merge)

	if mergeCollection.Count() != 6 {
		t.Error("failed merge...")
	}

	mergeCollection.For(func(v int, i int) {
		log.Print(v)
	})
}

func TestBaseCollection_Pop(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	pop, err := collection.Pop()
	if err != nil {
		t.Error(err)
	}
	log.Print(*pop)
}

func TestBaseCollection_Push(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	collection.Push(4)
	last, err := collection.Last()
	if err != nil {
		t.Error(err)
	}

	log.Print(*last)
}

func TestBaseCollection_Copy(t *testing.T) {
	var collection = gollection.NewCollection(testData)
	collection.Push(4)
	l, err := collection.Last()
	if err != nil {
		t.Error(err)
	}

	log.Print(*l)

	pop, err := collection.Copy().Pop()
	if err != nil {
		t.Error(err)
	}

	l, err = collection.Last()
	if err != nil {
		t.Error(err)
	}

	log.Print(*pop)
	log.Print(*l)
}

func TestBaseCollection_Slice(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	log.Print(collection.Slice(0, 1))
}

func TestBaseCollection_Reverse(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	log.Print(collection.Reverse())
}

func TestBaseCollection_Sort(t *testing.T) {
	var collection = gollection.NewCollection(testData)

	log.Print(collection.Sort(func(i, j int) bool {
		return i > j
	}))
}
