package perm

import (
	"reflect"
	"testing"
)

func Test_perm(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{3}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := perm(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("perm() = %v, want %v", got, tt.want)
			}
		})
	}
}
