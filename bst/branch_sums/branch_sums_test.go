package branch_sums

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBranchSums(t *testing.T) {

	tests := []struct {
		input       *BinaryTree
		expectation []int
	}{
		{
			input:       NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expectation: []int{15, 16, 18, 10, 11},
		},
		{
			input:       NewBinaryTree(0, 1, 10, 100),
			expectation: []int{101, 10},
		},
		{
			input:       NewBinaryTree(1),
			expectation: []int{1},
		},
		{
			input:       NewBinaryTree(0, 1, 10, 100),
			expectation: []int{111},
		},
	}

	for _, tt := range tests {
		actual := BranchSums(tt.input)
		require.EqualValues(t, tt.expectation, actual)
	}

	////tree := NewBinaryTree(1, 1,2,3,4,5,6,7,8,9,10)
	//tree := NewBinaryTree(0, 1, 10, 100)
	//
	//fmt.Println(BranchSums(tree))

}
