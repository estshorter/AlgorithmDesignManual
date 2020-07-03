package comb

import "testing"

func Test_comb(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comb(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("comb() = %v, want %v", got, tt.want)
			}
		})
	}
}
