package main

import (
	"fmt"
	"time"
)

var matrix = [][]int{
	{1, 2, 10, 3, 49, 5, 22, 1, -2, -3, 4, 6},
	{5, 3, 5, 3, 5, -1, 0, 0, 0, 11, -100, 0},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{44, 23, -1, 3, -1, 32, -4, -35, 3, 5, 10, 45},
	{3, 6, 2, 7, 9, 4, 5, 7, 1, 9, 1, 3},
	{5, 2, 4, 7, 4, 8, 3, 7, 9, 1, 7, 3},
}

func scalarMultiply(a, b []int) int {
	result := 0
	for i := range a {
		for j := range b {
			result += a[i] * b[j]
		}
	}
	return result
}

func main() {
	t1 := time.Now()

	// create a channel which will synchronize processes
	results := make(chan []int)
	defer close(results)

	// arrange processes
	for i := range matrix {
		for j := range matrix {
			if i >= j {
				continue
			}
			// asyncronuosly call scalarMultiply
			go func(i, j int) {
				// always pass result with vectors location in matrix (which is i or j)
				results <- []int{scalarMultiply(matrix[i], matrix[j]), i, j}
			}(i, j)
		}
	}

	// cach results witch are running asyncronuosly
	min := int(^uint(0) >> 1) // set as the maximum possible int so every number should be less
	v1, v2 := 0, 0
	for i := range matrix {
		for j := range matrix {
			if i >= j {
				continue
			}
			res := <-results
			if min > res[0] {
				min = res[0]
				v1 = res[1]
				v2 = res[2]
			}
		}
	}

	t2 := time.Since(t1)
	fmt.Printf("finished in %d milliseconds\n", t2*time.Second)
	fmt.Printf("result is: %d\n", min)
	fmt.Printf("vector1: %#v\n", matrix[v1])
	fmt.Printf("vector2: %#v\n", matrix[v2])
}
