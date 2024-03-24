package sort

import "github.com/xederro/AIZO-Projekt/algo"

// Sort is an interface for sorting algorithms
type Sort[T algo.AllowedTypes] interface {
	// Sort sorts the array
	Sort() algo.Array[T]
}
