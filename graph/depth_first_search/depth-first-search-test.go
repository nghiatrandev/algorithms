package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	Name     string
	Children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func TestDepthFirstSearch(t *testing.T) {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	output := graph.DepthFirstSearch([]string{})
	expected := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}

	assert.Equal(t, output, expected, "depth first search wrong!!!")
}

//========================================================================

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.

	fmt.Println(array)

	return nil
}
