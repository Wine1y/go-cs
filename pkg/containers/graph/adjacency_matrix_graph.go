package graph

import (
	"fmt"

	"github.com/Wine1y/go-cs/pkg/utils"
)

type adjacencyMatrixGraphImpl[T comparable, V any] struct {
	valueToVertexIndex map[T]int
	VertexIndexToValue map[int]T
	matrix             [][]V
	defaultAdjValue    V
}

func newAdjacencyMatrixGraphImpl[T comparable, V any](defaultAdjValue V) adjacencyMatrixGraphImpl[T, V] {
	return adjacencyMatrixGraphImpl[T, V]{
		valueToVertexIndex: make(map[T]int),
		VertexIndexToValue: make(map[int]T),
		matrix:             make([][]V, 0),
		defaultAdjValue:    defaultAdjValue,
	}
}

func (graph *adjacencyMatrixGraphImpl[T, V]) AddVertex(vertex T) {
	graph.valueToVertexIndex[vertex] = len(graph.valueToVertexIndex)
	graph.VertexIndexToValue[len(graph.VertexIndexToValue)] = vertex

	newRow := make([]V, len(graph.matrix))
	for i := 0; i < len(newRow); i++ {
		newRow[i] = graph.defaultAdjValue
	}
	graph.matrix = append(graph.matrix, newRow)

	for i := 0; i < len(graph.matrix); i++ {
		graph.matrix[i] = append(graph.matrix[i], graph.defaultAdjValue)
	}
}

func (graph adjacencyMatrixGraphImpl[T, V]) getAdjacentValue(i1, i2 int) V {
	return graph.matrix[i1][i2]
}

func (graph adjacencyMatrixGraphImpl[T, V]) setAdjacentValue(i1, i2 int, value V) {
	graph.matrix[i1][i2] = value
}

func (graph adjacencyMatrixGraphImpl[T, V]) getVertexIndex(vertex T) int {
	i, ok := graph.valueToVertexIndex[vertex]
	if !ok {
		panic(fmt.Sprintf("Vertex not found in graph: %v", vertex))
	}
	return i
}

func (graph adjacencyMatrixGraphImpl[T, V]) getValue(index int) T {
	return graph.VertexIndexToValue[index]
}

type AdjacencyMatrixGraph[T comparable] struct {
	adjacencyMatrixGraphImpl[T, bool]
}

func NewAdjacencyMatrixGraph[T comparable]() AdjacencyMatrixGraph[T] {
	return AdjacencyMatrixGraph[T]{
		adjacencyMatrixGraphImpl: newAdjacencyMatrixGraphImpl[T, bool](false),
	}
}

func (graph *AdjacencyMatrixGraph[T]) Connect(vert1, vert2 T) {
	v1 := graph.getVertexIndex(vert1)
	v2 := graph.getVertexIndex(vert2)
	graph.setAdjacentValue(v1, v2, true)
}

func (graph AdjacencyMatrixGraph[T]) IsAdjacent(vert1, vert2 T) bool {
	v1 := graph.getVertexIndex(vert1)
	v2 := graph.getVertexIndex(vert2)
	return graph.getAdjacentValue(v1, v2)
}

func (graph AdjacencyMatrixGraph[T]) AdjacentVertices(vertex T) []T {
	vIndex := graph.getVertexIndex(vertex)
	adjVertices := make([]T, 0, len(graph.matrix))
	for i := 0; i < len(graph.matrix); i++ {
		if graph.getAdjacentValue(vIndex, i) {
			adjVertices = append(adjVertices, graph.getValue(i))
		}
	}
	return adjVertices
}

func (graph *AdjacencyMatrixGraph[T]) BFSIterator(startVertex T) utils.Iterator[T] {
	iterator := NewGraphBFSIterator[T](graph, startVertex)
	return &iterator
}
func (graph *AdjacencyMatrixGraph[T]) DFSIterator(startVertex T) utils.Iterator[T] {
	iterator := NewGraphDFSIterator[T](graph, startVertex)
	return &iterator
}

type AdjacencyMatrixWeightedGraph[T comparable] struct {
	adjacencyMatrixGraphImpl[T, *int]
}

func NewAdjacencyMatrixWeightedGraph[T comparable]() AdjacencyMatrixWeightedGraph[T] {
	return AdjacencyMatrixWeightedGraph[T]{
		adjacencyMatrixGraphImpl: newAdjacencyMatrixGraphImpl[T, *int](nil),
	}
}

func (graph *AdjacencyMatrixWeightedGraph[T]) Connect(vert1, vert2 T, weight int) {
	v1 := graph.getVertexIndex(vert1)
	v2 := graph.getVertexIndex(vert2)
	graph.setAdjacentValue(v1, v2, &weight)
}

func (graph AdjacencyMatrixWeightedGraph[T]) IsAdjacent(vert1, vert2 T) bool {
	v1 := graph.getVertexIndex(vert1)
	v2 := graph.getVertexIndex(vert2)
	return graph.getAdjacentValue(v1, v2) != nil
}

func (graph AdjacencyMatrixWeightedGraph[T]) AdjacentVertices(vertex T) []T {
	vIndex := graph.getVertexIndex(vertex)
	adjVertices := make([]T, 0, len(graph.matrix))
	for i := 0; i < len(graph.matrix); i++ {
		if graph.getAdjacentValue(vIndex, i) != nil {
			adjVertices = append(adjVertices, graph.getValue(i))
		}
	}
	return adjVertices
}

func (graph AdjacencyMatrixWeightedGraph[T]) EdgeWeight(vert1, vert2 T) int {
	v1 := graph.getVertexIndex(vert1)
	v2 := graph.getVertexIndex(vert2)
	adjValue := graph.getAdjacentValue(v1, v2)
	if adjValue == nil {
		panic(fmt.Sprintf("Edge between %v and %v is not found", vert1, vert2))
	}
	return *adjValue
}

func (graph *AdjacencyMatrixWeightedGraph[T]) BFSIterator(startVertex T) utils.Iterator[T] {
	iterator := NewGraphBFSIterator[T](graph, startVertex)
	return &iterator
}
func (graph *AdjacencyMatrixWeightedGraph[T]) DFSIterator(startVertex T) utils.Iterator[T] {
	iterator := NewGraphDFSIterator[T](graph, startVertex)
	return &iterator
}
