package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"1", []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := New()
			for _, val := range tt.vals {
				stack.Push(val)
			}
			for idx := range tt.vals {
				popped := stack.Pop()
				val := tt.vals[len(tt.vals)-1-idx]
				if val != popped {
					t.Errorf("Pop() = %v, want %v", popped, val)
				}

			}
		})
	}
}

func TestCapSize(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := New()
	for _, val := range a {
		s.Push(val)
	}
	for idx := range a {
		s.Pop()
		if idx == 3 && cap(s.a) != 10 {
			t.Errorf("wrong cap %v, want %v", cap(s.a), 10)
		}
		if idx == 5 && cap(s.a) != 6 {
			t.Errorf("wrong cap %v, want %v", cap(s.a), 6)
		}
		if idx == 6 && cap(s.a) != 4 {
			t.Errorf("wrong cap %v, want %v", cap(s.a), 4)
		}
		if idx == 7 && cap(s.a) != 2 {
			t.Errorf("wrong cap %v, want %v", cap(s.a), 2)
		}
		if idx == 8 && cap(s.a) != 1 {
			t.Errorf("wrong cap %v, want %v", cap(s.a), 1)
		}
	}
}
