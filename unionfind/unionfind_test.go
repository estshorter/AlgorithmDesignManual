package unionfind

import (
	"testing"
)

func TestUnion_Find(t *testing.T) {
	u := New(10)
	u.Union(1, 9)
	u.Union(1, 2)
	u.Union(1, 4)
	type args struct {
		x int
	}
	tests := []struct {
		name string
		u    *Union
		args args
		want int
	}{
		{"1", u, args{9}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Find(tt.args.x); got != tt.want {
				t.Errorf("Union.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion_SameComp(t *testing.T) {
	u := New(10)
	u.Union(1, 9)
	u.Union(1, 2)
	u.Union(1, 4)

	type args struct {
		s1 int
		s2 int
	}
	tests := []struct {
		name string
		u    *Union
		args args
		want bool
	}{
		{"1", u, args{1, 9}, true},
		{"2", u, args{1, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.SameComp(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Union.SameComp() = %v, want %v", got, tt.want)
			}
		})
	}
}
