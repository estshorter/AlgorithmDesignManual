package graph

import (
	"reflect"
	"testing"
)

func TestGraph_RemoveEdge(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.RemoveEdge(tt.args.i, tt.args.j)
		})
	}
}

func TestGraph_HasEdge(t *testing.T) {
	g := New(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
		want bool
	}{
		{"case1", g, args{0, 1}, true},
		{"case1", g, args{2, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.HasEdge(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Graph.HasEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_OutEdges(t *testing.T) {
	g := New(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
		want []int
	}{
		{"case1", g, args{0}, []int{1, 2}},
		{"case2", g, args{2}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.OutEdges(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.OutEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_InEdges(t *testing.T) {
	g := New(3, [][]int{{0, 1}, {0, 2}, {1, 2}})
	type args struct {
		i int
	}
	tests := []struct {
		name string
		g    *Graph
		args args
		want []int
	}{
		{"case1", g, args{2}, []int{0, 1}},
		{"case1", g, args{0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.InEdges(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.InEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
