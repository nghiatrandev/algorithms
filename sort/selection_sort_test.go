package search

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lst1 = []int{1, 9, 3, 2, 6, 5, 8, 7, 4}

func TestSelectionSort(t *testing.T) {

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
	fmt.Println(lst)
	fmt.Println("SELECTION SORT:")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, lst1, "selection sort wrong")
}
