package main

import (
	"fmt"
	"generics/interfaces"
)

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	sumIntSlice := SumSlice(intSlice)
	fmt.Printf("sumIntSlice: %v\n", sumIntSlice)

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sumFloatSlice := SumSlice(floatSlice)
	fmt.Printf("sumFloatSlice: %v\n", sumFloatSlice)

}

func SumSlice[T interfaces.Numeric](slice []T) T {

	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum

}
