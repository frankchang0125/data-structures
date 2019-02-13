package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ColorWhite = iota
	ColorGray
	ColorBlack
)

// Find undirected graph's connected components key points:
//  1. If Vertex Y's predecessor: predecessor[Y] is Vertex X,
//     we can say that there's a path from Vertex X to Vertex Y
//  2. Since it's an undirected graph, if we can go from Vertex X
//     to Vertex Y, we can also go from Vertex Y to Vertex X,
//     which means Vertex X and Vertex Y are connected
//  3. Both DFS and BFS can produce Predecessors list of vertexes,
//     thus we can use it to find undirected graph's connected components
//     by iterating all vertex's predecessors
//  4. In order to determine whether Vertex M is in Component N (Set N),
//     we have to find out the root vertex of Component N (Set N) by
//     searching predecessors, which the time complexity is O(n)
//     To reduce search time from O(n) to O(1) of vertex M,
//     use Set Collapsing to inflate predecessors list
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

// FindCCDFS finds undirected graph's connected components by DFS
func (g *Graph) FindCCDFS(start int) {
	g.DFS(start)

	// Set Collapsing
	for i := 0; i < g.NumOfVertexes; i++ {
		g.SetCollapsing(i)
	}
}

// FindCCBFS finds undirected graph's connected components by BFS
func (g *Graph) FindCCBFS(start int) {
	g.BFS(start)

	// Set Collapsing
	for i := 0; i < g.NumOfVertexes; i++ {
		g.SetCollapsing(i)
	}
}

func (g *Graph) SetCollapsing(vertex int) {
	current := vertex

	for {
		if g.Predecessors[current] == -1 {
			break
		}
		current = g.Predecessors[current]
	}

	if current != vertex {
		g.Predecessors[vertex] = current
	}
}

func (g *Graph) PrintConnectedComponents() {
	set := make(map[int][]string)
	for i := 0; i < g.NumOfVertexes; i++ {
		var root int

		if g.Predecessors[i] == -1 {
			root = i
		} else {
			root = g.Predecessors[i]
		}

		if _, ok := set[root]; !ok {
			set[root] = make([]string, 0)
		}

		set[root] = append(set[root], strconv.Itoa(i))
	}

	for root, elements := range set {
		fmt.Printf("Component %d: %s\n", root, strings.Join(elements, ", "))
	}
}

// Reference: https://goo.gl/3AYTek
func main() {
	graph := Constructor(9)
	graph.FindCCDFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
	fmt.Println()
	graph.FindCCBFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
	fmt.Println()
	graph.AddEdgeList(0, 1)
	graph.AddEdgeList(1, 0)
	graph.AddEdgeList(1, 4)
	graph.AddEdgeList(1, 5)
	graph.AddEdgeList(3, 6)
	graph.AddEdgeList(4, 1)
	graph.AddEdgeList(4, 5)
	graph.AddEdgeList(5, 1)
	graph.AddEdgeList(5, 4)
	graph.AddEdgeList(5, 7)
	graph.AddEdgeList(6, 3)
	graph.AddEdgeList(6, 8)
	graph.AddEdgeList(7, 5)
	graph.AddEdgeList(8, 6)
	graph.FindCCDFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
	fmt.Println()
	graph.FindCCBFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
	fmt.Println()
	graph.ClearEdges()
	graph.FindCCDFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
	fmt.Println()
	graph.FindCCBFS(0)
	fmt.Printf("Predecessors: %v\n", graph.Predecessors)
	graph.PrintConnectedComponents()
}
