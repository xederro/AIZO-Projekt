package main

import (
	"AIZO-Projekt/algo/sort"
	"AIZO-Projekt/algo/sort/insertionsort"
	"AIZO-Projekt/utils"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	tc := flag.Int("c", 0, "Count of tests")
	flag.Parse()
	ts := flag.Args()

	if *tc != 0 {
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
	} else {
		arr, err := utils.ReadFile[int](ts[0])
		if err != nil {
			log.Fatalln("Invalid file or path")
		}

		a := insertionsort.NewInsertionSort[int](arr).Sort()

		fmt.Println(arr)
		fmt.Println(a)
	}
}

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
