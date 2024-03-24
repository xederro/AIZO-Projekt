package quicksort

import (
	"github.com/xederro/AIZO-Projekt/algo"
	"testing"
)

func quicksortRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewQuickSort(algo.NewArray[T](b.N).PopulateWithRandomValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkQuickSort_Random_Int(b *testing.B) {
	quicksortRandom[int](b)
}

func BenchmarkQuickSort_Random_Int32(b *testing.B) {
	quicksortRandom[int32](b)
}

func BenchmarkQuickSort_Random_Int64(b *testing.B) {
	quicksortRandom[int64](b)
}

func BenchmarkQuickSort_Random_Float32(b *testing.B) {
	quicksortRandom[float32](b)
}

func BenchmarkQuickSort_Random_Float64(b *testing.B) {
	quicksortRandom[float64](b)
}

func quicksortAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewQuickSort(algo.NewArray[T](b.N).PopulateWithAscendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkQuickSort_Ascending_Int(b *testing.B) {
	quicksortAscending[int](b)
}

func BenchmarkQuickSort_Ascending_Int32(b *testing.B) {
	quicksortAscending[int32](b)
}

func BenchmarkQuickSort_Ascending_Int64(b *testing.B) {
	quicksortAscending[int64](b)
}

func BenchmarkQuickSort_Ascending_Float32(b *testing.B) {
	quicksortAscending[float32](b)
}

func BenchmarkQuickSort_Ascending_Float64(b *testing.B) {
	quicksortAscending[float64](b)
}

func quicksortDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewQuickSort(algo.NewArray[T](b.N).PopulateWithDescendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkQuickSort_Descending_Int(b *testing.B) {
	quicksortDescending[int](b)
}

func BenchmarkQuickSort_Descending_Int32(b *testing.B) {
	quicksortDescending[int32](b)
}

func BenchmarkQuickSort_Descending_Int64(b *testing.B) {
	quicksortDescending[int64](b)
}

func BenchmarkQuickSort_Descending_Float32(b *testing.B) {
	quicksortDescending[float32](b)
}

func BenchmarkQuickSort_Descending_Float64(b *testing.B) {
	quicksortDescending[float64](b)
}

func quicksortSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewQuickSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkQuickSort_SortOneThirds_Int(b *testing.B) {
	quicksortSortOneThirds[int](b)
}

func BenchmarkQuickSort_SortOneThirds_Int32(b *testing.B) {
	quicksortSortOneThirds[int32](b)
}

func BenchmarkQuickSort_SortOneThirds_Int64(b *testing.B) {
	quicksortSortOneThirds[int64](b)
}

func BenchmarkQuickSort_SortOneThirds_Float32(b *testing.B) {
	quicksortSortOneThirds[float32](b)
}

func BenchmarkQuickSort_SortOneThirds_Float64(b *testing.B) {
	quicksortSortOneThirds[float64](b)
}
func quicksortSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewQuickSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkQuickSort_SortTwoThirds_Int(b *testing.B) {
	quicksortSortTwoThirds[int](b)
}

func BenchmarkQuickSort_SortTwoThirds_Int32(b *testing.B) {
	quicksortSortTwoThirds[int32](b)
}

func BenchmarkQuickSort_SortTwoThirds_Int64(b *testing.B) {
	quicksortSortTwoThirds[int64](b)
}

func BenchmarkQuickSort_SortTwoThirds_Float32(b *testing.B) {
	quicksortSortTwoThirds[float32](b)
}

func BenchmarkQuickSort_SortTwoThirds_Float64(b *testing.B) {
	quicksortSortTwoThirds[float64](b)
}
