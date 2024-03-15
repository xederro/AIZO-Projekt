package main

import (
	"AIZO-Projekt/algo"
	"AIZO-Projekt/algo/sort"
	"AIZO-Projekt/algo/sort/heapsort"
	"AIZO-Projekt/algo/sort/insertionsort"
	"AIZO-Projekt/algo/sort/quicksort"
	"AIZO-Projekt/algo/sort/shellsort"
	"AIZO-Projekt/utils"
	"flag"
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
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
	tc := flag.Int("c", 0, "Count of tests")
	flag.Parse()
	ts := flag.Args()

	if *tc != 0 {
		auto(tc, ts)
	} else {
		manual()
	}
}

// auto is a function that allows user to choose type of values and file with values
func auto(tc *int, ts []string) {
	if *tc < 0 {
		log.Fatalln("Count of tests must be greater than 0")
	}
	if ts == nil || len(ts) <= 0 {
		log.Fatalln("No test values")
	}

	args, err := parseArgs(&ts)
	if err != nil {
		log.Fatalln("Invalid test values")
	}

	sort.Test(*tc, args)
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

// manual is a function that allows user to choose type of values and file with values
func manual() {
	var path string
	var typ int

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Insert path to file with values").
				Prompt(">").
				Value(&path),
		),
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose type").
				Options(
					huh.NewOption("int", INT).Selected(true),
					huh.NewOption("int32", INT32),
					huh.NewOption("int64", INT64),
					huh.NewOption("float32", FLOAT32),
					huh.NewOption("float64", FLOAT64),
				).
				Value(&typ),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}

	switch typ {
	case INT:
		arr, err := utils.ReadFile[int](path)
		if err != nil {
			log.Fatalln("Invalid file or path")
		}
		menu[int](arr)
		break
	case INT32:
		arr, err := utils.ReadFile[int32](path)
		if err != nil {
			log.Fatalln("Invalid file or path")
		}
		menu[int32](arr)
		break
	case INT64:
		arr, err := utils.ReadFile[int64](path)
		if err != nil {
			log.Fatalln("Invalid file or path")
		}
		menu[int64](arr)
		break
	case FLOAT32:
		arr, err := utils.ReadFile[float32](path)
		if err != nil {
			log.Fatalln("Invalid file or path")
		}
		menu[float32](arr)
		break
	case FLOAT64:
		arr, err := utils.ReadFile[float64](path)
		if err != nil {
			log.Fatalln("Invalid file or path")
		}
		menu[float64](arr)
		break
	default:
		log.Fatalln("Invalid type")
	}
}

// menu is a function that allows user to choose sorting algorithm
func menu[T algo.AllowedTypes](arr algo.Array[T]) {
	next := true
	for next {
		var a int
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[int]().
					Title("Choose an algorithm").
					Options(
						huh.NewOption("Heap Sort", HEAPSORT).Selected(true),
						huh.NewOption("Insertion Sort", INSERTIONSORT),
						huh.NewOption("Quick Sort", QUICKSORT),
						huh.NewOption("Shell Sort", SHELLSORT),
					).
					Value(&a),
			),
		).WithTheme(huh.ThemeCharm())

		err := form.Run()
		if err != nil {
			log.Fatalln("Error with form")
		}

		var s sort.Sort[T]
		switch a {
		case HEAPSORT:
			s = heapsort.NewHeapSort[T](arr)
			fmt.Println("\nHeap sort")
			break
		case INSERTIONSORT:
			s = insertionsort.NewInsertionSort[T](arr)
			fmt.Println("\nInsertion sort")
			break
		case QUICKSORT:
			s = quicksortWithPivotSelector[T](arr)
			fmt.Println("\nQuick sort")
			break
		case SHELLSORT:
			s = shellsortWithGapSelector[T](arr)
			fmt.Println("\nShell sort")
			break
		default:
			log.Fatalln("Invalid algorithm")
		}

		fmt.Println("Unsorted:")
		fmt.Println(arr)
		fmt.Println("Sorted: ")
		fmt.Println(s.Sort())

		form = huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Continue?").
					Value(&next),
			),
		).WithTheme(huh.ThemeCharm())

		err = form.Run()
		if err != nil {
			log.Fatalln("Error with form")
		}
	}
	fmt.Println("Bye!")
}

// quicksortWithPivotSelector is a function that allows user to choose pivot function for quick sort
func quicksortWithPivotSelector[T algo.AllowedTypes](arr algo.Array[T]) sort.Sort[T] {
	s := quicksort.NewQuickSort[T](arr)

	var pp int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose pivot point").
				Options(
					huh.NewOption("First", FIRST),
					huh.NewOption("Middle", MIDDLE).Selected(true),
					huh.NewOption("Last", LAST),
					huh.NewOption("Random", RANDOM),
				).
				Value(&pp),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}

	switch pp {
	case FIRST:
		s.SetPivotCalcFunc(quicksort.First)
		break
	case MIDDLE:
		s.SetPivotCalcFunc(quicksort.Middle)
		break
	case LAST:
		s.SetPivotCalcFunc(quicksort.Last)
		break
	case RANDOM:
		s.SetPivotCalcFunc(quicksort.Random)
		break
	default:
		log.Fatalln("Invalid pivot point")
	}
	return s
}

// shellsortWithGapSelector is a function that allows user to choose gap function for shell sort
func shellsortWithGapSelector[T algo.AllowedTypes](arr algo.Array[T]) sort.Sort[T] {
	s := shellsort.NewShellSort[T](arr)

	var pp int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Choose pivot point").
				Options(
					huh.NewOption("Shell's series", SHELL).Selected(true),
					huh.NewOption("Lazarus's series", LAZARUS),
				).
				Value(&pp),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		log.Fatalln("Error with form")
	}

	switch pp {
	case SHELL:
		s.SetGapCalcFunc(shellsort.Shell)
		break
	case LAZARUS:
		s.SetGapCalcFunc(shellsort.Lazarus)
		break
	default:
		log.Fatalln("Invalid gap function")
	}

	return s
}
