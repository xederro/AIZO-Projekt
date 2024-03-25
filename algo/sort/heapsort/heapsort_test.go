package heapsort

import (
	"github.com/xederro/AIZO-Projekt/algo"
	"slices"
	"testing"
)

func heapsortRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewHeapSort(algo.NewArray[T](1000).PopulateWithRandomValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestHeapSort_Random_Int(t *testing.T) {
	heapsortRandomTest[int8](t)
}

func TestHeapSort_Random_Int32(t *testing.T) {
	heapsortRandomTest[int32](t)
}

func TestHeapSort_Random_Int64(t *testing.T) {
	heapsortRandomTest[int64](t)
}

func TestHeapSort_Random_Float32(t *testing.T) {
	heapsortRandomTest[float32](t)
}

func TestHeapSort_Random_Float64(t *testing.T) {
	heapsortRandomTest[float64](t)
}

func heapsortAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewHeapSort(algo.NewArray[T](1000).PopulateWithAscendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestHeapSort_Ascending_Int(t *testing.T) {
	heapsortAscendingTest[int8](t)
}

func TestHeapSort_Ascending_Int32(t *testing.T) {
	heapsortAscendingTest[int32](t)
}

func TestHeapSort_Ascending_Int64(t *testing.T) {
	heapsortAscendingTest[int64](t)
}

func TestHeapSort_Ascending_Float32(t *testing.T) {
	heapsortAscendingTest[float32](t)
}

func TestHeapSort_Ascending_Float64(t *testing.T) {
	heapsortAscendingTest[float64](t)
}

func heapsortDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewHeapSort(algo.NewArray[T](1000).PopulateWithDescendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestHeapSort_Descending_Int(t *testing.T) {
	heapsortDescendingTest[int8](t)
}

func TestHeapSort_Descending_Int32(t *testing.T) {
	heapsortDescendingTest[int32](t)
}

func TestHeapSort_Descending_Int64(t *testing.T) {
	heapsortDescendingTest[int64](t)
}

func TestHeapSort_Descending_Float32(t *testing.T) {
	heapsortDescendingTest[float32](t)
}

func TestHeapSort_Descending_Float64(t *testing.T) {
	heapsortDescendingTest[float64](t)
}

func heapsortSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewHeapSort(algo.NewArray[T](1000).PopulateAndSortOneThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestHeapSort_SortOneThirds_Int(t *testing.T) {
	heapsortSortOneThirdsTest[int8](t)
}

func TestHeapSort_SortOneThirds_Int32(t *testing.T) {
	heapsortSortOneThirdsTest[int32](t)
}

func TestHeapSort_SortOneThirds_Int64(t *testing.T) {
	heapsortSortOneThirdsTest[int64](t)
}

func TestHeapSort_SortOneThirds_Float32(t *testing.T) {
	heapsortSortOneThirdsTest[float32](t)
}

func TestHeapSort_SortOneThirds_Float64(t *testing.T) {
	heapsortSortOneThirdsTest[float64](t)
}
func heapsortSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewHeapSort(algo.NewArray[T](1000).PopulateAndSortTwoThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestHeapSort_SortTwoThirds_Int(t *testing.T) {
	heapsortSortTwoThirdsTest[int8](t)
}

func TestHeapSort_SortTwoThirds_Int32(t *testing.T) {
	heapsortSortTwoThirdsTest[int32](t)
}

func TestHeapSort_SortTwoThirds_Int64(t *testing.T) {
	heapsortSortTwoThirdsTest[int64](t)
}

func TestHeapSort_SortTwoThirds_Float32(t *testing.T) {
	heapsortSortTwoThirdsTest[float32](t)
}

func TestHeapSort_SortTwoThirds_Float64(t *testing.T) {
	heapsortSortTwoThirdsTest[float64](t)
}
