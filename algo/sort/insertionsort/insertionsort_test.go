package insertionsort

import (
	"AIZO-Projekt/algo"
	"slices"
	"testing"
)

func insertionsortRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewInsertionSort(algo.NewArray[T](1000).PopulateWithRandomValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestInsertionSort_Random_Int(t *testing.T) {
	insertionsortRandomTest[int](t)
}

func TestInsertionSort_Random_Int32(t *testing.T) {
	insertionsortRandomTest[int32](t)
}

func TestInsertionSort_Random_Int64(t *testing.T) {
	insertionsortRandomTest[int64](t)
}

func TestInsertionSort_Random_Float32(t *testing.T) {
	insertionsortRandomTest[float32](t)
}

func TestInsertionSort_Random_Float64(t *testing.T) {
	insertionsortRandomTest[float64](t)
}

func insertionsortAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewInsertionSort(algo.NewArray[T](1000).PopulateWithAscendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestInsertionSort_Ascending_Int(t *testing.T) {
	insertionsortAscendingTest[int](t)
}

func TestInsertionSort_Ascending_Int32(t *testing.T) {
	insertionsortAscendingTest[int32](t)
}

func TestInsertionSort_Ascending_Int64(t *testing.T) {
	insertionsortAscendingTest[int64](t)
}

func TestInsertionSort_Ascending_Float32(t *testing.T) {
	insertionsortAscendingTest[float32](t)
}

func TestInsertionSort_Ascending_Float64(t *testing.T) {
	insertionsortAscendingTest[float64](t)
}

func insertionsortDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewInsertionSort(algo.NewArray[T](1000).PopulateWithDescendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestInsertionSort_Descending_Int(t *testing.T) {
	insertionsortDescendingTest[int](t)
}

func TestInsertionSort_Descending_Int32(t *testing.T) {
	insertionsortDescendingTest[int32](t)
}

func TestInsertionSort_Descending_Int64(t *testing.T) {
	insertionsortDescendingTest[int64](t)
}

func TestInsertionSort_Descending_Float32(t *testing.T) {
	insertionsortDescendingTest[float32](t)
}

func TestInsertionSort_Descending_Float64(t *testing.T) {
	insertionsortDescendingTest[float64](t)
}

func insertionsortSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewInsertionSort(algo.NewArray[T](1000).PopulateAndSortOneThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestInsertionSort_SortOneThirds_Int(t *testing.T) {
	insertionsortSortOneThirdsTest[int](t)
}

func TestInsertionSort_SortOneThirds_Int32(t *testing.T) {
	insertionsortSortOneThirdsTest[int32](t)
}

func TestInsertionSort_SortOneThirds_Int64(t *testing.T) {
	insertionsortSortOneThirdsTest[int64](t)
}

func TestInsertionSort_SortOneThirds_Float32(t *testing.T) {
	insertionsortSortOneThirdsTest[float32](t)
}

func TestInsertionSort_SortOneThirds_Float64(t *testing.T) {
	insertionsortSortOneThirdsTest[float64](t)
}
func insertionsortSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewInsertionSort(algo.NewArray[T](1000).PopulateAndSortTwoThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestInsertionSort_SortTwoThirds_Int(t *testing.T) {
	insertionsortSortTwoThirdsTest[int](t)
}

func TestInsertionSort_SortTwoThirds_Int32(t *testing.T) {
	insertionsortSortTwoThirdsTest[int32](t)
}

func TestInsertionSort_SortTwoThirds_Int64(t *testing.T) {
	insertionsortSortTwoThirdsTest[int64](t)
}

func TestInsertionSort_SortTwoThirds_Float32(t *testing.T) {
	insertionsortSortTwoThirdsTest[float32](t)
}

func TestInsertionSort_SortTwoThirds_Float64(t *testing.T) {
	insertionsortSortTwoThirdsTest[float64](t)
}
