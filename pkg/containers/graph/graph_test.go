package graph

import (
	"slices"
	"testing"

	tu "github.com/Wine1y/go-cs/internal/testing_utils"
)

func testGraph(t *testing.T, graph Graph[int]) {
	tu.AssertPanics(t, "Graph panics when \"IsAdjacent\" is used on not existing vertices", func() {
		graph.IsAdjacent(10, 20)
	})

	tu.AssertPanics(t, "Graph panics when \"AdjacentVertices\" is used on not existing vertex", func() {
		graph.AdjacentVertices(10)
	})

	tu.AssertPanics(t, "Graph panics when \"Connect\" is used on not existing vertices", func() {
		graph.Connect(10, 20)
	})

	graph.AddVertex(10)
	graph.AddVertex(20)
	graph.AddVertex(30)

	tu.AssertEqualsNamed(t, "Graph vertices are initially not adjacent", graph.IsAdjacent(10, 20), false)
	tu.AssertEqualsNamed(t, "Graph vertices are initially not adjacent", graph.IsAdjacent(20, 30), false)

	graph.Connect(10, 20)
	graph.Connect(20, 30)
	tu.AssertEqualsNamed(t, "Graph vertices are adjacent after \"Connect\" operation", graph.IsAdjacent(10, 20), true)
	tu.AssertEqualsNamed(t, "Graph vertices are adjacent after \"Connect\" operation", graph.IsAdjacent(20, 30), true)

	graph.Connect(10, 30)
	expAdjVertices, adjVertices := []int{20, 30}, graph.AdjacentVertices(10)

	tu.AssertEqualsNamed(t, "Graph returns correct amount of adjacent vertices", len(adjVertices), len(expAdjVertices))
	for _, expVertex := range expAdjVertices {
		tu.AssertEqualsNamed(
			t,
			"Graph return correct adjacent vertices",
			slices.Contains(adjVertices, expVertex),
			true,
		)
	}
}

func testWeightedGraph(t *testing.T, graph WeightedGraph[int]) {
	tu.AssertPanics(t, "WeightedGraph panics when \"IsAdjacent\" is used on not existing vertices", func() {
		graph.IsAdjacent(10, 20)
	})

	tu.AssertPanics(t, "WeightedGraph panics when \"AdjacentVertices\" is used on not existing vertex", func() {
		graph.AdjacentVertices(10)
	})

	tu.AssertPanics(t, "WeightedGraph panics when \"Connect\" is used on not existing vertices", func() {
		graph.Connect(10, 20, 0)
	})

	graph.AddVertex(10)
	graph.AddVertex(20)
	graph.AddVertex(30)

	tu.AssertEqualsNamed(t, "WeightedGraph vertices are initially not adjacent", graph.IsAdjacent(10, 20), false)
	tu.AssertEqualsNamed(t, "WeightedGraph vertices are initially not adjacent", graph.IsAdjacent(20, 30), false)

	tu.AssertPanics(t, "WeightedGraph panics when \"EdgeWeight\" is used on not existing edge", func() {
		graph.EdgeWeight(10, 20)
	})

	graph.Connect(10, 20, 15)
	graph.Connect(20, 30, 25)
	tu.AssertEqualsNamed(t, "WeightedGraph vertices are adjacent after \"Connect\" operation", graph.IsAdjacent(10, 20), true)
	tu.AssertEqualsNamed(t, "WeightedGraph vertices are adjacent after \"Connect\" operation", graph.IsAdjacent(20, 30), true)
	tu.AssertEqualsNamed(t, "WeightedGraph edges have correct weights", graph.EdgeWeight(10, 20), 15)
	tu.AssertEqualsNamed(t, "WeightedGraph edges have correct weights", graph.EdgeWeight(20, 30), 25)

	graph.Connect(10, 30, 20)
	expAdjVertices, adjVertices := []int{20, 30}, graph.AdjacentVertices(10)

	tu.AssertEqualsNamed(t, "WeightedGraph returns correct amount of adjacent vertices", len(adjVertices), len(expAdjVertices))
	for _, expVertex := range expAdjVertices {
		tu.AssertEqualsNamed(
			t,
			"WeightedGraph return correct adjacent vertices",
			slices.Contains(adjVertices, expVertex),
			true,
		)
	}
}

func TestAdjacencyMatrixGraph(t *testing.T) {
	graph := NewAdjacencyMatrixGraph[int]()
	testGraph(t, &graph)
}

func TestAdjacencyMatrixWeightedGraph(t *testing.T) {
	graph := NewAdjacencyMatrixWeightedGraph[int]()
	testWeightedGraph(t, &graph)
}

func TestAdjacencyListsGraph(t *testing.T) {
	graph := NewAdjacencyListsGraph[int]()
	testGraph(t, &graph)
}

func TestAdjacencyListsWeightedGraph(t *testing.T) {
	graph := NewAdjacencyListsWeightedGraph[int]()
	testWeightedGraph(t, &graph)
}
