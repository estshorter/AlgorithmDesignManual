package toposort

import (
	"adm/graph"
	"adm/stack"
	"log"
)

var visited []bool
var processed []bool
var s *stack.Stack

func toposort(g *graph.Graph) []int {
	s = stack.New()
	visited = make([]bool, g.NumVertexes())
	processed = make([]bool, g.NumVertexes())
	for i := 0; i < g.NumVertexes(); i++ {
		if !visited[i] {
			dfs(g, i)
		}
	}
	ret := make([]int, s.Len())

	for i := 0; !s.IsEmpty(); i++ {
		ret[i] = s.Pop()
	}
	return ret
}

func dfs(g *graph.Graph, es int) {
	visited[es] = true
	for _, e := range g.OutEdges(es) {
		if visited[e] && !processed[e] { // 逆辺
			log.Fatalln("Not DAG!!")
			return
		}
		if !visited[e] {
			dfs(g, e)
		}
	}
	processed[es] = true
	s.Push(es)
}
