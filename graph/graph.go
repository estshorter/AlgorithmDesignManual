package graph

import (
	"adm/stack"
)

// Vertex defines a vertex of a graph
type Vertex struct {
	ID int
}

// Graph defines a graph data structure
type Graph struct {
	vertexes []Vertex
	adj      []stack.Stack // adjacency list
}

// New initializes graph
func New(n int, edges [][]int) *Graph {
	vertexes := make([]Vertex, n)
	adj := make([]stack.Stack, n)
	for k := range vertexes {
		vertexes[k].ID = k
	}
	for _, edge := range edges {
		in := edge[0]
		out := edge[1]
		adj[in].Push(out)
	}
	return &Graph{vertexes, adj}
}

// AddEdge adds the (i,j) edge to a graph
func (g *Graph) AddEdge(i, j int) {
	g.adj[i].Push(j)
}

// RemoveEdge removes a (i,j) edge from a graph if exists
func (g *Graph) RemoveEdge(i, j int) {
	for k := 0; k < g.adj[i].Len(); k++ {
		if g.adj[i].Get(k) == j {
			g.adj[i].Remove(k)
			return
		}
	}
}

// HasEdge return true if a graph has an edge (i,j)
func (g *Graph) HasEdge(i, j int) bool {
	return g.adj[i].Contain(j)
}

// OutEdges returns edges which point from a vertex i
func (g *Graph) OutEdges(i int) []int {
	if g.adj[i].Len() == 0 {
		return nil
	}
	edges := make([]int, g.adj[i].Len())
	for k := 0; k < g.adj[i].Len(); k++ {
		edges[k] = g.adj[i].Get(k)
	}
	return edges
}

// NumOutEdges returns number of edges which point from a vertex i
func (g *Graph) NumOutEdges(i int) int {
	return g.adj[i].Len()
}

// InEdges returns edges which point a vertex i
func (g *Graph) InEdges(i int) []int {
	var edges []int
	for j := range g.vertexes {
		if g.adj[j].Contain(i) {
			edges = append(edges, j)
		}
	}
	return edges
}

// NumVertexes returns number of vertexes of a graph
func (g *Graph) NumVertexes() int {
	return len(g.vertexes)
}
