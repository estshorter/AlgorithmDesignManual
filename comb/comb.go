package comb

func comb(n, m int) int {
	bc := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		bc[i] = make([]int, n+1)
	}

	for i := 0; i <= n; i++ {
		bc[i][0] = 1
	}
	for i := 0; i <= n; i++ {
		bc[i][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			bc[i][j] = bc[i-1][j-1] + bc[i-1][j]
		}
	}
	return bc[n][m]
}
