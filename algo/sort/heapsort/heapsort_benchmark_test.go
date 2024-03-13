package heapsort

import (
	"AIZO-Projekt/algo"
	"testing"
)

func heapsortRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewHeapSort(algo.NewArray[T](b.N).PopulateWithRandomValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkHeapSort_Random_Int(b *testing.B) {
	heapsortRandom[int](b)
}

func BenchmarkHeapSort_Random_Int32(b *testing.B) {
	heapsortRandom[int32](b)
}

func BenchmarkHeapSort_Random_Int64(b *testing.B) {
	heapsortRandom[int64](b)
}

func BenchmarkHeapSort_Random_Float32(b *testing.B) {
	heapsortRandom[float32](b)
}

func BenchmarkHeapSort_Random_Float64(b *testing.B) {
	heapsortRandom[float64](b)
}

func heapsortAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewHeapSort(algo.NewArray[T](b.N).PopulateWithAscendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkHeapSort_Ascending_Int(b *testing.B) {
	heapsortAscending[int](b)
}

func BenchmarkHeapSort_Ascending_Int32(b *testing.B) {
	heapsortAscending[int32](b)
}

func BenchmarkHeapSort_Ascending_Int64(b *testing.B) {
	heapsortAscending[int64](b)
}

func BenchmarkHeapSort_Ascending_Float32(b *testing.B) {
	heapsortAscending[float32](b)
}

func BenchmarkHeapSort_Ascending_Float64(b *testing.B) {
	heapsortAscending[float64](b)
}

func heapsortDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewHeapSort(algo.NewArray[T](b.N).PopulateWithDescendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkHeapSort_Descending_Int(b *testing.B) {
	heapsortDescending[int](b)
}

func BenchmarkHeapSort_Descending_Int32(b *testing.B) {
	heapsortDescending[int32](b)
}

func BenchmarkHeapSort_Descending_Int64(b *testing.B) {
	heapsortDescending[int64](b)
}

func BenchmarkHeapSort_Descending_Float32(b *testing.B) {
	heapsortDescending[float32](b)
}

func BenchmarkHeapSort_Descending_Float64(b *testing.B) {
	heapsortDescending[float64](b)
}

func heapsortSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewHeapSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkHeapSort_SortOneThirds_Int(b *testing.B) {
	heapsortSortOneThirds[int](b)
}

func BenchmarkHeapSort_SortOneThirds_Int32(b *testing.B) {
	heapsortSortOneThirds[int32](b)
}

func BenchmarkHeapSort_SortOneThirds_Int64(b *testing.B) {
	heapsortSortOneThirds[int64](b)
}

func BenchmarkHeapSort_SortOneThirds_Float32(b *testing.B) {
	heapsortSortOneThirds[float32](b)
}

func BenchmarkHeapSort_SortOneThirds_Float64(b *testing.B) {
	heapsortSortOneThirds[float64](b)
}
func heapsortSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewHeapSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkHeapSort_SortTwoThirds_Int(b *testing.B) {
	heapsortSortTwoThirds[int](b)
}

func BenchmarkHeapSort_SortTwoThirds_Int32(b *testing.B) {
	heapsortSortTwoThirds[int32](b)
}

func BenchmarkHeapSort_SortTwoThirds_Int64(b *testing.B) {
	heapsortSortTwoThirds[int64](b)
}

func BenchmarkHeapSort_SortTwoThirds_Float32(b *testing.B) {
	heapsortSortTwoThirds[float32](b)
}

func BenchmarkHeapSort_SortTwoThirds_Float64(b *testing.B) {
	heapsortSortTwoThirds[float64](b)
}
