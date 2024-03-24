package main

import (
	"flag"
	"github.com/xederro/AIZO-Projekt/algo/sort"
	"os"
	"runtime/pprof"
	"strconv"
)

const (
	// Constants for types of values
	INT = iota
	INT32
	INT64
	FLOAT32
	FLOAT64
	// Constants for types of sorts
	HEAPSORT
	INSERTIONSORT
	QUICKSORT
	SHELLSORT
	// Constants for pivot point positions
	RANDOM
	MIDDLE
	FIRST
	LAST
	// Constants for series of gaps
	SHELL
	LAZARUS
)

func main() {
	perf := flag.Bool("p", false, "Measure performance")
	testConf := sort.TestConfig{
		TestCount:         flag.Int("c", 0, "Count of tests"),
		TestHeapSort:      flag.Bool("h", false, "Test heap sort"),
		TestInsertionSort: flag.Bool("i", false, "Test Insertion sort"),
		TestQuickSort:     flag.Bool("q", false, "Test Quick sort"),
		TestShellSort:     flag.Bool("s", false, "Test Shell sort"),
		TestAsync:         flag.Bool("a", false, "Test Asynchronously"),
	}

	flag.Parse()
	ts := flag.Args()

	if *perf {
		measureCPUPerf()
	}

	if *testConf.TestCount != 0 {
		auto(&testConf, ts)
	} else {
		manual()
	}
}

// parseArgs is a function that parses string arguments to int
func parseArgs(args *[]string) ([]int, error) {
	var result []int
	for _, arg := range *args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// measureCPUPerf is a function that measures CPU performance
func measureCPUPerf() {
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
}
