package perm

func perm(n int) [][]int {
	used := make(map[int]bool)
	return permHelper(n, 0, used, []int{}, [][]int{})
}

func permHelper(n, k int, used map[int]bool, p []int, ret [][]int) [][]int {
	if k == n {
		p2 := make([]int, len(p))
		copy(p2, p)
		return append(ret, p2)
	}

	for i := 0; i < n; i++ {
		if ok, v := used[i]; !ok || !v {
			used[i] = true
			ret = permHelper(n, k+1, used, append(p, i), ret)
			used[i] = false
		}
	}
	return ret
}
