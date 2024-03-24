package insertionsort

import "github.com/xederro/AIZO-Projekt/algo"

// InsertionSort Description: This struct contains the copy of array.
type InsertionSort[T algo.AllowedTypes] struct {
	Arr algo.Array[T]
}

// NewInsertionSort Description: This function creates a new InsertionSort struct.
func NewInsertionSort[T algo.AllowedTypes](arr algo.Array[T]) InsertionSort[T] {
	a := algo.NewArray[T](len(arr))
	copy(a, arr)
	return InsertionSort[T]{
		Arr: a,
	}
}

// Sort Description: This function sorts the array using the insertion sort algorithm.
func (q InsertionSort[T]) Sort() algo.Array[T] {
	for i := 1; i < len(q.Arr); i++ {
		for j := i; j > 0 && q.Arr[j-1] > q.Arr[j]; j-- {
			q.Arr.Swap(j, j-1)
		}
	}

	return q.Arr
}
