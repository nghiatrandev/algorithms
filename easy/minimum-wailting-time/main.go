package main

import (
	"fmt"
	sort "github.com/nghiatrandev/algorithms/sort/bubble-sort"
)

func main() {
	queries := []int{3, 2, 1, 2, 6}
	result := MinimumWaitingTime(queries)
	fmt.Println(result)
}

func MinimumWaitingTime(queries []int) int {
	// Write your code here.

	queries = sort.BubboleSort(queries)
	fmt.Println(queries)
	result := 0
	sum := 0
	for i := 0; i < len(queries)-1; i++ {
		sum += queries[i]
		result += sum
	}

	return result
}
