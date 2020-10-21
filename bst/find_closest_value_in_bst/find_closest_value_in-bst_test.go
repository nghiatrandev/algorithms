package bst

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type BST struct {
	Value int `json:"value"`

	Left  *BST `json:"left"`
	Right *BST `json:"right"`
}

func TestFindClosestValueInBst(t *testing.T) {
	tree := BST{
		Value: 10,
		Left: &BST{
			Value: 2,
			Left: &BST{
				Value: 4,
				Left: &BST{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &BST{
				Value: 5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left:  nil,
				Right: &BST{
					Value: 14,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &BST{
				Value: 22,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := tree.FindClosestValue(12)

	fmt.Println("FIND CLOSEST VALUE IN BST 'S RESULT:")
	fmt.Println(result)

	assert.Equal(t, 13, result, "112")

}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findCurrentResult(target, tree.Value)
}

func (tree *BST) findCurrentResult(target, closest int) int {

	if absdiff(tree.Value, target) < closest {
		closest = tree.Value
	}

	if target < tree.Value && tree.Left != nil {
		return tree.Left.findCurrentResult(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findCurrentResult(target, closest)
	} else {
		return closest
	}
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
