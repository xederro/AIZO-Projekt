package shellsort

import (
	"github.com/xederro/AIZO-Projekt/algo"
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

func shellsortLazarusRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithRandomValues()).SetGapCalcFunc(Lazarus)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Lazarus_Random_Int(b *testing.B) {
	shellsortLazarusRandom[int](b)
}

func BenchmarkShellSort_Lazarus_Random_Int32(b *testing.B) {
	shellsortLazarusRandom[int32](b)
}

func BenchmarkShellSort_Lazarus_Random_Int64(b *testing.B) {
	shellsortLazarusRandom[int64](b)
}

func BenchmarkShellSort_Lazarus_Random_Float32(b *testing.B) {
	shellsortLazarusRandom[float32](b)
}

func BenchmarkShellSort_Lazarus_Random_Float64(b *testing.B) {
	shellsortLazarusRandom[float64](b)
}

func shellsortLazarusAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithAscendingValues()).SetGapCalcFunc(Lazarus)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Lazarus_Ascending_Int(b *testing.B) {
	shellsortLazarusAscending[int](b)
}

func BenchmarkShellSort_Lazarus_Ascending_Int32(b *testing.B) {
	shellsortLazarusAscending[int32](b)
}

func BenchmarkShellSort_Lazarus_Ascending_Int64(b *testing.B) {
	shellsortLazarusAscending[int64](b)
}

func BenchmarkShellSort_Lazarus_Ascending_Float32(b *testing.B) {
	shellsortLazarusAscending[float32](b)
}

func BenchmarkShellSort_Lazarus_Ascending_Float64(b *testing.B) {
	shellsortLazarusAscending[float64](b)
}

func shellsortLazarusDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateWithDescendingValues()).SetGapCalcFunc(Lazarus)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Lazarus_Descending_Int(b *testing.B) {
	shellsortLazarusDescending[int](b)
}

func BenchmarkShellSort_Lazarus_Descending_Int32(b *testing.B) {
	shellsortLazarusDescending[int32](b)
}

func BenchmarkShellSort_Lazarus_Descending_Int64(b *testing.B) {
	shellsortLazarusDescending[int64](b)
}

func BenchmarkShellSort_Lazarus_Descending_Float32(b *testing.B) {
	shellsortLazarusDescending[float32](b)
}

func BenchmarkShellSort_Lazarus_Descending_Float64(b *testing.B) {
	shellsortLazarusDescending[float64](b)
}

func shellsortLazarusSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds()).SetGapCalcFunc(Lazarus)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Lazarus_SortOneThirds_Int(b *testing.B) {
	shellsortLazarusSortOneThirds[int](b)
}

func BenchmarkShellSort_Lazarus_SortOneThirds_Int32(b *testing.B) {
	shellsortLazarusSortOneThirds[int32](b)
}

func BenchmarkShellSort_Lazarus_SortOneThirds_Int64(b *testing.B) {
	shellsortLazarusSortOneThirds[int64](b)
}

func BenchmarkShellSort_Lazarus_SortOneThirds_Float32(b *testing.B) {
	shellsortLazarusSortOneThirds[float32](b)
}

func BenchmarkShellSort_Lazarus_SortOneThirds_Float64(b *testing.B) {
	shellsortLazarusSortOneThirds[float64](b)
}
func shellsortLazarusSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewShellSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds()).SetGapCalcFunc(Lazarus)
	b.ResetTimer()
	h.Sort()
}

func BenchmarkShellSort_Lazarus_SortTwoThirds_Int(b *testing.B) {
	shellsortLazarusSortTwoThirds[int](b)
}

func BenchmarkShellSort_Lazarus_SortTwoThirds_Int32(b *testing.B) {
	shellsortLazarusSortTwoThirds[int32](b)
}

func BenchmarkShellSort_Lazarus_SortTwoThirds_Int64(b *testing.B) {
	shellsortLazarusSortTwoThirds[int64](b)
}

func BenchmarkShellSort_Lazarus_SortTwoThirds_Float32(b *testing.B) {
	shellsortLazarusSortTwoThirds[float32](b)
}

func BenchmarkShellSort_Lazarus_SortTwoThirds_Float64(b *testing.B) {
	shellsortLazarusSortTwoThirds[float64](b)
}
