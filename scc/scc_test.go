package scc

import (
	"adm/graph"
	"reflect"
	"testing"
)

func Test_stronglyConnectedComponents(t *testing.T) {
	g := graph.New(8, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 0},
		{3, 5}, {1, 4}, {4, 5}, {5, 6}, {6, 4}, {3, 7}, {7, 5}})
	type args struct {
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{g}, []int{0, 0, 0, 0, 2, 2, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stronglyConnectedComponents(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stronglyConnectedComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}
