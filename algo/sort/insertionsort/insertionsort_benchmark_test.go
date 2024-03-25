package insertionsort

import (
	"github.com/xederro/AIZO-Projekt/algo"
	"testing"
)

func insertionsortRandom[T algo.AllowedTypes](b *testing.B) {
	h := NewInsertionSort(algo.NewArray[T](b.N).PopulateWithRandomValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkInsertionSort_Random_Int(b *testing.B) {
	insertionsortRandom[int8](b)
}

func BenchmarkInsertionSort_Random_Int32(b *testing.B) {
	insertionsortRandom[int32](b)
}

func BenchmarkInsertionSort_Random_Int64(b *testing.B) {
	insertionsortRandom[int64](b)
}

func BenchmarkInsertionSort_Random_Float32(b *testing.B) {
	insertionsortRandom[float32](b)
}

func BenchmarkInsertionSort_Random_Float64(b *testing.B) {
	insertionsortRandom[float64](b)
}

func insertionsortAscending[T algo.AllowedTypes](b *testing.B) {
	h := NewInsertionSort(algo.NewArray[T](b.N).PopulateWithAscendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkInsertionSort_Ascending_Int(b *testing.B) {
	insertionsortAscending[int8](b)
}

func BenchmarkInsertionSort_Ascending_Int32(b *testing.B) {
	insertionsortAscending[int32](b)
}

func BenchmarkInsertionSort_Ascending_Int64(b *testing.B) {
	insertionsortAscending[int64](b)
}

func BenchmarkInsertionSort_Ascending_Float32(b *testing.B) {
	insertionsortAscending[float32](b)
}

func BenchmarkInsertionSort_Ascending_Float64(b *testing.B) {
	insertionsortAscending[float64](b)
}

func insertionsortDescending[T algo.AllowedTypes](b *testing.B) {
	h := NewInsertionSort(algo.NewArray[T](b.N).PopulateWithDescendingValues())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkInsertionSort_Descending_Int(b *testing.B) {
	insertionsortDescending[int8](b)
}

func BenchmarkInsertionSort_Descending_Int32(b *testing.B) {
	insertionsortDescending[int32](b)
}

func BenchmarkInsertionSort_Descending_Int64(b *testing.B) {
	insertionsortDescending[int64](b)
}

func BenchmarkInsertionSort_Descending_Float32(b *testing.B) {
	insertionsortDescending[float32](b)
}

func BenchmarkInsertionSort_Descending_Float64(b *testing.B) {
	insertionsortDescending[float64](b)
}

func insertionsortSortOneThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewInsertionSort(algo.NewArray[T](b.N).PopulateAndSortOneThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkInsertionSort_SortOneThirds_Int(b *testing.B) {
	insertionsortSortOneThirds[int8](b)
}

func BenchmarkInsertionSort_SortOneThirds_Int32(b *testing.B) {
	insertionsortSortOneThirds[int32](b)
}

func BenchmarkInsertionSort_SortOneThirds_Int64(b *testing.B) {
	insertionsortSortOneThirds[int64](b)
}

func BenchmarkInsertionSort_SortOneThirds_Float32(b *testing.B) {
	insertionsortSortOneThirds[float32](b)
}

func BenchmarkInsertionSort_SortOneThirds_Float64(b *testing.B) {
	insertionsortSortOneThirds[float64](b)
}
func insertionsortSortTwoThirds[T algo.AllowedTypes](b *testing.B) {
	h := NewInsertionSort(algo.NewArray[T](b.N).PopulateAndSortTwoThirds())
	b.ResetTimer()
	h.Sort()
}

func BenchmarkInsertionSort_SortTwoThirds_Int(b *testing.B) {
	insertionsortSortTwoThirds[int8](b)
}

func BenchmarkInsertionSort_SortTwoThirds_Int32(b *testing.B) {
	insertionsortSortTwoThirds[int32](b)
}

func BenchmarkInsertionSort_SortTwoThirds_Int64(b *testing.B) {
	insertionsortSortTwoThirds[int64](b)
}

func BenchmarkInsertionSort_SortTwoThirds_Float32(b *testing.B) {
	insertionsortSortTwoThirds[float32](b)
}

func BenchmarkInsertionSort_SortTwoThirds_Float64(b *testing.B) {
	insertionsortSortTwoThirds[float64](b)
}
