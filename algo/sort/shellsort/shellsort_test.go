package shellsort

import (
	"AIZO-Projekt/algo"
	"slices"
	"testing"
)

func shellsortShellRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithRandomValues()).SetGapCalcFunc(Shell)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Shell_Random_Int(t *testing.T) {
	shellsortShellRandomTest[int](t)
}

func TestShellSort_Shell_Random_Int32(t *testing.T) {
	shellsortShellRandomTest[int32](t)
}

func TestShellSort_Shell_Random_Int64(t *testing.T) {
	shellsortShellRandomTest[int64](t)
}

func TestShellSort_Shell_Random_Float32(t *testing.T) {
	shellsortShellRandomTest[float32](t)
}

func TestShellSort_Shell_Random_Float64(t *testing.T) {
	shellsortShellRandomTest[float64](t)
}

func shellsortShellAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithAscendingValues()).SetGapCalcFunc(Shell)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Shell_Ascending_Int(t *testing.T) {
	shellsortShellAscendingTest[int](t)
}

func TestShellSort_Shell_Ascending_Int32(t *testing.T) {
	shellsortShellAscendingTest[int32](t)
}

func TestShellSort_Shell_Ascending_Int64(t *testing.T) {
	shellsortShellAscendingTest[int64](t)
}

func TestShellSort_Shell_Ascending_Float32(t *testing.T) {
	shellsortShellAscendingTest[float32](t)
}

func TestShellSort_Shell_Ascending_Float64(t *testing.T) {
	shellsortShellAscendingTest[float64](t)
}

func shellsortShellDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithDescendingValues()).SetGapCalcFunc(Shell)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Shell_Descending_Int(t *testing.T) {
	shellsortShellDescendingTest[int](t)
}

func TestShellSort_Shell_Descending_Int32(t *testing.T) {
	shellsortShellDescendingTest[int32](t)
}

func TestShellSort_Shell_Descending_Int64(t *testing.T) {
	shellsortShellDescendingTest[int64](t)
}

func TestShellSort_Shell_Descending_Float32(t *testing.T) {
	shellsortShellDescendingTest[float32](t)
}

func TestShellSort_Shell_Descending_Float64(t *testing.T) {
	shellsortShellDescendingTest[float64](t)
}

func shellsortShellSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortOneThirds()).SetGapCalcFunc(Shell)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Shell_SortOneThirds_Int(t *testing.T) {
	shellsortShellSortOneThirdsTest[int](t)
}

func TestShellSort_Shell_SortOneThirds_Int32(t *testing.T) {
	shellsortShellSortOneThirdsTest[int32](t)
}

func TestShellSort_Shell_SortOneThirds_Int64(t *testing.T) {
	shellsortShellSortOneThirdsTest[int64](t)
}

func TestShellSort_Shell_SortOneThirds_Float32(t *testing.T) {
	shellsortShellSortOneThirdsTest[float32](t)
}

func TestShellSort_Shell_SortOneThirds_Float64(t *testing.T) {
	shellsortShellSortOneThirdsTest[float64](t)
}
func shellsortShellSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortTwoThirds()).SetGapCalcFunc(Shell)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Shell_SortTwoThirds_Int(t *testing.T) {
	shellsortShellSortTwoThirdsTest[int](t)
}

func TestShellSort_Shell_SortTwoThirds_Int32(t *testing.T) {
	shellsortShellSortTwoThirdsTest[int32](t)
}

func TestShellSort_Shell_SortTwoThirds_Int64(t *testing.T) {
	shellsortShellSortTwoThirdsTest[int64](t)
}

func TestShellSort_Shell_SortTwoThirds_Float32(t *testing.T) {
	shellsortShellSortTwoThirdsTest[float32](t)
}

func TestShellSort_Shell_SortTwoThirds_Float64(t *testing.T) {
	shellsortShellSortTwoThirdsTest[float64](t)
}

func shellsortHibbardRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithRandomValues()).SetGapCalcFunc(Hibbard)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Hibbard_Random_Int(t *testing.T) {
	shellsortHibbardRandomTest[int](t)
}

func TestShellSort_Hibbard_Random_Int32(t *testing.T) {
	shellsortHibbardRandomTest[int32](t)
}

func TestShellSort_Hibbard_Random_Int64(t *testing.T) {
	shellsortHibbardRandomTest[int64](t)
}

func TestShellSort_Hibbard_Random_Float32(t *testing.T) {
	shellsortHibbardRandomTest[float32](t)
}

func TestShellSort_Hibbard_Random_Float64(t *testing.T) {
	shellsortHibbardRandomTest[float64](t)
}

func shellsortHibbardAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithAscendingValues()).SetGapCalcFunc(Hibbard)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Hibbard_Ascending_Int(t *testing.T) {
	shellsortHibbardAscendingTest[int](t)
}

func TestShellSort_Hibbard_Ascending_Int32(t *testing.T) {
	shellsortHibbardAscendingTest[int32](t)
}

func TestShellSort_Hibbard_Ascending_Int64(t *testing.T) {
	shellsortHibbardAscendingTest[int64](t)
}

func TestShellSort_Hibbard_Ascending_Float32(t *testing.T) {
	shellsortHibbardAscendingTest[float32](t)
}

func TestShellSort_Hibbard_Ascending_Float64(t *testing.T) {
	shellsortHibbardAscendingTest[float64](t)
}

func shellsortHibbardDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithDescendingValues()).SetGapCalcFunc(Hibbard)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Hibbard_Descending_Int(t *testing.T) {
	shellsortHibbardDescendingTest[int](t)
}

func TestShellSort_Hibbard_Descending_Int32(t *testing.T) {
	shellsortHibbardDescendingTest[int32](t)
}

func TestShellSort_Hibbard_Descending_Int64(t *testing.T) {
	shellsortHibbardDescendingTest[int64](t)
}

func TestShellSort_Hibbard_Descending_Float32(t *testing.T) {
	shellsortHibbardDescendingTest[float32](t)
}

func TestShellSort_Hibbard_Descending_Float64(t *testing.T) {
	shellsortHibbardDescendingTest[float64](t)
}

func shellsortHibbardSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortOneThirds()).SetGapCalcFunc(Hibbard)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Hibbard_SortOneThirds_Int(t *testing.T) {
	shellsortHibbardSortOneThirdsTest[int](t)
}

func TestShellSort_Hibbard_SortOneThirds_Int32(t *testing.T) {
	shellsortHibbardSortOneThirdsTest[int32](t)
}

func TestShellSort_Hibbard_SortOneThirds_Int64(t *testing.T) {
	shellsortHibbardSortOneThirdsTest[int64](t)
}

func TestShellSort_Hibbard_SortOneThirds_Float32(t *testing.T) {
	shellsortHibbardSortOneThirdsTest[float32](t)
}

func TestShellSort_Hibbard_SortOneThirds_Float64(t *testing.T) {
	shellsortHibbardSortOneThirdsTest[float64](t)
}
func shellsortHibbardSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortTwoThirds()).SetGapCalcFunc(Hibbard)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Hibbard_SortTwoThirds_Int(t *testing.T) {
	shellsortHibbardSortTwoThirdsTest[int](t)
}

func TestShellSort_Hibbard_SortTwoThirds_Int32(t *testing.T) {
	shellsortHibbardSortTwoThirdsTest[int32](t)
}

func TestShellSort_Hibbard_SortTwoThirds_Int64(t *testing.T) {
	shellsortHibbardSortTwoThirdsTest[int64](t)
}

func TestShellSort_Hibbard_SortTwoThirds_Float32(t *testing.T) {
	shellsortHibbardSortTwoThirdsTest[float32](t)
}

func TestShellSort_Hibbard_SortTwoThirds_Float64(t *testing.T) {
	shellsortHibbardSortTwoThirdsTest[float64](t)
}
