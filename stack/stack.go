package stack

// Stack implements stack data structure
type Stack struct {
	a []int
}

// New initializes stack
func New() *Stack {
	return &Stack{[]int{}}
}

// Len returns number of elements
func (s *Stack) Len() int {
	return len(s.a)
}

// Slice returns slice of a stack
func (s *Stack) Slice() []int {
	return s.a
}

// Get returns ith element of stack
func (s *Stack) Get(i int) int {
	return s.a[i]
}

// Set sets a value to ith elment of stack
func (s *Stack) Set(i, val int) {
	s.a[i] = val
}

// Contain returns true if a stack contains a value v otherwise returns false
func (s *Stack) Contain(v int) bool {
	for _, val := range s.a {
		if val == v {
			return true
		}
	}
	return false
}

// Add adds a value v to a stack
func (s *Stack) Add(i, v int) {
	if len(s.a)+1 > cap(s.a) {
		s.resize()
	}
	s.a = s.a[:len(s.a)+1]
	for j := len(s.a) - 1; j > i; j-- {
		s.a[j] = s.a[j-1]
	}
	s.a[i] = v
}

// Remove removes a value v from a stack
func (s *Stack) Remove(i int) int {
	v := s.a[i]
	for j := i; j < len(s.a)-1; j++ {
		s.a[j] = s.a[j+1]
	}
	s.a = s.a[:len(s.a)-1]
	if cap(s.a) >= 3*len(s.a) {
		s.resize()
	}
	return v
}

// Push pushes a value v to a stack
func (s *Stack) Push(v int) {
	s.Add(s.Len(), v)
}

// Pop pops a value v from a stack
func (s *Stack) Pop() int {
	return s.Remove(s.Len() - 1)
}

func (s *Stack) resize() {
	b := make([]int, len(s.a), max(2*len(s.a), 1))
	copy(b, s.a)
	s.a = b
}

// IsEmpty returns true if stack size equal zero otherwise returns false
func (s *Stack) IsEmpty() bool {
	if len(s.a) == 0 {
		return true
	}
	return false
}

// Peek return the last-in element
func (s *Stack) Peek() int {
	return s.a[len(s.a)-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
