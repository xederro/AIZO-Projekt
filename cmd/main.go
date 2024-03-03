package main

import (
	"AIZO-Projekt/algo"
	"AIZO-Projekt/algo/sort"
	"AIZO-Projekt/framework"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	//run tests for QuickSort with different types and pivot selection strategies

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int32 Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](algo.NewArray[int32](size).PopulateWithRandomValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](algo.NewArray[int32](size).PopulateWithAscendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](algo.NewArray[int32](size).PopulateWithDescendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int32 First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int32 Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int64 Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](algo.NewArray[int64](size).PopulateWithRandomValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Ascending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](algo.NewArray[int64](size).PopulateWithAscendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Descending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](algo.NewArray[int64](size).PopulateWithDescendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int64 First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Ascending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Descending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort Int64 Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Ascending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int64 Sorted Descending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int64](
						algo.NewArray[int64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[int64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Wait()

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float32 Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](algo.NewArray[float32](size).PopulateWithRandomValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Ascending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](algo.NewArray[float32](size).PopulateWithAscendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Descending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](algo.NewArray[float32](size).PopulateWithDescendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float32 First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Ascending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Descending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float32 Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Ascending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float32 Sorted Descending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float32](
						algo.NewArray[float32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float32]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float64 Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](algo.NewArray[float64](size).PopulateWithRandomValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Ascending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](algo.NewArray[float64](size).PopulateWithAscendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Descending Middle Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](algo.NewArray[float64](size).PopulateWithDescendingValues())
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float64 First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Ascending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Descending First Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Add(1)
	go framework.NewTimeTestHarness(10, 1000, 2000, 5000, 10000, 20000, 50000, 100000).
		AddTest(
			framework.NewTTO("QuickSort float64 Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Ascending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort float64 Sorted Descending Last Pivot").
				SetBefore(func(size int) any {
					return sort.NewQuickSort[float64](
						algo.NewArray[float64](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(*sort.QuickSort[float64]).Sort()
					return nil
				}),
		).
		ExecWG(&wg)

	wg.Wait()
}
