package lowlink

import (
	"adm/graph"
	"reflect"
	"testing"
)

func Test_lowlink(t *testing.T) {
	// g := graph.New(9, [][]int{{1, 2, 5, 6}, {0, 2}, {0, 1, 3}, {2, 4}, {3}, {0, 6}, {0, 5, 7, 8}, {6, 8, 9}, {6, 7, 9}, {7, 8}})
	g := graph.New(10, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {3, 4},
		{1, 0}, {2, 1}, {2, 0}, {3, 2}, {4, 3},
		{0, 5}, {0, 6}, {5, 6}, {5, 0}, {6, 0}, {6, 5}, {6, 7}, {6, 8},
		{7, 6}, {7, 8}, {7, 9}, {8, 7}, {8, 6}, {8, 9}, {9, 8}, {9, 7},
	})
	type args struct {
		g *graph.Graph
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{g}, []int{3, 2, 6, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowlink(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowlink() = %v, want %v", got, tt.want)
			}
		})
	}
}
