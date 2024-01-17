package graph

import (
	"fmt"

	linkedlist "github.com/Wine1y/go-cs/pkg/containers/linked_list"
)

type AdjacencyListsGraph[T comparable] struct {
	vertexToAdjList map[T]linkedlist.LinkedList[T]
}

func NewAdjacencyListsGraph[T comparable]() AdjacencyListsGraph[T] {
	return AdjacencyListsGraph[T]{
		vertexToAdjList: make(map[T]linkedlist.LinkedList[T]),
	}
}

func (graph *AdjacencyListsGraph[T]) AddVertex(vertex T) {
	graph.vertexToAdjList[vertex] = linkedlist.NewSinglyLinkedList[T]()
}

func (graph *AdjacencyListsGraph[T]) Connect(vert1, vert2 T) {
	list := graph.getVertexList(vert1)
	list.Append(vert2)
}

func (graph AdjacencyListsGraph[T]) IsAdjacent(vert1, vert2 T) bool {
	list := graph.getVertexList(vert1)
	if list.Length() == 0 {
		return false
	}
	for iterator := list.Iterator(); !iterator.Ended(); {
		if iterator.Next().Data() == vert2 {
			return true
		}
	}
	return false
}

func (graph AdjacencyListsGraph[T]) AdjacentVertices(vertex T) []T {
	list := graph.getVertexList(vertex)
	adjVertices := make([]T, 0, list.Length())
	if list.Length() == 0 {
		return adjVertices
	}
	for iterator := list.Iterator(); !iterator.Ended(); {
		adjVertices = append(adjVertices, iterator.Next().Data())
	}
	return adjVertices
}

func (graph AdjacencyListsGraph[T]) getVertexList(vertex T) linkedlist.LinkedList[T] {
	list, ok := graph.vertexToAdjList[vertex]
	if !ok {
		panic(fmt.Sprintf("Vertex not found in graph: %v", vertex))
	}
	return list
}

type AdjacencyListsWeightedGraph[T comparable] struct {
	vertexToAdjList map[T]linkedlist.LinkedList[graphEdge[T]]
}

type graphEdge[T comparable] struct {
	endVertex T
	weight    int
}

func NewAdjacencyListsWeightedGraph[T comparable]() AdjacencyListsWeightedGraph[T] {
	return AdjacencyListsWeightedGraph[T]{
		vertexToAdjList: make(map[T]linkedlist.LinkedList[graphEdge[T]]),
	}
}

func (graph *AdjacencyListsWeightedGraph[T]) AddVertex(vertex T) {
	graph.vertexToAdjList[vertex] = linkedlist.NewSinglyLinkedList[graphEdge[T]]()
}

func (graph *AdjacencyListsWeightedGraph[T]) Connect(vert1, vert2 T, weight int) {
	list := graph.getVertexList(vert1)
	list.Append(graphEdge[T]{endVertex: vert2, weight: weight})
}

func (graph AdjacencyListsWeightedGraph[T]) IsAdjacent(vert1, vert2 T) bool {
	list := graph.getVertexList(vert1)
	if list.Length() > 0 {
		for iterator := list.Iterator(); !iterator.Ended(); {
			if iterator.Next().Data().endVertex == vert2 {
				return true
			}
		}
	}
	return false
}

func (graph AdjacencyListsWeightedGraph[T]) AdjacentVertices(vertex T) []T {
	list := graph.getVertexList(vertex)
	adjVertices := make([]T, 0, list.Length())
	if list.Length() == 0 {
		return adjVertices
	}
	for iterator := list.Iterator(); !iterator.Ended(); {
		adjVertices = append(adjVertices, iterator.Next().Data().endVertex)
	}
	return adjVertices
}

func (graph AdjacencyListsWeightedGraph[T]) EdgeWeight(vert1, vert2 T) int {
	list := graph.getVertexList(vert1)
	if list.Length() > 0 {
		for iterator := list.Iterator(); !iterator.Ended(); {
			edge := iterator.Next().Data()
			if edge.endVertex == vert2 {
				return edge.weight
			}
		}
	}
	panic(fmt.Sprintf("Edge between %v and %v is not found", vert1, vert2))
}

func (graph AdjacencyListsWeightedGraph[T]) getVertexList(vertex T) linkedlist.LinkedList[graphEdge[T]] {
	list, ok := graph.vertexToAdjList[vertex]
	if !ok {
		panic(fmt.Sprintf("Vertex not found in graph: %v", vertex))
	}
	return list
}
