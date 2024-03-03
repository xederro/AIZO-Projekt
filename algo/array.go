// Package algo Description: This file contains the implementation of the Array helper type for algorithms and its methods
package algo

import (
	"log"
	"math"
	"math/rand/v2"
	"reflect"
	"slices"
)

// AllowedTypes is an interface that allows to use only int64, int32, float32, float64 types
type AllowedTypes interface {
	int64 | int32 | float64 | float32
}

// Array is a generic type that allows to create an array of AllowedTypes
type Array[T AllowedTypes] []T

// NewArray is a constructor for the Array type
func NewArray[T AllowedTypes](n int) Array[T] {
	return make([]T, n)
}

// PopulateWithRandomValues is a method that populates the array with random values
func (arr Array[T]) PopulateWithRandomValues() Array[T] {
	if arr == nil || len(arr) == 0 {
		log.Fatalln("Provided empty array or invalid pointer")
	}

	t := reflect.ValueOf((arr)[0])
	switch t.Kind() {
	case reflect.Int64:
		for i := 0; i < len(arr); i++ {
			arr[i] = T(rand.Int64N(255))
		}
		break
	case reflect.Int32:
		for i := 0; i < len(arr); i++ {
			arr[i] = T(rand.Int32())
		}
		break
	case reflect.Float64:
		for i := 0; i < len(arr); i++ {
			arr[i] = T(rand.Float64() * math.MaxFloat64)
		}
		break
	case reflect.Float32:
		for i := 0; i < len(arr); i++ {
			arr[i] = T(rand.Float32() * math.MaxFloat32)
		}
		break
	default:
		log.Fatalln("Provided array of wrong type")
	}

	return arr
}

// PopulateWithAscendingValues is a method that populates the array with ascending values
func (arr Array[T]) PopulateWithAscendingValues() Array[T] {
	arr.PopulateWithRandomValues()
	slices.Sort(arr)

	return arr
}

// PopulateWithDescendingValues is a method that populates the array with descending values
func (arr Array[T]) PopulateWithDescendingValues() Array[T] {
	arr.PopulateWithRandomValues()
	slices.Sort(arr)
	slices.Reverse(arr)

	return arr
}

// Swap is a method that swaps two elements in the array
// p1 and p2 are the positions of the elements to be swapped
func (arr Array[T]) Swap(p1, p2 int) Array[T] {
	if p1 != p2 {
		tmp := arr[p1]
		arr[p1] = arr[p2]
		arr[p2] = tmp
	}

	return arr
}
