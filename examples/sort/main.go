package main

import (
	"fmt"
	"generics/interfaces"
	"sort"
)

func main() {
	intSlice := []uint32{5, 3, 1, 4, 2}

	ascendingSorted := SortSlice(intSlice, true)
	fmt.Println("Ascending Sort:", ascendingSorted)

	descendingSorted := SortSlice(intSlice, false)
	fmt.Println("Descending Sort:", descendingSorted)
}

func SortSlice[T interfaces.Numeric](slice []T, ascending bool) []T {

	if ascending {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
	} else {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})
	}
	return slice

}
