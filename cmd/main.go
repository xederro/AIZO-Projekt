package main

import (
	"AIZO-Projekt/algo/sort"
	"AIZO-Projekt/framework"
	"sync"
)

var (
	testCount = 100
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

	framework.NewTimeTestHarness(1, 1).AddTest(
		framework.NewTTO("App", true).
			SetMeasure(func(data any) any {
				wg := sync.WaitGroup{}

				wg.Add(1)
				go sort.TestQuickSort[int32](&wg, "int32", testCount, testSizes...)

				wg.Add(1)
				go sort.TestQuickSort[int64](&wg, "int64", testCount, testSizes...)

				wg.Add(1)
				go sort.TestQuickSort[float32](&wg, "float32", testCount, testSizes...)

				wg.Add(1)
				go sort.TestQuickSort[float64](&wg, "float64", testCount, testSizes...)

				wg.Wait()
				return nil
			}),
	).Exec()
}
