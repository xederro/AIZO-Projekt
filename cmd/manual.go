package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/xederro/AIZO-Projekt/algo"
	"github.com/xederro/AIZO-Projekt/algo/sort"
	"github.com/xederro/AIZO-Projekt/algo/sort/heapsort"
	"github.com/xederro/AIZO-Projekt/algo/sort/insertionsort"
	"github.com/xederro/AIZO-Projekt/algo/sort/quicksort"
	"github.com/xederro/AIZO-Projekt/algo/sort/shellsort"
	"github.com/xederro/AIZO-Projekt/utils"
	"log"
)

// manual is a function that allows user to choose type of values and file with values
func manual() {
	path, typ := getFileAndType()

	readVal(typ, path)
}

// readVal is a function that reads values from file
func readVal(typ int, path string) {
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

// getFileAndType is a function that allows user to choose type of values and file with values
func getFileAndType() (string, int) {
	var path string
	var typ int

	// insert path to file with values, and choose type of values
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
	return path, typ
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

		s := chooseSortMethod(arr, a)

		fmt.Printf("Unsorted:\n%v\nSorted:\n%v\n", arr, s.Sort())
		//fmt.Println(arr)
		//fmt.Println("Sorted: ")
		//fmt.Println(s.Sort())

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

// chooseSortMethod is a function that allows user to choose sorting algorithm
func chooseSortMethod[T algo.AllowedTypes](arr algo.Array[T], a int) sort.Sort[T] {
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
	return s
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
