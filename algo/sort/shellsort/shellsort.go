package shellsort

import (
	"AIZO-Projekt/algo"
	"math"
)

// ShellSort Description: This struct contains the copy of array.
type ShellSort[T algo.AllowedTypes] struct {
	Arr algo.Array[T]
	gen func(N, k int) int
}

// NewShellSort Description: This function creates a new ShellSort struct.
func NewShellSort[T algo.AllowedTypes](arr algo.Array[T]) ShellSort[T] {
	a := algo.NewArray[T](len(arr))
	copy(a, arr)
	return ShellSort[T]{
		Arr: a,
		gen: Shell,
	}
}

// SetGapCalcFunc Description: This function sets the pivotCalc function.
func (q ShellSort[T]) SetGapCalcFunc(gcf func(N, k int) int) ShellSort[T] {
	q.gen = gcf
	return q
}

// Sort Description: This function sorts the array using the insertion sort algorithm.
func (q ShellSort[T]) Sort() algo.Array[T] {
	//# Sort an array a[0...n-1].
	for k := 1; k < len(q.Arr); k++ {
		gap := q.gen(len(q.Arr), k)
		for i := gap; i < len(q.Arr); i++ {
			//# save a[i] in temp and make a hole at position i
			temp := q.Arr[i]
			//# shift earlier gap-sorted elements up until the correct location for a[i] is found
			var j int
			for j = i; j >= gap && q.Arr[j-gap] > temp; j -= gap {
				q.Arr[j] = q.Arr[j-gap]
			}
			//# put temp (the original a[i]) in its correct location
			q.Arr[j] = temp
		}
		if gap == 1 {
			break
		}
	}

	return q.Arr
}

func Shell(N, k int) int {
	g := float64(N) / math.Pow(2, float64(k))
	return int(g)
}

func Lazarus(N, k int) int {
	g := 2*math.Floor(float64(N)/math.Pow(2, float64(k+1))) + 1
	return int(g)
}
