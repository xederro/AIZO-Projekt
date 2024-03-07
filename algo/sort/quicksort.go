// Package sort Description: This package contains the implementation of the quicksort algorithm.
package sort

import "AIZO-Projekt/algo"

// QuickSort Description: This struct contains the copy of array.
type QuickSort[T algo.AllowedTypes] struct {
	Arr       algo.Array[T]
	pivotCalc func(lo int, hi int) int
}

// NewQuickSort Description: This function creates a new QuickSort struct.
func NewQuickSort[T algo.AllowedTypes](arr algo.Array[T]) QuickSort[T] {
	a := algo.NewArray[T](len(arr))
	copy(a, arr)
	return QuickSort[T]{
		Arr:       a,
		pivotCalc: Middle,
	}
}

// SetPivotCalcFunc Description: This function sets the pivotCalc function.
func (q QuickSort[T]) SetPivotCalcFunc(pcf func(lo int, hi int) int) QuickSort[T] {
	q.pivotCalc = pcf
	return q
}

// Sort Description: This function sorts the array using the quicksort algorithm.
func (q QuickSort[T]) Sort() algo.Array[T] {
	q.qs(0, len(q.Arr)-1)

	return q.Arr
}

// qs Description: This function is a helper function for the Sort function.
// It sorts the array using the quicksort recursive algorithm.
func (q QuickSort[T]) qs(lo, hi int) {
	if lo < hi {
		index := q.partition(lo, hi)
		q.qs(lo, index-1)
		q.qs(index+1, hi)
	}
}

// partition Description: This function is a helper function for the qs function.
// It returns the index of the pivot element after dividing the array into two parts.
func (q QuickSort[T]) partition(lo, hi int) int {
	//find the index of the pivot element, pivot and swap it with the last element to avoid checking it
	p := q.pivotCalc(lo, hi)
	pivot := q.Arr[p]
	index := lo
	q.Arr.Swap(p, hi)

	for i := lo; i < hi; i++ {
		if q.Arr[i] < pivot {
			q.Arr.Swap(i, index)
			index++
		}
	}
	q.Arr.Swap(hi, index)

	return index
}

// Last Description: This function returns the index of the last element.
func Last(lo int, hi int) int {
	return hi
}

// First Description: This function returns the index of the first element.
func First(lo int, hi int) int {
	return lo
}

// Middle Description: This function returns the index of the middle element.
func Middle(lo int, hi int) int {
	return (hi + lo) / 2
}
