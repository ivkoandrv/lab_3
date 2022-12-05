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

 min := int(^uint(0) >> 1) // set as the maximum possible int so every number should be less
 v1, v2 := 0, 0
 for i := range matrix {
  for j := range matrix {
   if i >= j {
    continue
   }
   res := scalarMultiply(matrix[i], matrix[j])
   if min > res {
    min = res
    v1 = i
    v2 = j
   }
  }
 }

 t2 := time.Since(t1)
 fmt.Printf("finished in %d milliseconds\n", t2*time.Second)
 fmt.Printf("result is: %d\n", min)
 fmt.Printf("vector1: %#v\n", matrix[v1])
 fmt.Printf("vector2: %#v\n", matrix[v2])
}