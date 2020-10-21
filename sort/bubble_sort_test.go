package search

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

var lst = []int{8, 22, 7, 9, 31, 19, 5, 13}

func TestBubbleSort(t *testing.T) {
	count := 0
	for i := 0; i < len(lst)-1; i++ {
		for j := 1; j < len(lst); j++ {
			if lst[j] < lst[j-1] {
				lst[j], lst[j-1] = lst[j-1], lst[j]
				count++
			}
		}
	}

	fmt.Println("BUBBLE SORT:")
	assert.Equal(t, lst, []int{5, 7, 8, 9, 13, 19, 22, 31})

}
