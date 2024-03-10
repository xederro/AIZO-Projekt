package sort

import (
	"AIZO-Projekt/algo"
	"AIZO-Projekt/framework"
	"sync"
)

func TestQuickSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Middle Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Middle Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Middle Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Random Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(Random)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Random Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(Random)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Random Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(Random)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Random Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(Random)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Random Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(Random)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" First Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(First)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending First Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(First)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending First Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(First)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(First)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(First)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Last Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					).SetPivotCalcFunc(Last)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Ascending Last Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					).SetPivotCalcFunc(Last)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted Descending Last Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(Last)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 1 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(Last)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort "+typeName+" Sorted 2 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return NewQuickSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(Last)
				}).
				SetMeasure(func(data any) any {
					data.(QuickSort[T]).Sort()
					return nil
				}),
		).
		ExecWG(wg)
}

func TestInsertionSort[T algo.AllowedTypes](wg *sync.WaitGroup, typeName string, testCount int, testSizes ...int) {
	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName, true).
				SetBefore(func(size int) any {
					return NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(InsertionSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted Ascending", true).
				SetBefore(func(size int) any {
					return NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(InsertionSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted Descending", true).
				SetBefore(func(size int) any {
					return NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(InsertionSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted 1 over 3", true).
				SetBefore(func(size int) any {
					return NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(InsertionSort[T]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("InsertionSort "+typeName+" Sorted 2 over 3", true).
				SetBefore(func(size int) any {
					return NewInsertionSort[T](
						algo.NewArray[T](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(InsertionSort[T]).Sort()
					return nil
				}),
		).
		ExecWG(wg)
}
