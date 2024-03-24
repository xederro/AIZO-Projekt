package quicksort

import (
	"github.com/xederro/AIZO-Projekt/algo"
	"slices"
	"testing"
)

func quicksortMiddleRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithRandomValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Middle_Random_Int(t *testing.T) {
	quicksortMiddleRandomTest[int](t)
}

func TestQuickSort_Middle_Random_Int32(t *testing.T) {
	quicksortMiddleRandomTest[int32](t)
}

func TestQuickSort_Middle_Random_Int64(t *testing.T) {
	quicksortMiddleRandomTest[int64](t)
}

func TestQuickSort_Middle_Random_Float32(t *testing.T) {
	quicksortMiddleRandomTest[float32](t)
}

func TestQuickSort_Middle_Random_Float64(t *testing.T) {
	quicksortMiddleRandomTest[float64](t)
}

func quicksortMiddleAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithAscendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Middle_Ascending_Int(t *testing.T) {
	quicksortMiddleAscendingTest[int](t)
}

func TestQuickSort_Middle_Ascending_Int32(t *testing.T) {
	quicksortMiddleAscendingTest[int32](t)
}

func TestQuickSort_Middle_Ascending_Int64(t *testing.T) {
	quicksortMiddleAscendingTest[int64](t)
}

func TestQuickSort_Middle_Ascending_Float32(t *testing.T) {
	quicksortMiddleAscendingTest[float32](t)
}

func TestQuickSort_Middle_Ascending_Float64(t *testing.T) {
	quicksortMiddleAscendingTest[float64](t)
}

func quicksortMiddleDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithDescendingValues())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Middle_Descending_Int(t *testing.T) {
	quicksortMiddleDescendingTest[int](t)
}

func TestQuickSort_Middle_Descending_Int32(t *testing.T) {
	quicksortMiddleDescendingTest[int32](t)
}

func TestQuickSort_Middle_Descending_Int64(t *testing.T) {
	quicksortMiddleDescendingTest[int64](t)
}

func TestQuickSort_Middle_Descending_Float32(t *testing.T) {
	quicksortMiddleDescendingTest[float32](t)
}

func TestQuickSort_Middle_Descending_Float64(t *testing.T) {
	quicksortMiddleDescendingTest[float64](t)
}

func quicksortMiddleSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortOneThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Middle_SortOneThirds_Int(t *testing.T) {
	quicksortMiddleSortOneThirdsTest[int](t)
}

func TestQuickSort_Middle_SortOneThirds_Int32(t *testing.T) {
	quicksortMiddleSortOneThirdsTest[int32](t)
}

func TestQuickSort_Middle_SortOneThirds_Int64(t *testing.T) {
	quicksortMiddleSortOneThirdsTest[int64](t)
}

func TestQuickSort_Middle_SortOneThirds_Float32(t *testing.T) {
	quicksortMiddleSortOneThirdsTest[float32](t)
}

func TestQuickSort_Middle_SortOneThirds_Float64(t *testing.T) {
	quicksortMiddleSortOneThirdsTest[float64](t)
}
func quicksortMiddleSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortTwoThirds())
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Middle_SortTwoThirds_Int(t *testing.T) {
	quicksortMiddleSortTwoThirdsTest[int](t)
}

func TestQuickSort_Middle_SortTwoThirds_Int32(t *testing.T) {
	quicksortMiddleSortTwoThirdsTest[int32](t)
}

func TestQuickSort_Middle_SortTwoThirds_Int64(t *testing.T) {
	quicksortMiddleSortTwoThirdsTest[int64](t)
}

func TestQuickSort_Middle_SortTwoThirds_Float32(t *testing.T) {
	quicksortMiddleSortTwoThirdsTest[float32](t)
}

func TestQuickSort_Middle_SortTwoThirds_Float64(t *testing.T) {
	quicksortMiddleSortTwoThirdsTest[float64](t)
}

func quicksortFirstRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithRandomValues()).
			SetPivotCalcFunc(First)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_First_Random_Int(t *testing.T) {
	quicksortFirstRandomTest[int](t)
}

func TestQuickSort_First_Random_Int32(t *testing.T) {
	quicksortFirstRandomTest[int32](t)
}

func TestQuickSort_First_Random_Int64(t *testing.T) {
	quicksortFirstRandomTest[int64](t)
}

func TestQuickSort_First_Random_Float32(t *testing.T) {
	quicksortFirstRandomTest[float32](t)
}

func TestQuickSort_First_Random_Float64(t *testing.T) {
	quicksortFirstRandomTest[float64](t)
}

func quicksortFirstAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithAscendingValues()).
			SetPivotCalcFunc(First)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_First_Ascending_Int(t *testing.T) {
	quicksortFirstAscendingTest[int](t)
}

func TestQuickSort_First_Ascending_Int32(t *testing.T) {
	quicksortFirstAscendingTest[int32](t)
}

func TestQuickSort_First_Ascending_Int64(t *testing.T) {
	quicksortFirstAscendingTest[int64](t)
}

