package main

import (
	"fmt"
	sort "github.com/nghiatrandev/algorithms/sort/bubble-sort"
)

func main() {
	target := 61
	l := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	//l := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}

	fmt.Println(BinarySearch(l, target))

}

func BinarySearch(array []int, target int) int {

	array = sort.BubboleSort(array)
	currentArrayIndex := make([]int, 0)

	for i := 0; i < len(array); i++ {
		currentArrayIndex = append(currentArrayIndex, i)
	}

	for len(currentArrayIndex) > 1 {
		f := currentArrayIndex[0 : len(currentArrayIndex)/2]
		s := currentArrayIndex[len(currentArrayIndex)/2 : len(currentArrayIndex)]

		currentV := array[currentArrayIndex[len(currentArrayIndex)/2]-1]

		if currentV == target {
			return currentArrayIndex[len(currentArrayIndex)/2] - 1
		} else if (len(f) > 1 || len(s) > 1) && currentV > target {
			currentArrayIndex = f
		} else if (len(f) > 1 || len(s) > 1) && currentV < target {
			currentArrayIndex = s
		} else {
			break
		}
	}

	return -1
}
