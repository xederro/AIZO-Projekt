package sort

import (
	"AIZO-Projekt/algo"
	"AIZO-Projekt/algo/sort/heapsort"
	"AIZO-Projekt/algo/sort/insertionsort"
	"AIZO-Projekt/algo/sort/quicksort"
	"AIZO-Projekt/algo/sort/shellsort"
	"AIZO-Projekt/framework"
	"slices"
	"sync"
	"time"
)

func Test(testCount int, testSizes []int) {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	go TestHeapSort[int](wg, "int", testCount, testSizes...)
	go TestHeapSort[int32](wg, "int32", testCount, testSizes...)
	go TestHeapSort[int64](wg, "int64", testCount, testSizes...)
	go TestHeapSort[float32](wg, "float32", testCount, testSizes...)
	go TestHeapSort[float64](wg, "float64", testCount, testSizes...)
	wg.Wait()

	wg.Add(5)
	go TestInsertionSort[int](wg, "int", testCount, testSizes...)
	go TestInsertionSort[int32](wg, "int32", testCount, testSizes...)
	go TestInsertionSort[int64](wg, "int64", testCount, testSizes...)
	go TestInsertionSort[float32](wg, "float32", testCount, testSizes...)
	go TestInsertionSort[float64](wg, "float64", testCount, testSizes...)
	wg.Wait()

	wg.Add(5)
	go TestQuickSort[int](wg, "int", testCount, testSizes...)
	go TestQuickSort[int32](wg, "int32", testCount, testSizes...)
	go TestQuickSort[int64](wg, "int64", testCount, testSizes...)
	go TestQuickSort[float32](wg, "float32", testCount, testSizes...)
	go TestQuickSort[float64](wg, "float64", testCount, testSizes...)
	wg.Wait()

	wg.Add(5)
	go TestShellSort[int](wg, "int", testCount, testSizes...)
	go TestShellSort[int32](wg, "int32", testCount, testSizes...)
	go TestShellSort[int64](wg, "int64", testCount, testSizes...)
	go TestShellSort[float32](wg, "float32", testCount, testSizes...)
	go TestShellSort[float64](wg, "float64", testCount, testSizes...)
	wg.Wait()
}

func TestQuickSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Middle Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Middle Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Middle Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Random Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(quicksort.Random)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Random Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(quicksort.Random)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Random Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(quicksort.Random)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Random Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(quicksort.Random)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Random Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(quicksort.Random)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" First Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(quicksort.First)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending First Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(quicksort.First)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending First Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(quicksort.First)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(quicksort.First)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(quicksort.First)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Last Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(quicksort.Last)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Last Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(quicksort.Last)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Last Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(quicksort.Last)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(quicksort.Last)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return quicksort.NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(quicksort.Last)
				}).
				SetMeasure(func(data any) any {
					return data.(quicksort.QuickSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		ExecWG(wg)
}

func TestInsertionSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName, true).
				SetBefore(func(size int) any {
					return insertionsort.NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(insertionsort.InsertionSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted Ascending", true).
				SetBefore(func(size int) any {
					return insertionsort.NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(insertionsort.InsertionSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted Descending", true).
				SetBefore(func(size int) any {
					return insertionsort.NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(insertionsort.InsertionSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted 1 over 3", true).
				SetBefore(func(size int) any {
					return insertionsort.NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(insertionsort.InsertionSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted 2 over 3", true).
				SetBefore(func(size int) any {
					return insertionsort.NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(insertionsort.InsertionSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		ExecWG(wg)
}

func TestHeapSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("HeapSort "+typeName, true).
				SetBefore(func(size int) any {
					return heapsort.NewHeapSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(heapsort.HeapSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("HeapSort "+typeName+" Sorted Ascending", true).
				SetBefore(func(size int) any {
					return heapsort.NewHeapSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(heapsort.HeapSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("HeapSort "+typeName+" Sorted Descending", true).
				SetBefore(func(size int) any {
					return heapsort.NewHeapSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(heapsort.HeapSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("HeapSort "+typeName+" Sorted 1 over 3", true).
				SetBefore(func(size int) any {
					return heapsort.NewHeapSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(heapsort.HeapSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("HeapSort "+typeName+" Sorted 2 over 3", true).
				SetBefore(func(size int) any {
					return heapsort.NewHeapSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(heapsort.HeapSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		ExecWG(wg)
}

func TestShellSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("ShellSort Shell"+typeName, true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Shell"+typeName+" Sorted Ascending", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Shell"+typeName+" Sorted Descending", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Shell"+typeName+" Sorted 1 over 3", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Shell"+typeName+" Sorted 2 over 3", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Hibbard"+typeName, true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetGapCalcFunc(shellsort.Hibbard)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Hibbard"+typeName+" Sorted Ascending", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetGapCalcFunc(shellsort.Hibbard)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Hibbard"+typeName+" Sorted Descending", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetGapCalcFunc(shellsort.Hibbard)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Hibbard"+typeName+" Sorted 1 over 3", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetGapCalcFunc(shellsort.Hibbard)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		AddTest(
			framework.NewTTO("ShellSort Hibbard"+typeName+" Sorted 2 over 3", true).
				SetBefore(func(size int) any {
					return shellsort.NewShellSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetGapCalcFunc(shellsort.Hibbard)
				}).
				SetMeasure(func(data any) any {
					return data.(shellsort.ShellSort[T]).Sort()
				}).
				SetAfter(IsSorted[T]),
		).
		ExecWG(wg)
}

func IsSorted[T algo.AllowedTypes](name string, nr int, testSize int, time time.Duration, data any) {
	if !slices.IsSorted(data.(algo.Array[T])) {
		panic("Not sorted")
	}
	return
}
