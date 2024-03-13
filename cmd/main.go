package main

import "AIZO-Projekt/algo/sort"

var (
	testCount = 100
	testSizes = []int{1000, 2000, 4000, 8000, 16000, 32000, 64000}
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

	//sort.TestHeapSort[int32](nil, "int32", testCount, testSizes...)
	//sort.TestHeapSort[int64](nil, "int64", testCount, testSizes...)
	//sort.TestHeapSort[float32](nil, "float32", testCount, testSizes...)
	//sort.TestHeapSort[float64](nil, "float64", testCount, testSizes...)
	//sort.TestInsertionSort[int32](nil, "int32", testCount, testSizes...)
	//sort.TestInsertionSort[int64](nil, "int64", testCount, testSizes...)
	//sort.TestInsertionSort[float32](nil, "float32", testCount, testSizes...)
	//sort.TestInsertionSort[float64](nil, "float64", testCount, testSizes...)
	sort.TestQuickSort[int32](nil, "int32", testCount, testSizes...)
	//sort.TestQuickSort[int64](nil, "int64", testCount, testSizes...)
	//sort.TestQuickSort[float32](nil, "float32", testCount, testSizes...)
	//sort.TestQuickSort[float64](nil, "float64", testCount, testSizes...)
}
