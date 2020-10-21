package main

import (
	"fmt"
	"testing"
)

var lst1 = []int{1, 9, 3, 2, 6, 5, 8, 7, 4}

func TestSelectionSort(t testing.T) {

	for i := 0; i < len(lst1); i++ {
		minI, minV := i, lst1[i]
		for j := i + 1; j < len(lst1); j++ {
			if minV > lst1[j] {
				minI = j
				minV = lst1[j]
			}
		}

		if minV != lst1[i] {
			lst1[minI], lst1[i] = lst1[i], lst1[minI]
		}
	}
	fmt.Println(lst1)
}
