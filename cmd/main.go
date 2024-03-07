package main

import (
	"AIZO-Projekt/algo"
	"AIZO-Projekt/algo/sort"
	"AIZO-Projekt/framework"
)

var (
	testCount = 10
	testSizes = []int{1000, 2000, 4000, 8000, 16000, 32000, 64000}
)

func main() {
	// run tests for QuickSort with set data
	//test := []int32{12, 13, 0, 14, 15, 16, 17, 3, 2, 1, 4, 5, 6, 7, 8, 9, 10, 11, 18, 19, 20, 21, 22, 23, 24}
	//framework.NewTimeTestHarness(1).
	//	AddTest(
	//		framework.NewTTO("Custom test", false).
	//			SetBefore(func(size int) any {
	//				return sort.NewQuickSort[int32](
	//					test,
	//				).SetPivotCalcFunc(sort.Middle)
	//			}).
	//			SetMeasure(func(data any) any {
	//				return data.(*sort.QuickSort[int32]).Sort()
	//			}).
	//			SetAfter(func(name string, nr int, testSize int, time time.Duration, data any) {
	//				fmt.Println(time)
	//				fmt.Println(test)
	//				fmt.Println(data)
	//			}),
	//	).
	//	Exec()

	framework.NewTimeTestHarness(testCount, testSizes...).
		AddTest(
			framework.NewTTO("QuickSort Int32 Middle Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithRandomValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending Middle Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithAscendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending Middle Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 1 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortOneThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 2 over 3 Middle Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortTwoThirds(),
					)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 First Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending First Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending First Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 1 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 2 over 3 First Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(sort.First)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Last Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Ascending Last Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted Descending Last Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateWithDescendingValues(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 1 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortOneThirds(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		AddTest(
			framework.NewTTO("QuickSort Int32 Sorted 2 over 3 Last Pivot", true).
				SetBefore(func(size int) any {
					return sort.NewQuickSort[int32](
						algo.NewArray[int32](size).
							PopulateAndSortTwoThirds(),
					).SetPivotCalcFunc(sort.Last)
				}).
				SetMeasure(func(data any) any {
					data.(sort.QuickSort[int32]).Sort()
					return nil
				}),
		).
		Exec()

	//for range 1000 {
	//	arr := algo.NewArray[int32](1000).PopulateWithRandomValues()
	//	test := sort.NewQuickSort[int32](arr)
	//	start := time.Now()
	//	test.Sort()
	//	dur := time.Since(start)
	//	fmt.Println(dur.Nanoseconds(), slices.IsSorted(test.Arr))
	//}
}
