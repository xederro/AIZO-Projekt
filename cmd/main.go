package main

import (
	"AIZO-Projekt/framework"
	"sync"
)

var (
	testCount = 100
	testSizes = []int{10000, 20000, 40000, 80000, 160000, 320000, 640000}
)

func main() {
	// run tests for QuickSort with set data
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

	framework.NewTimeTestHarness(1, 1).AddTest(
		framework.NewTTO("App", true).
			SetMeasure(func(data any) any {
				wg := sync.WaitGroup{}

				//wg.Add(1)
				//go sort.TestHeapSort[int32](&wg, "int32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestHeapSort[int64](&wg, "int64", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestHeapSort[float32](&wg, "float32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestHeapSort[float64](&wg, "float64", testCount, testSizes...)

				//wg.Add(1)
				//go sort.TestInsertionSort[int32](&wg, "int32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestInsertionSort[int64](&wg, "int64", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestInsertionSort[float32](&wg, "float32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestInsertionSort[float64](&wg, "float64", testCount, testSizes...)

				//wg.Add(1)
				//go sort.TestQuickSort[int32](&wg, "int32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestQuickSort[int64](&wg, "int64", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestQuickSort[float32](&wg, "float32", testCount, testSizes...)
				//
				//wg.Add(1)
				//go sort.TestQuickSort[float64](&wg, "float64", testCount, testSizes...)
				//
				wg.Wait()
				return nil
			}),
	).Exec()
}
