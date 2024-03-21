package quicksort

import (
	"AIZO-Projekt/algo"
	"math/rand"
)

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
	stack := algo.NewArray[int](hi - lo + 2)

	top := -1

	top++
	stack[top] = lo
	top++
	stack[top] = hi

	// Keep popping from stack while is not empty
	for top >= 0 {
		hi = stack[top]
		top--
		lo = stack[top]
		top--
		if lo >= hi {
			return
		}
		p := q.partition(lo, hi)

		// If there are elements on left side of pivot,
		// then push left side to stack
		if p > lo {
			top++
			stack[top] = lo
			top++
			stack[top] = p
		}

		// If there are elements on right side of pivot,
		// then push right side to stack
		if p+1 < hi {
			top++
			stack[top] = p + 1
			top++
			stack[top] = hi
		}
	}

	//if lo >= hi {
	//	return
	//}
	//m := q.partition(lo, hi)
	//q.qs(lo, m)
	//q.qs(m+1, hi)
}

// // partition Description: This function is a helper function for the qs function.
// // It returns the index of the pivot element after dividing the array into two parts.
//
//	func (q QuickSort[T]) partition(lo, hi int) int {
//		//find the index of the pivot element, pivot and swap it with the last element to avoid checking it
//		p := q.pivotCalc(lo, hi)
//		pivot := q.Arr[p]
//		index := lo
//		q.Arr.Swap(p, hi)
//
//		for i := lo; i < hi; i++ {
//			if q.Arr[i] < pivot {
//				q.Arr.Swap(i, index)
//				index++
//			}
//		}
//		q.Arr.Swap(hi, index)
//
//		return index
//	}
//
// partition Description: This function is a helper function for the qs function.
// It returns the index of the pivot element after dividing the array into two parts.
func (q QuickSort[T]) partition(lo, hi int) int {
	p := q.pivotCalc(lo, hi)
	pivot := q.Arr[p]
	l := lo
	r := hi
	for {
		for q.Arr[l] < pivot {
			l++
		}
		for q.Arr[r] > pivot {
			r--
		}

		if l < r {
			q.Arr.Swap(l, r)
			l++
			r--
		} else {
			if r == hi {
				r--
			}
			return r
		}
	}
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

// Random Description: This function returns the index of the middle element.
func Random(lo int, hi int) int {
	return rand.Intn(hi-lo) + lo
}
