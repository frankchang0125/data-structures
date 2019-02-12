package main

import (
	"fmt"
)

const (
	ColorWhite = iota // Vertex is not been discovered yet
	ColorGray         // Vertex has been discovered but not finished yet
	ColorBlack        // Vertex has been discovered and finished
)

type Graph struct {
	NumOfVertexes int
	AdjacencyList [][]int
	Time          int
	Colors        []int
	Distances     []int
	Predecessors  []int
	Discovers     []int
	Finishes      []int
}

func Constructor(numOfVertexes int) *Graph {
	return &Graph{
		NumOfVertexes: numOfVertexes,
		AdjacencyList: make([][]int, numOfVertexes),
		Colors:        make([]int, numOfVertexes),
		Distances:     make([]int, numOfVertexes),
		Predecessors:  make([]int, numOfVertexes),
		Discovers:     make([]int, numOfVertexes),
		Finishes:      make([]int, numOfVertexes),
	}
}

func (g *Graph) Reset() {
	// Reset time
	g.Time = 0

	for i := 0; i < g.NumOfVertexes; i++ {
		// Initialized as White color, representing this vertex is not been visited yet
		g.Colors[i] = ColorWhite
		// Max distance is number of vertexes - 1's edges
		g.Distances[i] = g.NumOfVertexes - 1
		// Initialized as -1, representing no predecessor of this vertex
		g.Predecessors[i] = -1
		// Initialized as 0, representing not discovered yet
		g.Discovers[i] = 0
		// Initialized as 0, representing not finished yet
		g.Finishes[i] = 0
	}
}

func (g *Graph) AddEdgeList(from, to int) {
	if g.AdjacencyList[from] == nil {
		g.AdjacencyList[from] = make([]int, 0)
	}
	g.AdjacencyList[from] = append(g.AdjacencyList[from], to)
}

func (g *Graph) ClearEdges() {
	for i := 0; i < g.NumOfVertexes; i++ {
		g.AdjacencyList[i] = nil
	}
}

func (g *Graph) DFS(start int) {
	g.Reset()
	i := start

	for j := 0; j < g.NumOfVertexes; j++ {
		if g.Colors[i] == ColorWhite {
			g.Distances[i] = 0     // Each connected component's start vertex
			g.Predecessors[i] = -1 // Each connected component's start vertex has no predecessor
			g.DFSVisit(i)
		}

		// There may be some other vertexes not been visited yet even after all adjacencies of
		// current vertex have been visited (i.e. unconnected components)
		// Update i to j to ensue the remain unvisited vertexes will be visited
		i = j
	}
}

func (g *Graph) DFSVisit(vertex int) {
	g.Time++
	g.Colors[vertex] = ColorGray // Vertex is visited
	g.Discovers[vertex] = g.Time // Update vertex's discovered time

	// Iterate all adjacencies of vertex
	for _, adjacency := range g.AdjacencyList[vertex] {
		if g.Colors[adjacency] == ColorWhite {
			g.Distances[adjacency] = g.Distances[vertex] + 1
			g.Predecessors[adjacency] = vertex
			g.DFSVisit(adjacency) // Recursive call to visit adjacency vertex
		}
	}

	g.Time++
	g.Colors[vertex] = ColorBlack // Vertex is discovered and finished
	g.Finishes[vertex] = g.Time   // Update vertex's finished time
}

// Reference: https://goo.gl/ytmNC4
func main() {
	graph := Constructor(8)
	graph.DFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	fmt.Printf("Discovers: %v\n", graph.Discovers)
	fmt.Printf("Finished: %v\n", graph.Finishes)
	fmt.Println()
	graph.AddEdgeList(0, 1)
	graph.AddEdgeList(0, 2)
	graph.AddEdgeList(1, 3)
	graph.AddEdgeList(2, 1)
	graph.AddEdgeList(2, 5)
	graph.AddEdgeList(3, 4)
	graph.AddEdgeList(3, 5)
	graph.AddEdgeList(5, 1)
	graph.AddEdgeList(6, 4)
	graph.AddEdgeList(6, 7)
	graph.AddEdgeList(7, 6)
	graph.DFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	fmt.Printf("Discovers: %v\n", graph.Discovers)
	fmt.Printf("Finished: %v\n", graph.Finishes)
	fmt.Println()
	graph.ClearEdges()
	graph.DFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	fmt.Printf("Discovers: %v\n", graph.Discovers)
	fmt.Printf("Finished: %v\n", graph.Finishes)
}
