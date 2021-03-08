package selection_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lst1 = []int{1, 9, 3, 2, 6, 5, 8, 7, 4}

func TestSelectionSort(t *testing.T) {
	fmt.Println(SelectionSort(lst1))
	fmt.Println("SELECTION SORT:")
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, lst1, "selection sort wrong")
}
