package main

import (
	"fmt"
)

func sumSlice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func isSliceEmpty[T any](s []T) bool {
	return len(s) == 0
}

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isSliceEmpty[int](intSlice))

	var floatSlice = []float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sumSlice(floatSlice))
	fmt.Println(isSliceEmpty(floatSlice))

}
