package slice_test

import (
	"log"
	"testing"

	"github.com/meteormin/gollection/pkg/slice"
)

func TestChunk(t *testing.T) {
	testData := make([]int, 10)
	slice.Chunk(testData, 2, func(v []int) {
		if len(v) != 2 {
			t.Error(len(v))
		}
	})
}

func TestConcat(t *testing.T) {
	testData := make([]int, 3)
	rs := slice.Concat(testData, []int{4, 5})
	if len(rs) != 5 {
		t.Error(len(rs), rs)
	}
}

func TestEach(t *testing.T) {
	testData := make([]int, 10)
	slice.Each(testData, func(v int) {
		if v >= 10 {
			t.Error(v)
		}
	})
}

func TestMap(t *testing.T) {
	testData := make([]int, 10)
	rs := slice.Map(testData, func(v int) int {
		return v + 1
	})

	log.Print(rs)
}

func TestFlatMap(t *testing.T) {
	testData := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rs := slice.FlatMap(testData, func(v []int) []float64 {
		return slice.Map(v, func(v int) float64 {
			return float64(v) + float64(v)*0.1
		})
	})

	t.Log(rs)
}

func TestFilter(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Filter(testData, func(v int) bool {
		return v == 1
	})

	if rs[0] != 1 {
		t.Error(rs[0])
	}
}

func TestAdd(t *testing.T) {
	testData := []int{1, 2, 3, 4}
	rs := slice.Add(testData, 5)

	if rs[len(rs)-1] == 4 {
		t.Error(rs[len(rs)-1])
	}
}

func TestRemove(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Remove(testData, 0)

	if rs[0] == 1 {
		t.Error(rs[0])
	}
}

func TestClear(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Clear(testData)

	if len(rs) != 0 {
		t.Error("failed clear")
	}
}

func TestCopy(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Copy(testData)

	if testData[0] != rs[0] {
		t.Error("failed copy")
	}
}

func TestLast(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Last(testData)

	if rs != 5 {
		t.Error(rs)
	}
}

func TestPush(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Push(testData, 1)
	v := slice.Last(rs)
	if v != 1 {
		t.Error(v)
	}
}

func TestPop(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	_, pop := slice.Pop(testData)
	if pop != 5 {
		t.Error(pop)
	}
}

func TestFirst(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.First(testData)

	if rs != 1 {
		t.Error(rs)
	}
}

func TestReverse(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Reverse(testData)

	if rs[0] != 5 {
		t.Error(rs)
	}
}

func TestMerge(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Merge(testData, []int{4, 3, 2})
	log.Print(rs)
	if rs[len(rs)-1] != 2 {
		t.Error(rs)
	}
}

func TestSlice(t *testing.T) {
	testData := []int{1, 2, 3, 4, 5}
	rs := slice.Slice(testData, 0, 1)
	log.Print(rs)
}
