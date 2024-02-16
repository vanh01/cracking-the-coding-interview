package treesngraphs_test

import (
	"fmt"
	"testing"

	treesngraphs "github.com/vanh01/cracking-the-coding-interview/TreesnGraphs"
)

func TestNewGraph(t *testing.T) {
	graph := treesngraphs.NewGraph()

	graph.AddNode(0)
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 3)

	fmt.Println("Directed Graph:")
	graph.PrintGraph()
}
