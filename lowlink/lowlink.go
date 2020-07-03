package lowlink

import (
	"adm/graph"
)

// https://algo-logic.info/articulation-points/

var ord, low []int
var visited []bool

func lowlink(g *graph.Graph) []int {
	visited = make([]bool, g.NumVertexes())
	ord = make([]int, g.NumVertexes())
	low = make([]int, g.NumVertexes())

	aps := []int{}
	for i := 0; i < g.NumVertexes(); i++ {
		if !visited[i] {
			aps, _ = dfs(g, i, 0, -1, aps)
		}
	}
	// fmt.Println(ord, low)
	return aps
}

func dfs(g *graph.Graph, id, k, par int, aps []int) ([]int, int) {
	visited[id] = true
	ord[id] = k
	low[id] = k
	isAps := false
	cnt := 0
	k++
	// fmt.Println(id, k, par, g.OutEdges(id))
	for _, e := range g.OutEdges(id) {
		if !visited[e] {
			cnt++
			aps, k = dfs(g, e, k, id, aps)
			low[id] = min(low[id], low[e])
			if par != -1 && ord[id] <= low[e] {
				isAps = true
			}
		} else if e != par {
			low[id] = min(low[id], ord[e])
		}
	}
	if par == -1 && cnt == 2 {
		isAps = true
	}
	if isAps {
		aps = append(aps, id)
	}
	return aps, k
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
