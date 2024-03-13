package shellsort

import (
	"AIZO-Projekt/algo"
	"testing"
)

func shellsortShellRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithRandomValues()).SetGapCalcFunc(Shell)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Shell_Random_Int(b *testing.B) {
	shellsortShellRandom[int](b)
}

func BenchmarkShellSort_Shell_Random_Int32(b *testing.B) {
	shellsortShellRandom[int32](b)
}

func BenchmarkShellSort_Shell_Random_Int64(b *testing.B) {
	shellsortShellRandom[int64](b)
}

func BenchmarkShellSort_Shell_Random_Float32(b *testing.B) {
	shellsortShellRandom[float32](b)
}

func BenchmarkShellSort_Shell_Random_Float64(b *testing.B) {
	shellsortShellRandom[float64](b)
}

func shellsortShellAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithAscendingValues()).SetGapCalcFunc(Shell)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Shell_Ascending_Int(b *testing.B) {
	shellsortShellAscending[int](b)
}

func BenchmarkShellSort_Shell_Ascending_Int32(b *testing.B) {
	shellsortShellAscending[int32](b)
}

func BenchmarkShellSort_Shell_Ascending_Int64(b *testing.B) {
	shellsortShellAscending[int64](b)
}

func BenchmarkShellSort_Shell_Ascending_Float32(b *testing.B) {
	shellsortShellAscending[float32](b)
}

func BenchmarkShellSort_Shell_Ascending_Float64(b *testing.B) {
	shellsortShellAscending[float64](b)
}

func shellsortShellDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithDescendingValues()).SetGapCalcFunc(Shell)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Shell_Descending_Int(b *testing.B) {
	shellsortShellDescending[int](b)
}

func BenchmarkShellSort_Shell_Descending_Int32(b *testing.B) {
	shellsortShellDescending[int32](b)
}

func BenchmarkShellSort_Shell_Descending_Int64(b *testing.B) {
	shellsortShellDescending[int64](b)
}

func BenchmarkShellSort_Shell_Descending_Float32(b *testing.B) {
	shellsortShellDescending[float32](b)
}

func BenchmarkShellSort_Shell_Descending_Float64(b *testing.B) {
	shellsortShellDescending[float64](b)
}

func shellsortShellSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds()).SetGapCalcFunc(Shell)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Shell_SortOneThirds_Int(b *testing.B) {
	shellsortShellSortOneThirds[int](b)
}

func BenchmarkShellSort_Shell_SortOneThirds_Int32(b *testing.B) {
	shellsortShellSortOneThirds[int32](b)
}

func BenchmarkShellSort_Shell_SortOneThirds_Int64(b *testing.B) {
	shellsortShellSortOneThirds[int64](b)
}

func BenchmarkShellSort_Shell_SortOneThirds_Float32(b *testing.B) {
	shellsortShellSortOneThirds[float32](b)
}

func BenchmarkShellSort_Shell_SortOneThirds_Float64(b *testing.B) {
	shellsortShellSortOneThirds[float64](b)
}
func shellsortShellSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds()).SetGapCalcFunc(Shell)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Shell_SortTwoThirds_Int(b *testing.B) {
	shellsortShellSortTwoThirds[int](b)
}

func BenchmarkShellSort_Shell_SortTwoThirds_Int32(b *testing.B) {
	shellsortShellSortTwoThirds[int32](b)
}

func BenchmarkShellSort_Shell_SortTwoThirds_Int64(b *testing.B) {
	shellsortShellSortTwoThirds[int64](b)
}

func BenchmarkShellSort_Shell_SortTwoThirds_Float32(b *testing.B) {
	shellsortShellSortTwoThirds[float32](b)
}

func BenchmarkShellSort_Shell_SortTwoThirds_Float64(b *testing.B) {
	shellsortShellSortTwoThirds[float64](b)
}

func shellsortHibbardRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithRandomValues()).SetGapCalcFunc(Hibbard)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Hibbard_Random_Int(b *testing.B) {
	shellsortHibbardRandom[int](b)
}

func BenchmarkShellSort_Hibbard_Random_Int32(b *testing.B) {
	shellsortHibbardRandom[int32](b)
}

func BenchmarkShellSort_Hibbard_Random_Int64(b *testing.B) {
	shellsortHibbardRandom[int64](b)
}

func BenchmarkShellSort_Hibbard_Random_Float32(b *testing.B) {
	shellsortHibbardRandom[float32](b)
}

func BenchmarkShellSort_Hibbard_Random_Float64(b *testing.B) {
	shellsortHibbardRandom[float64](b)
}

func shellsortHibbardAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithAscendingValues()).SetGapCalcFunc(Hibbard)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Hibbard_Ascending_Int(b *testing.B) {
	shellsortHibbardAscending[int](b)
}

func BenchmarkShellSort_Hibbard_Ascending_Int32(b *testing.B) {
	shellsortHibbardAscending[int32](b)
}

func BenchmarkShellSort_Hibbard_Ascending_Int64(b *testing.B) {
	shellsortHibbardAscending[int64](b)
}

func BenchmarkShellSort_Hibbard_Ascending_Float32(b *testing.B) {
	shellsortHibbardAscending[float32](b)
}

func BenchmarkShellSort_Hibbard_Ascending_Float64(b *testing.B) {
	shellsortHibbardAscending[float64](b)
}

func shellsortHibbardDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithDescendingValues()).SetGapCalcFunc(Hibbard)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Hibbard_Descending_Int(b *testing.B) {
	shellsortHibbardDescending[int](b)
}

func BenchmarkShellSort_Hibbard_Descending_Int32(b *testing.B) {
	shellsortHibbardDescending[int32](b)
}

func BenchmarkShellSort_Hibbard_Descending_Int64(b *testing.B) {
	shellsortHibbardDescending[int64](b)
}

func BenchmarkShellSort_Hibbard_Descending_Float32(b *testing.B) {
	shellsortHibbardDescending[float32](b)
}

func BenchmarkShellSort_Hibbard_Descending_Float64(b *testing.B) {
	shellsortHibbardDescending[float64](b)
}

func shellsortHibbardSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds()).SetGapCalcFunc(Hibbard)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Hibbard_SortOneThirds_Int(b *testing.B) {
	shellsortHibbardSortOneThirds[int](b)
}

func BenchmarkShellSort_Hibbard_SortOneThirds_Int32(b *testing.B) {
	shellsortHibbardSortOneThirds[int32](b)
}

func BenchmarkShellSort_Hibbard_SortOneThirds_Int64(b *testing.B) {
	shellsortHibbardSortOneThirds[int64](b)
}

func BenchmarkShellSort_Hibbard_SortOneThirds_Float32(b *testing.B) {
	shellsortHibbardSortOneThirds[float32](b)
}

func BenchmarkShellSort_Hibbard_SortOneThirds_Float64(b *testing.B) {
	shellsortHibbardSortOneThirds[float64](b)
}
func shellsortHibbardSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds()).SetGapCalcFunc(Hibbard)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Hibbard_SortTwoThirds_Int(b *testing.B) {
	shellsortHibbardSortTwoThirds[int](b)
}

func BenchmarkShellSort_Hibbard_SortTwoThirds_Int32(b *testing.B) {
	shellsortHibbardSortTwoThirds[int32](b)
}

func BenchmarkShellSort_Hibbard_SortTwoThirds_Int64(b *testing.B) {
	shellsortHibbardSortTwoThirds[int64](b)
}

func BenchmarkShellSort_Hibbard_SortTwoThirds_Float32(b *testing.B) {
	shellsortHibbardSortTwoThirds[float32](b)
}

func BenchmarkShellSort_Hibbard_SortTwoThirds_Float64(b *testing.B) {
	shellsortHibbardSortTwoThirds[float64](b)
}
