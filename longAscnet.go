package main

import "fmt"

func find(A []int, n int) int {
	maxDiff := 0
	result := 0
	lastBegin := 0
	lastEnd := 0
	for i := 0; i < n-1; i++ {
		if A[i] <= A[i+1] {
			lastEnd = i
		} else {
			lastBegin = i + 1
		}
		diff := lastEnd - lastBegin
		if diff > maxDiff {
			maxDiff = diff
			result = lastBegin
			fmt.Println(lastBegin, lastEnd, result)
		}
	}
	return result
}

func main() {
	A := []int{1, 2, 3, 3, 2, -5, 5, 4, 3, 6, 7, 8, 1, 2, 3, 4, 5, 3, 2, 3, 3}
	result := find(A, len(A))
	fmt.Println(result)
}
