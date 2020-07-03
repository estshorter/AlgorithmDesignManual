package allroutes

import (
	"adm/graph"
)

func routes(g *graph.Graph, s, t int) [][]int {
	visited := make(map[int]bool)
	return routesHelper(g, t, 0, visited, []int{s}, [][]int{})
}

func routesHelper(g *graph.Graph, t, k int, visited map[int]bool, p []int, ret [][]int) [][]int {
	if p[k] == t {
		p2 := make([]int, len(p))
		copy(p2, p)
		return append(ret, p2)
	}

	for _, e := range g.OutEdges(p[k]) {
		if ok, v := visited[e]; !ok || !v {
			visited[e] = true
			ret = routesHelper(g, t, k+1, visited, append(p, e), ret)
			visited[e] = false
		}
	}
	return ret
}
