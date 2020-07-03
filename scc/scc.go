package scc

import (
	"adm/graph"
	"adm/stack"
)

// https://scrapbox.io/pocala-kyopro/%E5%BC%B7%E9%80%A3%E7%B5%90%E6%88%90%E5%88%86%E5%88%86%E8%A7%A3

var ord *stack.Stack
var compo []int
var visited []bool

var rg, ng *graph.Graph

func stronglyConnectedComponents(g *graph.Graph) []int {
	compo = make([]int, g.NumVertexes())
	ord = stack.New()
	visited = make([]bool, g.NumVertexes())
	rg = graph.New(g.NumVertexes(), nil)
	ng = g

	for i := 0; i < g.NumVertexes(); i++ {
		compo[i] = -1
		for _, j := range g.OutEdges(i) {
			rg.AddEdge(j, i)
		}
	}
	for i := 0; i < g.NumVertexes(); i++ {
		dfs(i)
	}
	group := 0
	for !ord.IsEmpty() {
		idx := ord.Pop()
		if compo[idx] == -1 {
			rdfs(idx, group)
			group++
		}
	}
	return compo
}

func dfs(id int) {
	if visited[id] {
		return
	}
	visited[id] = true
	for _, e := range ng.OutEdges(id) {
		dfs(e)
	}
	ord.Push(id)
}

func rdfs(id, group int) {
	if compo[id] != -1 {
		return
	}
	compo[id] = group
	for _, e := range rg.OutEdges(id) {
		rdfs(e, group)
	}
}
