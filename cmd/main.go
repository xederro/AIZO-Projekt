package main

import (
	"AIZO-Projekt/algo/sort"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	tc := flag.Int("c", 0, "Count of tests")
	flag.Parse()
	ts := flag.Args()

	if *tc != 0 {
		if *tc < 0 {
			panic("Count of tests must be greater than 0")
		}
		if ts == nil || len(ts) <= 0 {
			panic("No test values")
		}

		args, err := parseArgs(&ts)
		if err != nil {
			panic("Invalid test values")
		}

		sort.Test(*tc, args)
	} else {
		fmt.Println("No tests to run")
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
