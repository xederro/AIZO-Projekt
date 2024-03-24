package main

import (
	"github.com/xederro/AIZO-Projekt/algo/sort"
	"log"
)

// auto is a function that allows user to choose type of values and file with values
func auto(tc *sort.TestConfig, ts []string) {
	if *tc.TestCount < 0 {
		log.Fatalln("Count of tests must be greater than 0")
	}
	if ts == nil || len(ts) <= 0 {
		log.Fatalln("No test values")
	}

	args, err := parseArgs(&ts)
	if err != nil {
		log.Fatalln("Invalid test values")
	}

	tc.TestSizes = &args
	sort.Test(tc)
}
