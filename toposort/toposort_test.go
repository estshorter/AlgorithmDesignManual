package toposort

import (
	"adm/graph"
	"reflect"
	"testing"
)

func Test_toposort(t *testing.T) {
	g := graph.New(7, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 4}, {2, 5}, {4, 3}, {5, 4}, {6, 0}, {6, 5}})
	type args struct {
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{g}, []int{6, 0, 1, 2, 5, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toposort(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toposort() = %v, want %v", got, tt.want)
			}
		})
	}
}
