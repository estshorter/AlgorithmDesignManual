package subset

func subset(n int) [][]int {
	return subsetHelper(n, 0, []int{}, [][]int{})
}

func subsetHelper(n, k int, ss []int, ret [][]int) [][]int {
	if n == k {
		ss2 := make([]int, len(ss))
		copy(ss2, ss)
		return append(ret, ss)
	}

	ret = subsetHelper(n, k+1, append(ss, k+1), ret)
	ret = subsetHelper(n, k+1, ss, ret)
	return ret
}
