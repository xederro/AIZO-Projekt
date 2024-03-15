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

func shellsortLazarusRandomTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithRandomValues()).SetGapCalcFunc(Lazarus)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Lazarus_Random_Int(t *testing.T) {
	shellsortLazarusRandomTest[int](t)
}

func TestShellSort_Lazarus_Random_Int32(t *testing.T) {
	shellsortLazarusRandomTest[int32](t)
}

func TestShellSort_Lazarus_Random_Int64(t *testing.T) {
	shellsortLazarusRandomTest[int64](t)
}

func TestShellSort_Lazarus_Random_Float32(t *testing.T) {
	shellsortLazarusRandomTest[float32](t)
}

func TestShellSort_Lazarus_Random_Float64(t *testing.T) {
	shellsortLazarusRandomTest[float64](t)
}

func shellsortLazarusAscendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithAscendingValues()).SetGapCalcFunc(Lazarus)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Lazarus_Ascending_Int(t *testing.T) {
	shellsortLazarusAscendingTest[int](t)
}

func TestShellSort_Lazarus_Ascending_Int32(t *testing.T) {
	shellsortLazarusAscendingTest[int32](t)
}

func TestShellSort_Lazarus_Ascending_Int64(t *testing.T) {
	shellsortLazarusAscendingTest[int64](t)
}

func TestShellSort_Lazarus_Ascending_Float32(t *testing.T) {
	shellsortLazarusAscendingTest[float32](t)
}

func TestShellSort_Lazarus_Ascending_Float64(t *testing.T) {
	shellsortLazarusAscendingTest[float64](t)
}

func shellsortLazarusDescendingTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateWithDescendingValues()).SetGapCalcFunc(Lazarus)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Lazarus_Descending_Int(t *testing.T) {
	shellsortLazarusDescendingTest[int](t)
}

func TestShellSort_Lazarus_Descending_Int32(t *testing.T) {
	shellsortLazarusDescendingTest[int32](t)
}

func TestShellSort_Lazarus_Descending_Int64(t *testing.T) {
	shellsortLazarusDescendingTest[int64](t)
}

func TestShellSort_Lazarus_Descending_Float32(t *testing.T) {
	shellsortLazarusDescendingTest[float32](t)
}

func TestShellSort_Lazarus_Descending_Float64(t *testing.T) {
	shellsortLazarusDescendingTest[float64](t)
}

func shellsortLazarusSortOneThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortOneThirds()).SetGapCalcFunc(Lazarus)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Lazarus_SortOneThirds_Int(t *testing.T) {
	shellsortLazarusSortOneThirdsTest[int](t)
}

func TestShellSort_Lazarus_SortOneThirds_Int32(t *testing.T) {
	shellsortLazarusSortOneThirdsTest[int32](t)
}

func TestShellSort_Lazarus_SortOneThirds_Int64(t *testing.T) {
	shellsortLazarusSortOneThirdsTest[int64](t)
}

func TestShellSort_Lazarus_SortOneThirds_Float32(t *testing.T) {
	shellsortLazarusSortOneThirdsTest[float32](t)
}

func TestShellSort_Lazarus_SortOneThirds_Float64(t *testing.T) {
	shellsortLazarusSortOneThirdsTest[float64](t)
}
func shellsortLazarusSortTwoThirdsTest[T algo.AllowedTypes](t *testing.T) {
	for range 100 {
		h := NewShellSort(algo.NewArray[T](1000).PopulateAndSortTwoThirds()).SetGapCalcFunc(Lazarus)
		if !slices.IsSorted(h.Sort()) {
			t.Error("The array is not sorted")
		}
	}
}

func TestShellSort_Lazarus_SortTwoThirds_Int(t *testing.T) {
	shellsortLazarusSortTwoThirdsTest[int](t)
}

func TestShellSort_Lazarus_SortTwoThirds_Int32(t *testing.T) {
	shellsortLazarusSortTwoThirdsTest[int32](t)
}

func TestShellSort_Lazarus_SortTwoThirds_Int64(t *testing.T) {
	shellsortLazarusSortTwoThirdsTest[int64](t)
}

func TestShellSort_Lazarus_SortTwoThirds_Float32(t *testing.T) {
	shellsortLazarusSortTwoThirdsTest[float32](t)
}

func TestShellSort_Lazarus_SortTwoThirds_Float64(t *testing.T) {
	shellsortLazarusSortTwoThirdsTest[float64](t)
}
