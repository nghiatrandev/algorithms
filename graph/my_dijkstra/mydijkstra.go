package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	Name string
}

type Edge struct {
	Parent *Node
	Child  *Node
	Cost   int
}

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

const Infinity = int(^uint(0) >> 1)

func (g *Graph) AddEdge(parent, child *Node, cost int) {
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

func (g *Graph) AddNode(node *Node) {
	var isPresent bool

	for _, n := range g.Nodes {
		if n.Name == node.Name {
			isPresent = true
		}
	}

	if !isPresent {
		g.Nodes = append(g.Nodes, node)
	}
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Parent.Name + " -> " + edge.Child.Name + " = " + strconv.Itoa(edge.Cost)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}
	s += "\n"

	return s
}

func (g *Graph) newCostTable(startNode *Node) map[*Node]int {
	costTable := make(map[*Node]int)
	costTable[startNode] = 0

	for _, n := range g.Nodes {
		if n != startNode {
			costTable[n] = Infinity
		}
	}

	return costTable
}

func getClosestNonVisitedNode(costTable map[*Node]int, visited []*Node) *Node {
	type CostTableToSort struct {
		Node *Node
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
	fmt.Println(costTable)
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, CostTableToSort{node, cost})
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

func (g *Graph) getNodeEdges(node *Node) []*Edge {
	fmt.Println(g)
	nodeEdges := make([]*Edge, 0)
	for _, edge := range g.Edges {
		if edge.Parent == node {
			nodeEdges = append(nodeEdges, edge)
		}
	}

	return nodeEdges
}

func (g *Graph) dijkstra(startNode *Node) (shortestPathTable string) {
	costTable := g.newCostTable(startNode)

	// An empty list of "visited" Nodes. Everytime the algorithm runs on a Node, we add it here
	var visited []*Node

	for len(g.Nodes) != len(visited) {

		node := getClosestNonVisitedNode(costTable, visited)

		visited = append(visited, node)

		fmt.Println(node)

		nodeEdges := g.getNodeEdges(node)

		fmt.Println(nodeEdges)

		for _, edge := range nodeEdges {
			distanceToNeighbor := costTable[node] + edge.Cost

			if distanceToNeighbor < costTable[edge.Child] {
				costTable[edge.Child] = distanceToNeighbor
			}
		}
	}
	// Make the costTable nice to read :)
	for node, cost := range costTable {
		shortestPathTable += fmt.Sprintf("Distance from %s to %s = %d\n", startNode.Name, node.Name, cost)
	}

	return shortestPathTable
}

func main() {

	d := &Node{Name: "d"}
	c := &Node{Name: "c"}
	b := &Node{Name: "b"}
	g := &Node{Name: "g"}
	a := &Node{Name: "a"}
	e := &Node{Name: "e"}
	f := &Node{Name: "f"}

	//a := &Node{Name: "a"}
	//b := &Node{Name: "b"}
	//c := &Node{Name: "c"}
	//d := &Node{Name: "d"}
	//e := &Node{Name: "e"}
	//f := &Node{Name: "f"}
	//g := &Node{Name: "g"}

	graph := Graph{}
	graph.AddEdge(a, b, 2)
	graph.AddEdge(a, c, 5)
	graph.AddEdge(b, c, 6)
	graph.AddEdge(b, d, 1)
	graph.AddEdge(b, e, 3)
	graph.AddEdge(d, e, 4)
	graph.AddEdge(c, f, 8)
	graph.AddEdge(e, g, 9)
	graph.AddEdge(f, g, 7)

	//graph.dijkstra(a)

	//fmt.Println(graph.getNonVisitedClosest(a))
	fmt.Println(graph.dijkstra(a))
}
