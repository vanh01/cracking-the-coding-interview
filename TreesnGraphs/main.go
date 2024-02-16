package treesngraphs

import "fmt"

// Node represents a node in the directed graph
type Node struct {
	Value     int
	Neighbors []*Node
}

// Graph represents a directed graph
type Graph struct {
	Nodes map[int]*Node
}

// NewGraph creates a new directed graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(value int) {
	if _, exists := g.Nodes[value]; !exists {
		g.Nodes[value] = &Node{Value: value}
	}
}

// AddEdge adds a directed edge from node with value 'from' to node with value 'to'
func (g *Graph) AddEdge(from, to int) {
	fromNode, toNode := g.Nodes[from], g.Nodes[to]
	if fromNode == nil || toNode == nil {
		return // Node not found
	}
	fromNode.Neighbors = append(fromNode.Neighbors, toNode)
}

// PrintGraph prints the adjacency list representation of the graph
func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Printf("Node %d -> ", node.Value)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%d ", neighbor.Value)
		}
		fmt.Println()
	}
}
