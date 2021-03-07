package main

import (
	"fmt"
	sort "github.com/nghiatrandev/algorithms/sort/bubble-sort"
)

func main() {
	arr := []int{1, 4, 3, 6, 7, 6, 9, 0, 0, 9, 2, 4, 5, 6, 7, 8, 4, 1}

	sortedArr := sort.BubboleSort(arr)
	fmt.Print([]int{sortedArr[len(sortedArr)-3], sortedArr[len(sortedArr)-2], sortedArr[len(sortedArr)-1]})

}
