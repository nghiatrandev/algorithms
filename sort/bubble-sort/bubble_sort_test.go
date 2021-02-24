package bubble_sort

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

var lst = []int{8, 22, 7, 9, 31, 19, 5, 13}

func TestBubbleSort(t *testing.T) {

	result := BubboleSort(lst)

	fmt.Println("BUBBLE SORT:")
	assert.Equal(t, result, []int{5, 7, 8, 9, 13, 19, 22, 31})

}
