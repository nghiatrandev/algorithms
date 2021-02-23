package main

import "fmt"

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

func main() {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	output := graph.DepthFirstSearch([]string{})
	expectation := []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}

	fmt.Println(output)
	fmt.Println(expectation)

	//fmt.Println(test([]int{0}))
}

//func test(arr []int) []int {
//	count := arr[len(arr)-1]
//	arr = append(arr, count+1)
//	if count < 10 {
//		return test(arr)
//	}
//	return arr
//}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.

	//fmt.Println(n)
	output := []string{"A"}
	output = n.Search(output)

	return output
}

func (n *Node) Search(result []string) []string {

	for _, child := range n.Children {
		result = append(result, child.Name)
		//check child node, if != nil, call de quy
		if child.Children != nil {
			result = child.Search(result)
		}
	}

	return result
}
