package graph

/*
Graph is a data structure used to represent connections between elements.
Connections (edges) can be directed or undirected.
In a weighted graph, every edge has a weight.
Graphs can be used to represent database tables dependencies, electronic circuits and networks.
*/
type Graph[T any] interface {
	//Add new vertex to the graph
	AddVertex(vertex T)
	//Connect two vertices
	Connect(vert1, vert2 T)
	//Check if two vertices are adjacent to each other
	IsAdjacent(vert1, vert2 T) bool
	//Return all vertices adjacent to the specified vertex
	AdjacentVertices(vertex T) []T
}

type WeightedGraph[T any] interface {
	//Add new vertex to the graph
	AddVertex(vertex T)
	//Connect two vertices
	Connect(vert1, vert2 T, weight int)
	//Check if two vertices are adjacent to each other
	IsAdjacent(vert1, vert2 T) bool
	//Return all vertices adjacent to the specified vertex
	AdjacentVertices(vertex T) []T
	//Return weight of the edge between the specified vertices
	EdgeWeight(vert1, vert2 T) int
}
