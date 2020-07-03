package allroutes

import (
	"adm/graph"
	"reflect"
	"testing"
)

func Test_routes(t *testing.T) {
	g := graph.New(7, [][]int{{1, 2}, {2, 6}, {6, 3}, {1, 3}, {1, 4}, {4, 3}, {4, 6}, {6, 4}, {1, 5}, {5, 6}})
	type args struct {
		g *graph.Graph
		s int
		t int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{g, 1, 3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := routes(tt.args.g, tt.args.s, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("routes() = %v, want %v", got, tt.want)
			}
		})
	}
}