func TestQuickSort_First_Ascending_Float32(t *testing.T) {
	quicksortFirstAscendingTest[float32](t)
}

func TestQuickSort_First_Ascending_Float64(t *testing.T) {
	quicksortFirstAscendingTest[float64](t)
}

func quicksortFirstDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithDescendingValues()).
			SetPivotCalcFunc(First)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_First_Descending_Int(t *testing.T) {
	quicksortFirstDescendingTest[int](t)
}

func TestQuickSort_First_Descending_Int32(t *testing.T) {
	quicksortFirstDescendingTest[int32](t)
}

func TestQuickSort_First_Descending_Int64(t *testing.T) {
	quicksortFirstDescendingTest[int64](t)
}

func TestQuickSort_First_Descending_Float32(t *testing.T) {
	quicksortFirstDescendingTest[float32](t)
}

func TestQuickSort_First_Descending_Float64(t *testing.T) {
	quicksortFirstDescendingTest[float64](t)
}

func quicksortFirstSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortOneThirds()).
			SetPivotCalcFunc(First)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_First_SortOneThirds_Int(t *testing.T) {
	quicksortFirstSortOneThirdsTest[int](t)
}

func TestQuickSort_First_SortOneThirds_Int32(t *testing.T) {
	quicksortFirstSortOneThirdsTest[int32](t)
}

func TestQuickSort_First_SortOneThirds_Int64(t *testing.T) {
	quicksortFirstSortOneThirdsTest[int64](t)
}

func TestQuickSort_First_SortOneThirds_Float32(t *testing.T) {
	quicksortFirstSortOneThirdsTest[float32](t)
}

func TestQuickSort_First_SortOneThirds_Float64(t *testing.T) {
	quicksortFirstSortOneThirdsTest[float64](t)
}
func quicksortFirstSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortTwoThirds()).
			SetPivotCalcFunc(First)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_First_SortTwoThirds_Int(t *testing.T) {
	quicksortFirstSortTwoThirdsTest[int](t)
}

func TestQuickSort_First_SortTwoThirds_Int32(t *testing.T) {
	quicksortFirstSortTwoThirdsTest[int32](t)
}

func TestQuickSort_First_SortTwoThirds_Int64(t *testing.T) {
	quicksortFirstSortTwoThirdsTest[int64](t)
}

func TestQuickSort_First_SortTwoThirds_Float32(t *testing.T) {
	quicksortFirstSortTwoThirdsTest[float32](t)
}

func TestQuickSort_First_SortTwoThirds_Float64(t *testing.T) {
	quicksortFirstSortTwoThirdsTest[float64](t)
}

func quicksortLastRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithRandomValues()).
			SetPivotCalcFunc(Last)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Last_Random_Int(t *testing.T) {
	quicksortLastRandomTest[int](t)
}

func TestQuickSort_Last_Random_Int32(t *testing.T) {
	quicksortLastRandomTest[int32](t)
}

func TestQuickSort_Last_Random_Int64(t *testing.T) {
	quicksortLastRandomTest[int64](t)
}

func TestQuickSort_Last_Random_Float32(t *testing.T) {
	quicksortLastRandomTest[float32](t)
}

func TestQuickSort_Last_Random_Float64(t *testing.T) {
	quicksortLastRandomTest[float64](t)
}

func quicksortLastAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithAscendingValues()).
			SetPivotCalcFunc(Last)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Last_Ascending_Int(t *testing.T) {
	quicksortLastAscendingTest[int](t)
}

func TestQuickSort_Last_Ascending_Int32(t *testing.T) {
	quicksortLastAscendingTest[int32](t)
}

func TestQuickSort_Last_Ascending_Int64(t *testing.T) {
	quicksortLastAscendingTest[int64](t)
}

func TestQuickSort_Last_Ascending_Float32(t *testing.T) {
	quicksortLastAscendingTest[float32](t)
}

func TestQuickSort_Last_Ascending_Float64(t *testing.T) {
	quicksortLastAscendingTest[float64](t)
}

func quicksortLastDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithDescendingValues()).
			SetPivotCalcFunc(Last)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Last_Descending_Int(t *testing.T) {
	quicksortLastDescendingTest[int](t)
}

func TestQuickSort_Last_Descending_Int32(t *testing.T) {
	quicksortLastDescendingTest[int32](t)
}

func TestQuickSort_Last_Descending_Int64(t *testing.T) {
	quicksortLastDescendingTest[int64](t)
}

func TestQuickSort_Last_Descending_Float32(t *testing.T) {
	quicksortLastDescendingTest[float32](t)
}

func TestQuickSort_Last_Descending_Float64(t *testing.T) {
	quicksortLastDescendingTest[float64](t)
}

func quicksortLastSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortOneThirds()).
			SetPivotCalcFunc(Last)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Last_SortOneThirds_Int(t *testing.T) {
	quicksortLastSortOneThirdsTest[int](t)
}

func TestQuickSort_Last_SortOneThirds_Int32(t *testing.T) {
	quicksortLastSortOneThirdsTest[int32](t)
}

func TestQuickSort_Last_SortOneThirds_Int64(t *testing.T) {
	quicksortLastSortOneThirdsTest[int64](t)
}

