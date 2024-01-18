package graph

import (
	"github.com/Wine1y/go-cs/pkg/containers/queue"
	"github.com/Wine1y/go-cs/pkg/containers/stack"
	"github.com/Wine1y/go-cs/pkg/utils"
)

/*
Graph is a data structure used to represent connections between elements.
Connections (edges) can be directed or undirected.
In a weighted graph, every edge has a weight.
Graphs can be used to represent database tables dependencies, electronic circuits and networks.
*/

type Graph[T comparable] interface {
	//Add new vertex to the graph
	AddVertex(vertex T)
	//Check if two vertices are adjacent to each other
	IsAdjacent(vert1, vert2 T) bool
	//Return all vertices adjacent to the specified vertex
	AdjacentVertices(vertex T) []T
	//Return Breadth-First-Search iterator
	BFSIterator(startVertex T) utils.Iterator[T]
	//Return Depth-First-Search iterator
	DFSIterator(startVertex T) utils.Iterator[T]
}
type UnweightedGraph[T comparable] interface {
	Graph[T]
	//Connect two vertices
	Connect(vert1, vert2 T)
}

type WeightedGraph[T comparable] interface {
	Graph[T]
	//Connect two vertices
	Connect(vert1, vert2 T, weight int)
	//Return weight of the edge between the specified vertices
	EdgeWeight(vert1, vert2 T) int
}

type GraphDFSIterator[T comparable] struct {
	graph   Graph[T]
	stack   stack.Stack[T]
	visited map[T]bool
}

func NewGraphDFSIterator[T comparable](graph Graph[T], startVertex T) GraphDFSIterator[T] {
	stack := stack.NewLinkedListStack[T]()
	stack.Push(startVertex)
	return GraphDFSIterator[T]{
		graph:   graph,
		stack:   &stack,
		visited: make(map[T]bool),
	}
}

func (iterator *GraphDFSIterator[T]) Next() T {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	vertex := iterator.stack.Pop()
	iterator.visited[vertex] = true
	for _, adjVertex := range iterator.graph.AdjacentVertices(vertex) {
		_, visited := iterator.visited[adjVertex]
		if !visited {
			iterator.stack.Push(adjVertex)
		}
	}
	return vertex
}

func (iterator GraphDFSIterator[T]) Ended() bool {
	return iterator.stack.Empty()
}

type GraphBFSIterator[T comparable] struct {
	graph   Graph[T]
	queue   queue.Queue[T]
	visited map[T]bool
}

func NewGraphBFSIterator[T comparable](graph Graph[T], startVertex T) GraphBFSIterator[T] {
	queue := queue.NewLinkedListQueue[T]()
	queue.EnQueue(startVertex)
	return GraphBFSIterator[T]{
		graph:   graph,
		queue:   &queue,
		visited: make(map[T]bool),
	}
}

func (iterator *GraphBFSIterator[T]) Next() T {
	if iterator.Ended() {
		panic("Iterator has ended")
	}
	vertex := iterator.queue.DeQueue()
	iterator.visited[vertex] = true
	for _, adjVertex := range iterator.graph.AdjacentVertices(vertex) {
		_, visited := iterator.visited[adjVertex]
		if !visited {
			iterator.queue.EnQueue(adjVertex)
		}
	}
	return vertex
}

func (iterator GraphBFSIterator[T]) Ended() bool {
	return iterator.queue.Empty()
}
