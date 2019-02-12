package main

import (
	"fmt"
)

const (
	ColorWhite = iota // Vertex is not been visited yet
	ColorGray         // Vertex has been visited and still in the queue
	ColorBlack        // Vertex has been visited and was removed from the queue
)

type Graph struct {
	NumOfVertexes int
	AdjacencyList [][]int
	Colors        []int
	Distances     []int
	Predecessors  []int
}

func Constructor(numOfVertexes int) *Graph {
	return &Graph{
		NumOfVertexes: numOfVertexes,
		AdjacencyList: make([][]int, numOfVertexes),
		Colors:        make([]int, numOfVertexes),
		Distances:     make([]int, numOfVertexes),
		Predecessors:  make([]int, numOfVertexes),
	}
}

func (g *Graph) Reset() {
	for i := 0; i < g.NumOfVertexes; i++ {
		// Initialized as White color, representing this vertex is not been visited yet
		g.Colors[i] = ColorWhite
		// Max distance is number of vertexes - 1's edges
		g.Distances[i] = g.NumOfVertexes - 1
		// Initialized as -1, representing no predecessor of this vertex
		g.Predecessors[i] = -1
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

func (g *Graph) BFS(start int) {
	g.Reset()
	queue := []int{}
	i := start

	// Iterate the entire list to ensure all vertexes will be visited,
	// even for the unconnected components
	for j := 0; j < g.NumOfVertexes; j++ {
		if g.Colors[i] == ColorWhite {
			g.Colors[i] = ColorGray  // Vertex is visited
			g.Distances[i] = 0       // Each connected component's start vertex distance is 0
			g.Predecessors[i] = -1   // Each connected component's start vertex has no predecessor
			queue = append(queue, i) // Enqueue

			for len(queue) > 0 {
				// Get new search start vertex
				current := queue[0]

				// Iterate all adjacencies of current vertex
				for _, adjacency := range g.AdjacencyList[current] {
					if g.Colors[adjacency] == ColorWhite {
						g.Colors[adjacency] = ColorGray
						g.Distances[adjacency] = g.Distances[current] + 1
						g.Predecessors[adjacency] = current
						queue = append(queue, adjacency) // Enqueue
					}
				}

				queue = queue[1:]              // Dequeue
				g.Colors[current] = ColorBlack // Current vertex has been removed from queue
			}
		}

		// There may be some other vertexes not been visited yet even after all adjacencies of
		// current vertex have been visited (i.e. unconnected components)
		// Update i to j to ensue the remain unvisited vertexes will be visited
		i = j
	}
}

// Reference: https://goo.gl/ZMgFLe
func main() {
	graph := Constructor(9)
	graph.BFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	fmt.Println()
	graph.AddEdgeList(0, 1)
	graph.AddEdgeList(0, 2)
	graph.AddEdgeList(0, 3)
	graph.AddEdgeList(1, 0)
	graph.AddEdgeList(1, 4)
	graph.AddEdgeList(2, 0)
	graph.AddEdgeList(2, 4)
	graph.AddEdgeList(2, 5)
	graph.AddEdgeList(2, 6)
	graph.AddEdgeList(2, 7)
	graph.AddEdgeList(3, 0)
	graph.AddEdgeList(3, 7)
	graph.AddEdgeList(4, 1)
	graph.AddEdgeList(4, 2)
	graph.AddEdgeList(4, 5)
	graph.AddEdgeList(5, 2)
	graph.AddEdgeList(5, 4)
	graph.AddEdgeList(5, 8)
	graph.AddEdgeList(6, 2)
	graph.AddEdgeList(6, 7)
	graph.AddEdgeList(6, 8)
	graph.AddEdgeList(7, 2)
	graph.AddEdgeList(7, 3)
	graph.AddEdgeList(7, 6)
	graph.AddEdgeList(8, 5)
	graph.AddEdgeList(8, 6)
	graph.BFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	fmt.Println()
	graph.ClearEdges()
	graph.BFS(0)
	fmt.Printf("Colors: %v\n", graph.Colors)
	fmt.Printf("Distances: %v\n", graph.Distances)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
}
