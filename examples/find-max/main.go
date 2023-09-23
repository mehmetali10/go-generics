package main

import (
	"fmt"
	"generics/interfaces"
)

func main() {
	intSlice := []uint8{1, 2, 3, 4, 5}
	maxInt := FindMax(intSlice)
	fmt.Printf("maxInt: %v\n", maxInt)

	floatSlice := []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	maxFloat := FindMax(floatSlice)
	fmt.Printf("maxFloat: %v\n", maxFloat)
}

func FindMax[T interfaces.Numeric](slice []T) T {

	var max T
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max

}
