package branch_sums

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
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
