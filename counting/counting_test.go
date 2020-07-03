package counting

import "testing"

func Test_counting(t *testing.T) {
	type args struct {
		a    []int
		x    int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 1, 2, 2, 2, 2, 2, 3, 4}, 2, 0, 8}, 5},
		{"2", args{[]int{2, 2, 2, 2, 2}, 2, 0, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := counting(tt.args.a, tt.args.x, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("counting() = %v, want %v", got, tt.want)
			}
		})
	}
}
