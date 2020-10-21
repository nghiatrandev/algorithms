package branch_sums

import (
	"fmt"
	"testing"
)

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func TestBranchSums(t *testing.T) {

	//tree := NewBinaryTree(1, 1,2,3,4,5,6,7,8,9,10)
	tree := NewBinaryTree(0, 1, 10, 100)

	fmt.Println(BranchSums(tree))

}

func BranchSums(root *BinaryTree) []int {

	sumBranches := make([]int, 0)
	sumBranches = append(sumBranches, 0)

	return root.sumsBranch(sumBranches)
}

func (tree *BinaryTree) sumsBranch(sumBranches []int) []int {
	sumBranches[len(sumBranches)-1] = sumBranches[len(sumBranches)-1] + tree.Value
	currentSum := sumBranches[len(sumBranches)-1]

	if tree.Left != nil {
		sumBranches = tree.Left.sumsBranch(sumBranches)
	}

	if tree.Right != nil && tree.Left != nil {
		sumBranches = append(sumBranches, currentSum)
		sumBranches = tree.Right.sumsBranch(sumBranches)
	}

	return sumBranches
}

//-------------------------------\

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}
