package heapsort

import "AIZO-Projekt/algo"

// HeapSort Description: This struct contains the copy of array.
type HeapSort[T algo.AllowedTypes] struct {
	Heap algo.Heap[T]
}

// NewHeapSort Description: This function creates a new InsertionSort struct.
func NewHeapSort[T algo.AllowedTypes](arr algo.Array[T]) HeapSort[T] {
	a := algo.NewArray[T](len(arr))
	copy(a, arr)
	return HeapSort[T]{
		Heap: algo.NewHeap(a),
	}
}

// Sort Description: This function sorts the array using the heap sort algorithm.
func (q HeapSort[T]) Sort() algo.Array[T] {
	q.Heap.BuildHeap()
	q.heapSort()
	return algo.Array[T](q.Heap)
}

// heapSort Description: The heap sort algorithm.
func (q HeapSort[T]) heapSort() {
	for i := len(q.Heap); i > 0; i-- {
		q.Heap.Swap(0, i-1)
		q.Heap.ShiftDown(i-1, 0)
	}
}