func TestQuickSort_Last_SortOneThirds_Float32(t *testing.T) {
	quicksortLastSortOneThirdsTest[float32](t)
}

func TestQuickSort_Last_SortOneThirds_Float64(t *testing.T) {
	quicksortLastSortOneThirdsTest[float64](t)
}
func quicksortLastSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortTwoThirds()).
			SetPivotCalcFunc(Last)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Last_SortTwoThirds_Int(t *testing.T) {
	quicksortLastSortTwoThirdsTest[int](t)
}

func TestQuickSort_Last_SortTwoThirds_Int32(t *testing.T) {
	quicksortLastSortTwoThirdsTest[int32](t)
}

func TestQuickSort_Last_SortTwoThirds_Int64(t *testing.T) {
	quicksortLastSortTwoThirdsTest[int64](t)
}

func TestQuickSort_Last_SortTwoThirds_Float32(t *testing.T) {
	quicksortLastSortTwoThirdsTest[float32](t)
}

func TestQuickSort_Last_SortTwoThirds_Float64(t *testing.T) {
	quicksortLastSortTwoThirdsTest[float64](t)
}

func quicksortRandomRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithRandomValues()).
			SetPivotCalcFunc(Random)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Random_Random_Int(t *testing.T) {
	quicksortRandomRandomTest[int](t)
}

func TestQuickSort_Random_Random_Int32(t *testing.T) {
	quicksortRandomRandomTest[int32](t)
}

func TestQuickSort_Random_Random_Int64(t *testing.T) {
	quicksortRandomRandomTest[int64](t)
}

func TestQuickSort_Random_Random_Float32(t *testing.T) {
	quicksortRandomRandomTest[float32](t)
}

func TestQuickSort_Random_Random_Float64(t *testing.T) {
	quicksortRandomRandomTest[float64](t)
}

func quicksortRandomAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithAscendingValues()).
			SetPivotCalcFunc(Random)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Random_Ascending_Int(t *testing.T) {
	quicksortRandomAscendingTest[int](t)
}

func TestQuickSort_Random_Ascending_Int32(t *testing.T) {
	quicksortRandomAscendingTest[int32](t)
}

func TestQuickSort_Random_Ascending_Int64(t *testing.T) {
	quicksortRandomAscendingTest[int64](t)
}

func TestQuickSort_Random_Ascending_Float32(t *testing.T) {
	quicksortRandomAscendingTest[float32](t)
}

func TestQuickSort_Random_Ascending_Float64(t *testing.T) {
	quicksortRandomAscendingTest[float64](t)
}

func quicksortRandomDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateWithDescendingValues()).
			SetPivotCalcFunc(Random)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Random_Descending_Int(t *testing.T) {
	quicksortRandomDescendingTest[int](t)
}

func TestQuickSort_Random_Descending_Int32(t *testing.T) {
	quicksortRandomDescendingTest[int32](t)
}

func TestQuickSort_Random_Descending_Int64(t *testing.T) {
	quicksortRandomDescendingTest[int64](t)
}

func TestQuickSort_Random_Descending_Float32(t *testing.T) {
	quicksortRandomDescendingTest[float32](t)
}

func TestQuickSort_Random_Descending_Float64(t *testing.T) {
	quicksortRandomDescendingTest[float64](t)
}

func quicksortRandomSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortOneThirds()).
			SetPivotCalcFunc(Random)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Random_SortOneThirds_Int(t *testing.T) {
	quicksortRandomSortOneThirdsTest[int](t)
}

func TestQuickSort_Random_SortOneThirds_Int32(t *testing.T) {
	quicksortRandomSortOneThirdsTest[int32](t)
}

func TestQuickSort_Random_SortOneThirds_Int64(t *testing.T) {
	quicksortRandomSortOneThirdsTest[int64](t)
}

func TestQuickSort_Random_SortOneThirds_Float32(t *testing.T) {
	quicksortRandomSortOneThirdsTest[float32](t)
}

func TestQuickSort_Random_SortOneThirds_Float64(t *testing.T) {
	quicksortRandomSortOneThirdsTest[float64](t)
}
func quicksortRandomSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewQuickSort(algo.NewArray[T](1000).
			PopulateAndSortTwoThirds()).
			SetPivotCalcFunc(Random)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestQuickSort_Random_SortTwoThirds_Int(t *testing.T) {
	quicksortRandomSortTwoThirdsTest[int](t)
}

func TestQuickSort_Random_SortTwoThirds_Int32(t *testing.T) {
	quicksortRandomSortTwoThirdsTest[int32](t)
}

func TestQuickSort_Random_SortTwoThirds_Int64(t *testing.T) {
	quicksortRandomSortTwoThirdsTest[int64](t)
}

func TestQuickSort_Random_SortTwoThirds_Float32(t *testing.T) {
	quicksortRandomSortTwoThirdsTest[float32](t)
}

func TestQuickSort_Random_SortTwoThirds_Float64(t *testing.T) {
	quicksortRandomSortTwoThirdsTest[float64](t)
}
