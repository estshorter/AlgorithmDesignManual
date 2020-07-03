package unionfind

//Union defines union-find data structure
type Union struct {
	p    []int
	size []int
	n    int
}

func New(n int) *Union {
	u := &Union{
		make([]int, n), make([]int, n), n,
	}
	for i := 0; i < n; i++ {
		u.p[i] = i
		u.size[i] = 1
	}
	return u
}

func (u *Union) Find(x int) int {
	if u.p[x] == x {
		return x
	}
	return u.Find(u.p[x])
}

func (u *Union) Union(s1, s2 int) {
	r1 := u.Find(s1)
	r2 := u.Find(s2)

	if r1 == r2 {
		return
	}

	if u.size[r1] >= u.size[r2] {
		u.size[r1] += u.size[r2]
		u.p[r2] = r1
	} else {
		u.size[r2] += u.size[r1]
		u.p[r1] = r2
	}
}

func (u *Union) SameComp(s1, s2 int) bool {
	return u.Find(s1) == u.Find(s2)
}
